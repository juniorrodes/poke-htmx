package main

import (
	"github.com/juniorrodes/poke-htmx/internal/api/app/controller"
	internalRouter "github.com/juniorrodes/poke-htmx/internal/api/infrastructure/router"
)

func main() {
	c := controller.NewController()

	r := internalRouter.NewRouter()
	r.POST("/clicked", c.AddNameToGreet)

	r.ListenAndServe(":8080")
}
