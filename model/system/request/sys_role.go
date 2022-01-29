package request

type SysRoleListModel struct {
	Offset   int         `json:"offset"`
	Size     int         `json:"size"`
	Name     string      `json:"name"`
	Intro    string      `json:"intro"`
	CreateAt interface{} `json:"createAt"`
}
