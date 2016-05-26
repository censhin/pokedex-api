package pokemon

type Pokemons struct {
	Pokemons []Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string    `json:"name"`
	Number         int       `json:"number"`
	Generation     string    `json:"generation"`
	Types          []string  `json:"types"`
	BaseStats      BaseStats `json:"baseStats"`
	Genus          string    `json:"genus"`
	Height         int       `json:"height"`
	Weight         int       `json:"weight"`
	CaptureRate    int       `json:"captureRate"`
	BaseExperience int       `json:"baseExperience"`
	Abilities      []string  `json:"abilities"`
	EffortValues   BaseStats `json:"effortValues"`
	EggGroups      []string  `json:"eggGroups"`
	Evolution      Evolution `json:"evolution"`
}

type BaseStats struct {
	Hp             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"specialAttack"`
	SpecialDefence int `json:"specialDefense"`
	Speed          int `json:"speed"`
}

type Evolution struct {
	Level  int    `json:"level"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}
