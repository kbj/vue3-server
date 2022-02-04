package system

import (
	"reflect"
	"time"
	"vue3-server/common/global"
	base2 "vue3-server/entity/base"
	"vue3-server/entity/system"
	"vue3-server/model/base"
	"vue3-server/model/system/request"
	"vue3-server/model/system/response"
	"vue3-server/utils"
)

type RoleService struct{}

// GetRoleList 查询角色列表
func (roleService *RoleService) GetRoleList(param *request.SysRoleListModel) *base.ResponseEntity {
	db := global.Db.Model(&system.Role{}).Offset(param.Offset).Limit(param.Size)
	if param.Name != "" {
		db.Where("name like ?", "%"+param.Name+"%")
	}
	if param.Intro != "" {
		db.Where("intro like ?", "%"+param.Intro+"%")
	}
	if param.CreateAt != nil {
		switch reflect.TypeOf(param.CreateAt).Kind() {
		case reflect.Array, reflect.Slice:
			// 时间数组
			values := reflect.ValueOf(param.CreateAt)
			startString := values.Index(0).Interface().(string)
			endString := values.Index(1).Interface().(string)
			start, _ := time.Parse("2006-01-02T15:04:05.000Z", startString)
			end, _ := time.Parse("2006-01-02T15:04:05.000Z", endString)
			db.Where("create_at between ? and ?", start, end)
		}
	}

	var lists []response.SysRoleListModel
	var total int64
	db.Count(&total)
	db.Find(&lists)

	// 添加菜单信息
	var menuService MenuService
	for i := range lists {
		menus := *menuService.RoleMenus(lists[i].ID)
		if menus.Code == 0 {
			menu := menus.Data.(*[]response.SysMenuRecursive)
			lists[i].MenuList = menu
		}
	}

	var result = make(map[string]interface{})
	result["list"] = lists
	result["totalCount"] = total
	return utils.ResponseSuccess(&result)
}

// CreateRole 创建角色
func (*RoleService) CreateRole(param *request.SysRoleMenuModel) *base.ResponseEntity {
	// 保存角色信息
	roleInfo := system.Role{
		Model: base2.Model{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
		Name:  param.Name,
		Intro: param.Intro,
	}
	result := global.Db.Save(&roleInfo)
	if result.Error != nil {
		return utils.ResponseFail("保存失败！")
	}

	// 保存角色菜单
	menus := *param.MenuList
	for i := range menus {
		roleMenu := system.RoleMenu{
			Model: base2.Model{
				CreateAt: time.Now(),
				UpdateAt: time.Now(),
			},
			RoleId: roleInfo.ID,
			MenuId: menus[i],
		}
		global.Db.Save(&roleMenu)
	}

	return utils.ResponseSuccess("保存成功！")
}

// UpdateRole 更新角色
func (*RoleService) UpdateRole(param *request.SysRoleMenuModel) *base.ResponseEntity {
	role := system.Role{
		Model: base2.Model{
			ID: param.Id,
		},
		Name:  param.Name,
		Intro: param.Intro,
	}
	global.Db.Model(&role).Updates(role)

	if param.MenuList != nil {
		// 更改菜单
		global.Db.Where("role_id = ?", param.Id).Delete(&system.RoleMenu{})
		menus := *param.MenuList
		if len(menus) > 0 {
			for i := range menus {
				roleMenu := system.RoleMenu{
					Model: base2.Model{
						CreateAt: time.Now(),
						UpdateAt: time.Now(),
					},
					RoleId: role.ID,
					MenuId: menus[i],
				}
				global.Db.Save(&roleMenu)
			}
		}
	}

	return utils.ResponseSuccess("保存成功！")
}

// DeleteRole 删除角色
func (*RoleService) DeleteRole(id uint) *base.ResponseEntity {
	// 删除菜单
	global.Db.Where("role_id = ?", id).Delete(&system.RoleMenu{})
	global.Db.Delete(&system.Role{}, id)
	return utils.ResponseSuccess("删除成功！")
}
