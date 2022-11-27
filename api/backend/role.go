package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleCreateReq struct {
	g.Meta `path:"/backend/role" tags:"Role" method:"post" summary:"添加角色"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc"    v:"required#角色描述不能为空" dc:"角色描述"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role" tags:"Role" method:"put" summary:"修改角色"`
	Id     uint   `json:"id" v:"required#角色id不能为空" dc:"角色id"`
	Name   string `json:"name"  dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleCreateRes struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateRes struct {
	RoleId uint `json:"role_id"`
}
