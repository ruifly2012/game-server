package web_service

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	startHttpServer()
}

/*
 * Http Server
 * serve client request: register|login|loginByGuest
 */
func startHttpServer() {
	app := iris.New()

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	registerHandlers(app)

	app.Run(iris.Addr(":8080"))
}

func registerHandlers(app *iris.Application) {
	app.Post("/register", registerHandler)
	app.Post("/login", loginHandler)
	app.Post("/loginByGuest", loginByGuestHandler)
}

func registerHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":     "failed",
		"error_code": "error_password_invalid"})
	return
}

func loginHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":     "failed",
		"error_code": "error_password_invalid"})
	return
}

func loginByGuestHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"status":     "failed",
		"error_code": "error_password_invalid"})
	return
}
