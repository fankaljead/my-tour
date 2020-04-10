package models

import (
	"github.com/astaxie/beego/orm"
)

type Hobby struct {
	Id          int64
	Title       string
	Description string
}

type UserHobby struct {
	Id      int64
	UserId  int64
	HobbyId int64
}

func init() {

	orm.RegisterModel(new(UserHobby), new(Hobby))
	// create table
	// orm.RunSyncdb("default", false, true)

	AddATable("hobby")
	AddATable("user_hobby")
}
func AddHobby(title string) int64 {
	o := orm.NewOrm()
	var hobby Hobby
	hobby.Title = title
	// o.ReadOrCreate(&hobby, "id")
	// fmt.Printf("hobby title is :   %s\n\n", title)
	// fmt.Printf("hobby is :   %v\n\n", hobby)
	if created, id, err := o.ReadOrCreate(&hobby, "Title"); err == nil {
		if created {
			// fmt.Println("New Insert an object. Id:", id)
			return id
		} else {
			// fmt.Println("Get an object. Id:", id)
			return id
		}
	}
	return -1
}

func AddUserHobby(user_id int64, hobby_id int64) int64 {
	o := orm.NewOrm()
	var user_hobby UserHobby
	user_hobby.UserId = user_id
	user_hobby.HobbyId = hobby_id

	id, err := o.Insert(&user_hobby)
	if err == nil {
		return id
	}
	return -1
}

func DeleteUserHobby(user_hobby_id int64) (num int64, err error) {
	o := orm.NewOrm()
	if num, err = o.Delete(&UserHobby{Id: user_hobby_id}); err == nil {
		return num, nil
	}
	return
}

func GetUserHobbies(user_id string) map[string]interface{} {
	user_hobbies := make(map[string]interface{})

	user_hobbies["user_id"] = user_id

	var res []struct {
		HobbyId     int64  `json:"hobby_id"`
		UserHobbyId int64  `json:"user_hobby_id"`
		Title       string `json:"title"`
	}

	// res := make(orm.Params)
	o := orm.NewOrm()
	// nums, err := o.Raw("SELECT hobby.title as title, hobby.id as hobbyId,user_hobby.id as userHobbyId FROM user_hobby left join hobby on user_hobby.hobby_id=hobby.id where user_hobby.user_id="+user_id).RowsToStruct(res, "userHobbyId", "title")
	nums, err := o.Raw("SELECT hobby.title as title, hobby.id as hobby_id,user_hobby.id as user_hobby_id FROM user_hobby left join hobby on user_hobby.hobby_id=hobby.id where user_hobby.user_id=" + user_id).QueryRows(&res)

	if err == nil {
		user_hobbies["hobbies"] = res
	}

	user_hobbies["number"] = nums

	return user_hobbies
}
