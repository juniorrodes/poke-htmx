package utils

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}
