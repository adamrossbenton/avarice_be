package models

type Sword struct {
	ID				uint		`json:"id" gorm:"primary_key"`
	Name			string		`json:"name"`
	Image			string		`json:"image"`
	Price			float32		`json:"price,string"`
	Inches			int			`json:"inches,string"`
	Ounces			int			`json:"ounces,string"`
	Mats			string		`json:"mats"`
	Description		string		`json:"description"`
}