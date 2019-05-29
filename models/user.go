package models

import "time"

type User struct {
	AId uint64 `json:"a_id";gorm:"primary_key"`
	AName string `json:"a_name";gorm:"type:varchar(35);not null`
	ALoginTime time.Time `json:"a_login_time"`
	ATrueName string `json:"a_true_name";gorm:"type:varchar(35);not null`
	//Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
}
func (c *User) TableName() string {
	return "admin_users"
}
func GetById(id uint64) *User {
	u := &User{}
	DB.Where("a_id = ?", id).First(u)
	return u
}

func GetBySql(sql string,id uint64) (*User,error) {
	u := &User{}
	//DB.Exec(sql,id)
	//row := DB.Where("a_id = ?", id).Select("a_login_time, a_name").Row() // (*sql.Row)
	//DB.Raw(sql, id).Row().Scan(&u.aLoginTime,&u.AName)
	DB.Raw(sql, id).Scan(u)

	return u,nil
	/*rows, _ := DB.Model(&User{}).Where("a_id = ?", "id").Select("name, age, email").Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {

		rows.Scan(u)

	}*/

	// Raw SQL
	/*rows, _ := DB.Raw("select name, age, email from users where name = ?", "jinzhu").Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {

		rows.Scan(u)

	}*/
}