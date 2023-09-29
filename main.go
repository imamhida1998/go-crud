package main

import (
	"go-crud/lib/db"
	"go-crud/service/config"
	"go-crud/service/handler"
	"go-crud/service/repo"
	"go-crud/service/usecase"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	sesion "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

func main() {
	route := fiber.New()

	store := sesion.New()
	ctx := route.AcquireCtx(&fasthttp.RequestCtx{})
	s, err := store.Get(ctx)
	if err != nil {
		panic(err)
	}

	set := config.Config{}
	set.CatchError(set.InitEnv())
	Database := set.GetDBConfig()
	db, err := db.ConnectiontoMYSQL(Database)
	if err != nil {
		log.Println(err)
		return
	}

	UserRepo := repo.NewRepoUser(db)
	AuthUsecase := usecase.NewJWTService()
	UserUsecase := usecase.NewUsecaseUser(UserRepo, AuthUsecase)

	HanderUser := handler.NewHandlerUser(UserUsecase, s)

	apps := route.Group("/api/")
	users := apps.Group("/users")
	users.Use(authMiddleware(s, AuthUsecase, UserRepo))

	apps.Post("/registration", HanderUser.RegistrationDataUser)
	apps.Get("/login", HanderUser.LoginDataUser)
	users.Put("/:id", HanderUser.UpdateDataUser)
	users.Get("/:id", HanderUser.DetailUsers)

	err = route.Listen("127.0.0.1:8002")
	if err != nil {
		return
	}

}

func authMiddleware(s *sesion.Session, authService usecase.Auth, userService repo.UserRepo) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.Status(http.StatusUnauthorized).SendString("Unauthorized")
			return nil
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.Status(http.StatusUnauthorized).SendString("Unauthorized")
			return err
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.Status(http.StatusUnauthorized).SendString("Unauthorized")
			return err
		}

		Email := (claim["email"].(string))

		user, err := userService.GetUsersByEmail(Email)
		if err != nil {
			c.Status(http.StatusUnauthorized).SendString("Unauthorized")
			return err
		}

		s.Set("CurrentUser", user)

		c.Next()
		return nil

	}
}
