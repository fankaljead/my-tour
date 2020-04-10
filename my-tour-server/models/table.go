package models

import "github.com/astaxie/beego/orm"

type DatabaseTable struct {
	Id        int64  `json:"id"`
	TableName string `json:"table_name"`
}

func init() {
	orm.RegisterModel(new(DatabaseTable))
	AddATable("database_table")
	AutoCreateTabel()
}

func AddTable(table_name string) (int64, error) {
	o := orm.NewOrm()

	database_table := DatabaseTable{
		TableName: table_name,
	}

	return o.InsertOrUpdate(&database_table)
}

func AutoCreateTabel() (successNums int64, err error) {
	var tables []DatabaseTable

	for _, name := range table_names {
		table := DatabaseTable{
			TableName: name,
		}
		tables = append(tables, table)
	}

	o := orm.NewOrm()

	successNums, err = o.InsertMulti(len(tables), tables)
	return
}

var table_names []string

func AddATable(table_name string) {
	table_names = append(table_names, table_name)
}
