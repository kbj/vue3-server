package v1

import "vue3-server/service"

var (
	userService       = service.ServiceApp.System.UserService
	menuService       = service.ServiceApp.System.MenuService
	departmentService = service.ServiceApp.System.DepartmentService
	roleService       = service.ServiceApp.System.RoleService
)
