package entities

type Seeder struct {
	Tag      string `gorm:"not null; uniqueIndex" json:"tag"`
	Filename string `gorm:"not null" json:"filename"`
	User     string `gorm:"not null" json:"user"`
}

func (seed *Seeder) TableName() string {
	return "seeders"
}
