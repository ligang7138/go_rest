package models

import (
	"fmt"
	"go_rest/storage/mysql"
	"time"

	"go_rest/pkg/auth"
	"go_rest/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

type User struct {
	AId uint64 `json:"id";gorm:"primary_key;"`
	AName string `json:"name";gorm:"type:varchar(35);not null`
	APwd string `json:"passwd";gorm:"type:varchar(35);not null`
	ALoginTime time.Time `json:"login_time"`
	AStatus int8 `json:"status";gorm:"type:tinyint(1);not null;default 0`
	ARoles int8 `json:"roles";gorm:"type:tinyint(1);not null;default 0`
	AType int8  `json:"type";gorm:"type:tinyint(1);not null;default 0`
	AAddTime time.Time `json:"add_time";gorm:"type:varchar(35);not null`
	OpName string `json:"op_name";gorm:"type:varchar(35);not null`
	OpTime string `json:"op_time";gorm:"type:date(35);not null`

	ATrueName string `json:"a_true_name";gorm:"type:varchar(35);not null`
	//Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
}

var db = mysql.DB.Self
func (c *User) TableName() string {
	return "admin_users"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return db.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return db.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return db.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := db.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := db.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := db.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func GetById(id uint64) *User {
	u := &User{}
	db.Where("a_id = ?", id).First(u)
	return u
}

func GetBySql(sql string,id uint64) (*User,error) {
	u := &User{}
	//DB.Exec(sql,id)
	//row := DB.Where("a_id = ?", id).Select("a_login_time, a_name").Row() // (*sql.Row)
	db.Raw("" +
		"select u.a_name, i.a_true_name,u.a_login_time from admin_users u left join admin_user_info i on u.a_id = i.a_id where u.a_id = ?" +
		"", id).Row().Scan(u)

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
