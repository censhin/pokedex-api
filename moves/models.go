package moves

type Moves struct{
	Moves	[]Move `json:"moves"`
}

type Move struct {
	Name       string   `json:"name"`
	Number     int      `json:"number"`
	Type       string   `json:"type"`
	Category   string   `json:"category"`
	Contest    string   `json:"contest"`
	Pp         int      `json:"pp"`
	Power      int      `json:"power"`
	Accuracy   int      `json:"accuracy"`
	Generation string   `json:"generation"`
	Tm         []string `json:"tm"`
	Hm         []string `json:"hm"`
	Tutor      bool     `json:"tutor"`
}
