package pokeapi

type ResExplore struct {
	EncounterRates []any `json:"encounter_method_rates"`
	GameIndex      int   `json:"game_index"`
	ID             int   `json:"id"`
	Location       struct {
		name string
		url  string
	} `json:"location"`
	Name              string `json:"name"`
	Names             []any  `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []any `json:"version_details"`
	} `json:"pokemon_encounters"`
}
