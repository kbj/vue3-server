package system

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/common/global"
	"vue3-server/entity/system"
	"vue3-server/model/base"
	"vue3-server/utils"
)

type MenuService struct{}

// RoleMenus 根据角色ID查询该角色拥有的菜单
func (menu *MenuService) RoleMenus(ctx *fiber.Ctx, roleId uint) *base.ResponseEntity {
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
			select t1.id, t1.name, t1.url, t1.icon, t1.sort, t1.parent_id from t_menu t1 where (t1.parent_id in ? or t1.id in ?) and t1.delete_at is null
			union all
			select t2.id, t2.name, t2.url, t2.icon, t2.sort, t2.parent_id from t_menu t2 join tree t3 on t2.parent_id = t3.id where t2.delete_at is null
		)
		select * from tree
	`, params, params).Scan(&menus)
	// 递归组合成前端需要的格式
	return utils.ResponseSuccess(&menus)
}
