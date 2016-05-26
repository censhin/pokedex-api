package pokemon

type Pokemons struct {
	Pokemons []Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string                 `bson:"name" json:"name"`
	Number         int                    `bson:"number" json:"number"`
	Generation     string                 `bson:"generation" json:"generation"`
	Types          []string               `bson:"types" json:"types"`
	BaseStats      BaseStats              `bson:"baseStats" json:"baseStats"`
	Genus          string                 `bson:"genus" json:"genus"`
	Height         int                    `bson:"height" json:"height"`
	Weight         int                    `bson:"weight" json:"weight"`
	CaptureRate    int                    `bson:"captureRate" json:"captureRate"`
	BaseExperience int                    `bson:"baseExperience" json:"baseExperience"`
	Abilities      []string               `bson:"abilities" json:"abilities"`
	EffortValues   map[string]interface{} `bson:"effortValues" json:"effortValues"`
	EggGroups      []string               `bson:"eggGroups" json:"eggGroups"`
	Evolution      Evolution              `bson:"evolution" json:"evolution"`
}

type BaseStats struct {
	Hp             int `bson:"hp" json:"hp"`
	Attack         int `bson:"attack" json:"attack"`
	Defense        int `bson:"defense" json:"defense"`
	SpecialAttack  int `bson:"specialAttack" json:"specialAttack"`
	SpecialDefence int `bson:"specialDefense" json:"specialDefense"`
	Speed          int `bson:"speed" json:"speed"`
}

type Evolution struct {
	Level  int    `bson:"level" json:"level"`
	Name   string `bson:"name" json:"name"`
	Number int    `bson:"number" json:"number"`
}
