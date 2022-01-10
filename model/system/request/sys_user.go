package request

type SysUserModel struct {
	Username string `json:"username" validate:"required,min=3,max=32" name:"用户名"`
	Password string `json:"password" validate:"required,min=8,max=32" name:"密码"`
}
