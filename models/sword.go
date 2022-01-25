package models

type Sword struct {
	ID				uint		`json:"id" gorm:"primary_key"`
	Name			string		`json:"name"`
	Image			string		`json:"image"`
	Price			float32		`json:"price"`
	Inches			int			`json:"inches"`
	Ounces			int			`json:"ounces"`
	Mats			string		`json:"mats"`
	Description		string		`json:"description"`
}