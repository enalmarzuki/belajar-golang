//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-dependency-injection/app"
	"belajar-golang-dependency-injection/controller"
	"belajar-golang-dependency-injection/middleware"
	"belajar-golang-dependency-injection/repository"
	"belajar-golang-dependency-injection/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		app.NewValidatorOptions,
		app.NewValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
