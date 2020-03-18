package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

var (
	PersonList map[string]*Person
)

func init() {
	orm.RegisterModel(new(Person))
}

type Person struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func AddPerson(p Person) string {

	// register model

	// create table
	// orm.RunSyncdb("default", false, true)

	o := orm.NewOrm()

	// insert
	id, err := o.Insert(&p)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	return strconv.Itoa(int(id))
}

func UpdatePerson(uid string, pp *Person) (p *Person, err error) {
	o := orm.NewOrm()
	err := o.Read(&pp)
	if err != nil {
	}
}
