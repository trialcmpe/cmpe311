package router

import (
	"context"
	"net/http"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/labstack/echo"
)

func (r *Router) SignUPRoute(ctx context.Context) *Route {
	return &Route{
		path:   "/signup",
		method: echo.POST,
		handlerFunc: func(c echo.Context) error {
			var user model.User
			c.Bind(&user)
			token, err := r.service.SignUP(ctx, &user)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, *token)
		},
	}
}
