package entities

import "github.com/juniorrodes/poke-htmx/internal/api/app/entities/utils"

type Berry struct {
	Id               int                    `json:"id"`
	Name             string                 `json:"name"`
	GrowthTime       int                    `json:"growth_time"`
	MaxHarvest       int                    `json:"max_harvest"`
	NaturalGiftPower int                    `json:"natural_gift_power"`
	Size             int                    `json:"size"`
	Smoothness       int                    `json:"smoothness"`
	SoilDryness      int                    `json:"soil_dryness"`
	Fimrness         utils.NamedAPIResource `json:"firmness"`
	Flavors          []BerryFlavor          `json:"flavors"`
	Item             utils.NamedAPIResource `json:"item"`
	NaturalGiftType  utils.NamedAPIResource `json:"natural_gift_type"`
}

type BerryFlavorMap struct {
	Potency int                    `json:"potency"`
	Flavor  utils.NamedAPIResource `json:"flavor"`
}

type BerryFirmness struct {
	Id      int                    `json:"id"`
	Name    string                 `json:"name"`
	Berries utils.NamedAPIResource `json:"berries"`
	Names   []utils.Name           `json:"names"`
}

type BerryFlavor struct {
	Id          int                    `json:"id"`
	Name        string                 `json:"name"`
	Berries     []BerryFlavorMap       `json:"berries"`
	ContestType utils.NamedAPIResource `json:"contest_type"`
	Names       []utils.Name           `json:"names"`
}
