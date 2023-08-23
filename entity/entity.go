package entity

type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Villain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CriminalReport struct {
	ID          int    `json:"id"`
	HeroID      int    `json:"hero_id"`
	VillainID   int    `json:"villain_id"`
	Description string `json:"description"`
	Occurrence  string `json:"occurrence"`
}
