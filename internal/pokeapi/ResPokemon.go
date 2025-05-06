package pokeapi

type ResPokemon struct {
	Abilities              []any  `json:"abilities"`
	BaseExperience         int    `json:"base_experience"`
	Cries                  any    `json:"cries"`
	Forms                  []any  `json:"forms"`
	GameIndices            []any  `json:"game_indices"`
	Height                 int    `json:"height"`
	HeldItems              []any  `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []any  `json:"moves"`
	Name                   string `json:"name"`
	Order                  int    `json:"order"`
	PastAbilities          []any  `json:"past_abilities"`
	PastTypes              []any  `json:"past_types"`
	Weight                 int    `json:"weight"`
	Species                any    `json:"species"`
	Sprites                any    `json:"sprites"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	} `json:"types"`
}
