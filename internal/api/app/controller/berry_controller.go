package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juniorrodes/poke-htmx/internal/api/app/entities"
	"github.com/juniorrodes/poke-htmx/internal/api/infrastructure/client"
	internalRouter "github.com/juniorrodes/poke-htmx/internal/api/infrastructure/router"
	"github.com/juniorrodes/poke-htmx/internal/components"
)

type BerryController interface {
	GetBerry(*http.Request) internalRouter.Response
}

type berryController struct {
	pokeClient *client.PokeClient
}

func NewBerryController() BerryController {
	return &berryController{
		pokeClient: client.NewPokeClient(),
	}
}

func (c *berryController) GetBerry(r *http.Request) internalRouter.Response {
	name := r.FormValue("name")

	pokeResponse, err := c.pokeClient.DoRequest(fmt.Sprintf("/v2/berry/%s", name))
	response := internalRouter.Response{}
	if err != nil {
		return internalRouter.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(err.Error()),
		}
	}
	berry := entities.Berry{}
	err = json.Unmarshal(pokeResponse, &berry)
	if err != nil {
		return internalRouter.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(err.Error()),
		}
	}

	err = components.BerryTable(berry).Render(r.Context(), &response)
	if err != nil {
		return internalRouter.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(err.Error()),
		}
	}
	response.StatusCode = 200

	return response
}
