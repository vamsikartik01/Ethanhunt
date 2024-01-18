package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	helper "github.com/vamsikartik01/Ethanhunt/helpers"
	"github.com/vamsikartik01/Ethanhunt/models"
	"github.com/vamsikartik01/Ethanhunt/routes"
	"github.com/vamsikartik01/Ethanhunt/services/mysql"
)

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie("auth")
		if err != nil {
			c.Response().Header().Set("Location", "https://jamesbond.3dns.me/main")
			return c.String(http.StatusTemporaryRedirect, "Redirecting...")
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(helper.Config.Jwt.Secret), nil
		})

		if err != nil {
			c.Response().Header().Set("Location", "https://jamesbond.3dns.me/main")
			return c.String(http.StatusTemporaryRedirect, "Redirecting...")
		}

		if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
			c.Set("user", claims)
		} else {
			c.Response().Header().Set("Location", "https://jamesbond.3dns.me/main/auth/signin")
			return c.String(http.StatusTemporaryRedirect, "Redirecting...")
		}

		return next(c)
	}
}

func main() {
	log.Println("Hello Ethan! Api is listening on port :2000")

	err := helper.LoadConfig()
	if err != nil {
		log.Println("Config file with error - ", err)
	}

	err = mysql.InitConnection()
	if err != nil {
		log.Println("My sql with error - ", err)
	}

	ethan := echo.New()
	ethan.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	ethan.Use(jwtMiddleware)
	routes.DefineRoutes(ethan)
	ethan.Logger.Fatal(ethan.Start(":2000"))
}
