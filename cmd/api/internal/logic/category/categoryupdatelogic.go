package category

import (
	"context"

	"mummy-storage/cmd/api/internal/svc"
	"mummy-storage/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryUpdateLogic) CategoryUpdate(req *types.CategoryUpdateRequest) (resp *types.CategoryUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
