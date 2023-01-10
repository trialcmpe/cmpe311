package middleware

import (
	"fmt"

	"github.com/canergulay/bilgipedia/internal/pkg/authentication"
	"github.com/labstack/echo"
)

var (
	noTokenFound  string = "Unauthorized, please enter a valid token !"
	tokenNotValid string = "Unauthorized, your token is not valid !"
)

func JwtVerifer(jmanager *authentication.JwtManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			fmt.Println(token)
			// IF THERE IS NO A HEADER VALUE WITH THE KEY ABOVE, WE WILL SIMPLE TERMINATE THIS REQUEST AND RETURN 403
			if len(token) < 10 {
				return c.JSON(403, noTokenFound)
			}
			fmt.Println("here")
			//IF THERE IS A TOKEN, WE WILL THEN CHECK IF IT IS A VALID ONE
			userID, err := jmanager.JwtCredentialsVerifier(token)
			if err != nil {
				fmt.Println(err)
				return c.JSON(402, err)
			}
			fmt.Println(userID)
			// OTHERWISE WE WILL SET OUR USER OBJECT THAT WE JUST PARSED FROM JWT PAYLOAD
			c.Set("user", *userID)
			next(c)
			return nil
		}
	}
}
