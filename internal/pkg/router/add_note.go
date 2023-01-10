package router

import (
	"context"
	"net/http"

	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/canergulay/bilgipedia/internal/pkg/utils"
	"github.com/labstack/echo"
)

func (r *Router) AddNoteRoute(ctx context.Context) *Route {
	return &Route{
		path:   "/notes/add",
		method: echo.POST,
		handlerFunc: func(c echo.Context) error {
			userID := c.Get("user")
			var note model.Note
			c.Bind(&note)
			r.service.AddNote(utils.UserIDToContext(ctx, userID), &note)
			return c.String(http.StatusOK, "note added !")
		},
	}
}
