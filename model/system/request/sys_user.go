package request

type SysUserModel struct {
	Name     string `json:"name" validate:"required,min=3,max=32" name:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=32" name:"密码"`
}

type SysUserListModel struct {
	Id        string      `json:"id"`
	Offset    int         `json:"offset"`
	Size      int         `json:"size"`
	Name      *string     `json:"name"`
	Realname  *string     `json:"realname"`
	Cellphone *string     `json:"cellphone"`
	Enable    interface{} `json:"enable"`
	CreateAt  interface{} `json:"createAt"`
}

type SysUser struct {
	Id           uint   `json:"id"`
	Cellphone    string `json:"cellphone"`
	DepartmentId uint   `json:"departmentId"`
	Name         string `json:"name"`
	Realname     string `json:"realname"`
	RoleId       uint   `json:"roleId"`
}
