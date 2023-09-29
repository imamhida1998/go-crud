package main

import (
	"go-crud/lib/db"
	"go-crud/service/config"
	"go-crud/service/handler"
	"go-crud/service/repository"
	"go-crud/service/usecase"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func main() {
	route := gin.Default()
	set := config.Config{}
	set.CatchError(set.InitEnv())
	Database := set.GetDBConfig()
	db, err := db.ConnectiontoMYSQL(Database)
	if err != nil {
		log.Println(err)
		return
	}

	UserRepo := repository.NewRepoUser(db)
	AuthUsecase := usecase.NewJWTService()
	UserUsecase := usecase.NewUsecaseUser(UserRepo, AuthUsecase)

	HanderUser := handler.NewHandlerUser(UserUsecase)

	user := route.Group("/api/user")

	user.POST("/registration", HanderUser.RegistrationDataUser)
	user.GET("/login", HanderUser.LoginDataUser)
	user.PUT("/:id", authMiddleware(AuthUsecase, UserRepo), HanderUser.UpdateDataUser)
	user.GET("/:id", authMiddleware(AuthUsecase, UserRepo), HanderUser.DetailUsers)
	server := http.Server{
		Addr:    "127.0.0.1:8002",
		Handler: route,
	}
	server.ListenAndServe()
	if err != nil {
		return
	}

}

func authMiddleware(authService usecase.Auth, userService repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		Email := (claim["email"].(string))

		user, err := userService.GetUsersByEmail(Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Set("CurrentUser", user)

	}
}
