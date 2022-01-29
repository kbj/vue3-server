package system

import (
	"vue3-server/common/global"
	"vue3-server/entity/system"
	"vue3-server/model/base"
	"vue3-server/model/system/request"
	"vue3-server/utils"
)

type DepartmentService struct{}

// GetDepartmentList 部门列表
func (departmentService *DepartmentService) GetDepartmentList(param *request.SysUserListModel) *base.ResponseEntity {
	db := global.Db.Model(&system.Department{}).Offset(param.Offset).Limit(param.Size)
	var lists []system.Department
	var total int64

	db.Count(&total)
	db.Find(&lists)

	var result = make(map[string]interface{})
	result["list"] = lists
	result["totalCount"] = total
	return utils.ResponseSuccess(&result)
}
