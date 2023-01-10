package router

import (
	"context"
	"net/http"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/labstack/echo"
)

func (r *Router) SignINRoute(ctx context.Context) *Route {
	return &Route{
		path:   "/signin",
		method: echo.POST,
		handlerFunc: func(c echo.Context) error {
			var user model.User
			c.Bind(&user)
			jwt, err := r.service.SignIN(ctx, &model.User{
				Email:    user.Email,
				Password: user.Password,
			})
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, *jwt)
		},
	}
}
