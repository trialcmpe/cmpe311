package router

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

func (r *Router) GetCoursesRoute(ctx context.Context) *Route {
	return &Route{
		path:   "/courses",
		method: echo.GET,
		handlerFunc: func(c echo.Context) error {
			courses, err := r.service.GetCourses(ctx)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, courses)
		},
	}
}
