package router

import (
	"context"

	"github.com/canergulay/bilgipedia/internal/app"
	"github.com/canergulay/bilgipedia/internal/pkg/authentication"
	"github.com/canergulay/bilgipedia/internal/pkg/model"
	"github.com/canergulay/bilgipedia/internal/pkg/router/middleware"
	"github.com/labstack/echo"
)

type IService interface {
	SignUP(ctx context.Context, user *model.User) (*string, error)
	SignIN(ctx context.Context, user *model.User) (*string, error)
	GetNotes(ctx context.Context, course string) (*[]model.Note, error)
	AddNote(ctx context.Context, note *model.Note) error
	AddCourse(ctx context.Context, course *model.Course) error
	GetCourse(ctx context.Context, courseName string) (*model.Course, error)
	GetCourses(ctx context.Context) (*[]model.Course, error)
}

type Router struct {
	Routes         []app.Route
	service        IService
	AuthMiddleware echo.MiddlewareFunc
}

func InitializeRouter(ctx context.Context, s IService, jmanager *authentication.JwtManager) Router {
	router := Router{Routes: []app.Route{}, service: s}
	router.Routes = append(router.Routes, *router.SignUPRoute(ctx))
	router.Routes = append(router.Routes, *router.SignINRoute(ctx))
	router.Routes = append(router.Routes, *router.AddNoteRoute(ctx))
	router.Routes = append(router.Routes, *router.GetNotes(ctx))
	router.Routes = append(router.Routes, *router.GetCoursesRoute(ctx))
	router.Routes = append(router.Routes, *router.AddCourseRoute(ctx))

	router.AuthMiddleware = middleware.JwtVerifer(jmanager)
	return router
}

type Route struct {
	path        string
	method      string
	handlerFunc func(echo.Context) error
}

func (r Route) Path() string {
	return r.path
}

func (r Route) Method() string {
	return r.method
}
func (r Route) HandlerFunc() func(echo.Context) error {
	return r.handlerFunc
}

func NewRoute(Path, method string, handlerFunc echo.HandlerFunc) *Route {
	return &Route{path: Path, method: method, handlerFunc: handlerFunc}
}
