package service

import (
	"fmt"
	"go_rest/models"
	"strconv"
)

func GetUserName(id string) (string,error) {
	//models.DB.Table("qy_partner_users").Find(id)
	userId,_ := strconv.ParseUint(id, 10 , 64)
	user := models.GetById(userId)
	return user.AName,nil
}

func GetUserBySql(id string) (string,error) {
	//models.DB.Table("qy_partner_users").Find(id)
	userId,_ := strconv.ParseUint(id, 10 , 64)
	sql := "select u.a_name, i.a_true_name,u.a_login_time from admin_users u left join admin_user_info i on u.a_id = i.a_id where u.a_id = ? limit 1"
	user,_ := models.GetBySql(sql,userId)
	fmt.Printf(user.ALoginTime.String())
	return user.AName + user.ATrueName + user.ALoginTime.Format(models.Time_Format),nil
}