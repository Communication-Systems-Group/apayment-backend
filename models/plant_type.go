package models

import "github.com/astaxie/beego/orm"

type PlantType struct {
	Id   int64  `json:"id"`
	Code uint8  `json:"code"`
	Name string `json:"name"`
}

func init() {
	// Register model
	orm.RegisterModel(new(PlantType))
}
