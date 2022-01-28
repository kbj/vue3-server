package response

type SysMenuRecursive struct {
	ID         uint                `json:"id,omitempty"`
	Name       string              `json:"name,omitempty"`
	Type       uint8               `json:"type,omitempty"`
	Url        string              `json:"url"`
	Icon       string              `json:"icon,omitempty"`
	Sort       uint                `json:"sort"`
	ParentId   uint                `json:"parentId,omitempty"`
	Permission string              `json:"permission,omitempty"`
	Children   *[]SysMenuRecursive `json:"children,omitempty"`
}
