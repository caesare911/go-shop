package role

import (
	"context"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"
)

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	//插入数据返回id
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: int(lastInsertID)}, err
}

func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	//插入数据返回id
	_, err := dao.RoleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	return err
}
