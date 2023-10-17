package controller

import (
	"net/http"

	internalRouter "github.com/juniorrodes/poke-htmx/internal/api/infrastructure/router"
	"github.com/juniorrodes/poke-htmx/internal/components"
)

type clickerController struct {
	names []string
}

func NewController() *clickerController {
	return &clickerController{}
}

func (c *clickerController) AddNameToGreet(r *http.Request) internalRouter.Response {
	response := internalRouter.Response{}
	name := r.FormValue("name")
	c.names = append(c.names, name)

	components.Header(c.names).Render(r.Context(), &response)
	response.StatusCode = http.StatusOK

	return response
}
