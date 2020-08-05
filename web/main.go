package main

import (
	"github.com/wangjian890523/superstar/bootstrap"
	"github.com/wangjian890523/superstar/web/middleware/identity"
	"github.com/wangjian890523/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "王健")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)

	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")

}
