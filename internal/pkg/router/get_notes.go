package router

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

func (r *Router) GetNotes(ctx context.Context) *Route {
	return &Route{
		path:   "/notes/:course",
		method: echo.GET,
		handlerFunc: func(c echo.Context) error {
			course := c.Param("course")
			notes, err := r.service.GetNotes(ctx, course)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, notes)
		},
	}
}
