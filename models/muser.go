package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
}

type MUser struct {
}

//检查用户名是否可用
func (u *MUser) CheckUserNameValid(userName string)int{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=?", userName).Values(&maps)
	if err == nil && num <= 0 {
		return 0
	}
	return -2

}

//检查用户名是否存在
func (u *MUser) CheckUserNameExists(userName string)bool{
	if len(userName) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=?", userName).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}

//检查用户名密码是否正确
func (u *MUser) CheckUserAndPwd(userName string,userPwd string)bool{
	if len(userName) <= 0 || len(userPwd) <= 0{
		return false
	}
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_user_name FROM t_user WHERE F_user_name=? AND F_user_password = ?", userName,userPwd).Values(&maps)
	if err == nil && num > 0 {
		return true
	}
	return false
}


//检查用户密码
func (u *MUser) CheckUserPwdValid(userPwd string)int{
	if len(userPwd) > 0{
		return 0
	}
	return -3
}

//添加用户
func (u *MUser) AddUser(userName string,userPwd string)int{
	result := u.CheckUserNameValid(userName)
	if result == 0{
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		res, err := o.Raw("INSERT INTO t_user SET F_user_name = ?,F_user_password=?", userName,userPwd).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				result = 0
			}
		}
	}
	return result
}

//修改用户密码
func (u *MUser) ModifyUserPwd(userName string,userPwd string)int{
	result := -1
	res := u.CheckUserNameExists(userName)
	if res {
		result = u.CheckUserPwdValid(userPwd)
	}
	if result == 0{
		result = -1
		//写入数据库
		o := orm.NewOrm()
		res, err := o.Raw("UPDATE t_user SET F_user_password=? WHERE F_user_name=?",userPwd,userName).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num >0{
				result = 0
			}
		}
	}
	return result
}

//获取用户的密码
func (u *MUser) GetUserPwd(userName string)string{
	pwd := ""
	if len(userName) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT F_user_password FROM t_user WHERE F_user_name=? LIMIT 1", userName).Values(&maps)
		if err == nil && num > 0 {
			pwd = maps[0]["F_user_password"].(string)
		}
	}
	return pwd

}