package category

import (
	"context"

	"mummy-storage/cmd/api/internal/errcode"
	"mummy-storage/cmd/api/internal/svc"
	"mummy-storage/cmd/api/internal/types"
	"mummy-storage/internal/crtl/categoryctrl"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryCreateLogic {
	return &CategoryCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryCreateLogic) CategoryCreate(req *types.CategoryCreateRequest) (resp *types.CategoryCreateResponse, err error) {
	ctrl := categoryctrl.NewCategoryCtrl(l.svcCtx.Env)
	err = ctrl.Create(l.ctx, req.Name, req.Name)
	if err != nil {
		return nil, errcode.ErrService
	}
	return
}
