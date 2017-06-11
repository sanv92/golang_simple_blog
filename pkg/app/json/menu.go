package json

const MenuFile = "pkg/config/menu.json"

type Menu struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
}
