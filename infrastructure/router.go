package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hukurou-s/user-auth-api-sample/crypt"
	"github.com/hukurou-s/user-auth-api-sample/interfaces/controllers"
)

var Echo *echo.Echo

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	controller := controllers.NewUserController(NewSqlHandler())

	// controller.Login(c)の引数は自分で定義したcontextなので、
	// echo.Context型のcを引数として受け取る関数の中で呼び出す必要がある
	e.POST("/login", func(c echo.Context) error { return controller.Login(c) })

	pubKey := crypt.NewPublicKey()

	// /user以下ではjwtによる認証が必要になる
	u := e.Group("/user")
	u.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    pubKey,
		SigningMethod: "RS256",
	}))
	u.GET("/profile", func(c echo.Context) error { return controller.UserProfile(c) })

	Echo = e
}
