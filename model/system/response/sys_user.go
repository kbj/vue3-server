package response

import "vue3-server/entity/system"

type SysUserModel struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Realname   string            `json:"realname"`
	Cellphone  string            `json:"cellphone"`
	Enable     uint              `json:"enable"`
	Role       system.Role       `json:"role"`
	Department system.Department `json:"department"`
}
