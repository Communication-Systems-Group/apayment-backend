package models

import (
	"github.com/astaxie/beego/orm"
	"math/big"
)

type Request struct {
	Id             int64 `json:"id"`
	User           *User `orm:"rel(fk)" json:"user"`

	Address        string `json:"address"`
	Contributions  []*Contribution `orm:"-" json:"contributions"`
	Remark         string `orm:"-" json:"remark"`
	Created        *big.Int `orm:"-" json:"created"`
	Modified        *big.Int `orm:"-" json:"modified"`

	Inspector      *User `orm:"rel(fk);null" json:"inspector"`
	InspectionLacks []*InspectionLack `orm:"-" json:"inspectionLacks"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Request))
}