package services

type Movie struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Year            int    `json:"year"`
	Rating          int    `json:"rating"`
	ScoutbaseRating string `json:"scoutbase_rating"`
	CreatedAt       string `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt       string `gorm:"column:updatedAt" json:"updated_at"`
}

type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Country   string `json:"country"`
	CreatedAt string `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt string `gorm:"column:updatedAt" json:"updated_at"`
}

type Director struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Country   string `json:"country"`
	CreatedAt string `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt string `gorm:"column:updatedAt" json:"updated_at"`
}

func (Actor) TableName() string {
	return "Actor"
}

func (Director) TableName() string {
	return "Director"
}

func (Movie) TableName() string {
	return "Movie"
}
