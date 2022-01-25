package models

type Sword struct {
	ID				uint		`json:"id" gorm:"primary_key"`
	Name			string		`json:"name"`
	Price			float32		`json:"price"`
	Inches			int			`json:"Inches"`
	Ounces			int			`json:"ounces"`
	Mats			string		`json:"mats"`
	Description		string		`json:"description"`
}