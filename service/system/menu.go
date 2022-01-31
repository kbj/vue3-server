package system

import (
	"vue3-server/common/global"
	"vue3-server/entity/system"
	"vue3-server/model/base"
	"vue3-server/model/system/response"
	"vue3-server/utils"
)

type MenuService struct{}

// RoleMenus 根据角色ID查询该角色拥有的菜单
func (*MenuService) RoleMenus(roleId uint) *base.ResponseEntity {
	// 查询角色拥有的菜单
	var params []uint
	global.Db.Raw(`
		select menu_id from t_role_menu where delete_at is null and role_id = ?
	`, roleId).Scan(&params)
	if len(params) < 1 {
		return utils.ResponseFail("暂时未查询到菜单信息！")
	}

	// 查询拥有的菜单信息
	var menus []system.Menu
	global.Db.Raw(`
		with recursive tree as (
			select t1.id, t1.name, t1.url, t1.icon, t1.sort, t1.parent_id, t1.type, t1.permission from t_menu t1 where t1.id in ? and t1.delete_at is null
			union all
			select t2.id, t2.name, t2.url, t2.icon, t2.sort, t2.parent_id, t2.type, t2.permission from t_menu t2 join tree t3 on t2.parent_id = t3.id where t2.delete_at is null
		)
		select * from tree
	`, params).Scan(&menus)

	// 递归组合成前端需要的格式
	var respMenu response.SysMenuRecursive
	respMenu = *recursiveMenus(&respMenu, &menus)
	return utils.ResponseSuccess(respMenu.Children)
}

// 递归查询出树的数组
func recursiveMenus(menuPoint *response.SysMenuRecursive, dbMenus *[]system.Menu) *response.SysMenuRecursive {
	menus := *dbMenus
	menuEntity := *menuPoint

	for i := range menus {
		if menus[i].ParentId == menuEntity.ID {
			// 将此时的menu封装到返回要的实体
			menu := response.SysMenuRecursive{
				ID:         menus[i].ID,
				Name:       menus[i].Name,
				Type:       menus[i].Type,
				Url:        menus[i].Url,
				Icon:       menus[i].Icon,
				Sort:       menus[i].Sort,
				ParentId:   menus[i].ParentId,
				Permission: menus[i].Permission,
			}
			menu = *recursiveMenus(&menu, dbMenus)

			// 判断menuEntity是否有children对象
			if menuEntity.Children == nil {
				var children []response.SysMenuRecursive
				children = append(children, menu)
				menuEntity.Children = &children
			} else {
				children := *menuEntity.Children
				children = append(children, menu)
				menuEntity.Children = &children
			}
		}
	}
	return &menuEntity
}

// GetMenuList 菜单列表
func (*MenuService) GetMenuList(offset int, size int) *base.ResponseEntity {
	// 查询拥有的菜单信息
	var menus []system.Menu
	var total int
	global.Db.Find(&menus)
	total = len(menus)

	// 递归组合成前端需要的格式
	var respMenu response.SysMenuRecursive
	respMenu = *recursiveMenus(&respMenu, &menus)

	var result = make(map[string]interface{})
	result["list"] = respMenu.Children
	result["totalCount"] = total
	return utils.ResponseSuccess(&result)
}

// DeleteMenu 删除菜单
func (*MenuService) DeleteMenu(id int) *base.ResponseEntity {
	global.Db.Delete(&system.Menu{}, id)
	return utils.ResponseSuccess("删除成功！")
}
