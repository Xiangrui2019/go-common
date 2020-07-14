package services

import (
	"context"
	"go-common/app/infra/config/internel/dao"
	"go-common/app/infra/config/internel/models"
	rep "go-common/app/infra/config/reply"
	"go-common/app/infra/config/request"
	"go-common/library/ecode"

	"github.com/pkg/errors"
)

type ConfigMapService struct {
}

func (service *ConfigMapService) CreateConfigMap(ctx context.Context, args *request.CreateConfigMapRequest, reply *rep.CreateConfigMapReply) error {
	config_map, err := dao.CreateConfigMap(&models.ConfigMap{
		Name:        args.Name,
		Description: args.Description,
	})

	if err != nil {
		reply.ErrorCode = ecode.ServerException
		reply.Message = ""
		reply.Data = nil
		reply.Details = "服务器开小差了, 无法创建配置组."
		return errors.Wrap(err, "Create config map had an exception.")
	}

	reply.ErrorCode = ecode.OK
	reply.Message = ""
	reply.Data = &rep.CreateConfigMapDto{
		ID:          config_map.ID,
		Name:        config_map.Name,
		Description: config_map.Description,
		CreatedAt:   config_map.CreatedAt,
		UpdatedAt:   config_map.UpdatedAt,
		DeletedAt:   config_map.DeletedAt,
	}
	reply.Details = "创建配置组成功."

	return nil
}
