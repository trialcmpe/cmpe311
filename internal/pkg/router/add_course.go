package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/labstack/echo"
)

type AddCourseBody struct {
	Name string `json:"name"`
}

func (r *Router) AddCourseRoute(ctx context.Context) *Route {
	return &Route{
		path:   "/courses/add",
		method: echo.POST,
		handlerFunc: func(c echo.Context) error {
			var body AddCourseBody
			c.Bind(&body)
			if err := r.service.AddCourse(ctx, &model.Course{
				Name: body.Name,
			}); err != nil {
				fmt.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
			fmt.Println("bura..")
			return c.String(http.StatusOK, "course	 added !")
		},
	}
}
