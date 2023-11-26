package endpoint

import (
	"avitoStart/internal/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Service interface {
	//users
	AddUser(user model.User) (bool, error)
	DeleteUser(id string) (bool, error)
	ExtractUsers() ([]model.User, error)

	//slug
	DeleteSlug(name string) (bool, error)
	CreateSlug(name string) (bool, error)
	ExecSlugNamesUser(iduser string) ([]model.Slug, error)
	MasterFunc(data model.MasterData) (bool, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) AddUser(ctx echo.Context) error {
	user := model.User{}

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Println(user)
	res, err := e.s.AddUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (e *Endpoint) DeleteUser(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Println(user)
	res, err := e.s.DeleteUser(user.Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (e *Endpoint) ExtractUsers(ctx echo.Context) error {
	res, err := e.s.ExtractUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (e *Endpoint) DeleteSlug(ctx echo.Context) error {
	var name model.SlugQuery
	err := ctx.Bind(&name)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Println(name.Name)
	res, err := e.s.DeleteSlug(name.Name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (e *Endpoint) CreateSlug(ctx echo.Context) error {
	var name model.SlugQuery
	err := ctx.Bind(&name)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Println(name.Name)
	res, err := e.s.CreateSlug(name.Name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (e *Endpoint) ExecSlugNamesUser(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Println(user)
	res, err := e.s.ExecSlugNamesUser(user.Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (e *Endpoint) MasterFunc(ctx echo.Context) error {
	var data model.MasterData
	err := ctx.Bind(&data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	log.Println(data)
	res, err := e.s.MasterFunc(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
