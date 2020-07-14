package services

import (
	"context"
	"go-common/app/infra/config/internel/dao"
	"go-common/app/infra/config/internel/models"
	rep "go-common/app/infra/config/reply"
	"go-common/app/infra/config/request"
	"go-common/library/ecode"
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
		return err
	}

	reply.ErrorCode = ecode.OK
	reply.Message = ""
	reply.Data = &rep.CreateConfigMapDto{
		ID:          config_map.ID,
		Name:        config_map.Name,
		Description: config_map.Description,
		CreatedAt:   config_map.CreatedAt,
		UpdatedAt:   config_map.UpdatedAt,
	}
	reply.Details = "创建配置组成功."

	return nil
}

func (service *ConfigMapService) GetConfigMaps(ctx context.Context, args *request.GetConfigMapsRequest, reply *rep.GetConfigMapsReply) error {
	var config_map_dto []*rep.GetConfigMapDto
	config_maps, err := dao.ConfigMaps()

	if err != nil {
		reply.ErrorCode = ecode.ServerException
		reply.Message = ""
		reply.Details = "获取配置组失败."
		reply.Data = nil
		return err
	}

	for _, v := range config_maps {
		config_map_dto = append(config_map_dto, &rep.GetConfigMapDto{
			Name:        v.Name,
			Description: v.Description,
			ID:          v.ID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	reply.ErrorCode = ecode.OK
	reply.Message = ""
	reply.Details = "您成功获取了所有的配置组"
	reply.Data = config_map_dto

	return nil
}

func (service *ConfigMapService) UpdateConfigMap(ctx context.Context, args *request.UpdateConfigMapRequest, reply *rep.UpdateConfigMapReply) error {
	config_map := &models.ConfigMap{}
	config_map.ID = args.ID
	config_map.Name = args.Name
	config_map.Description = args.Description

	if err := dao.UpdateConfigMap(config_map); err != nil {
		reply.ErrorCode = ecode.ServerException
		reply.Message = ""
		reply.Details = "更新配置组配置失败."
		return err
	}

	reply.ErrorCode = ecode.OK
	reply.Message = ""
	reply.Details = "更新配置组配置成功."

	return nil
}

func (service *ConfigMapService) DeleteConfigMap(ctx context.Context, args *request.DeleteConfigMapRequest, reply *rep.DeleteConfigMapReply) error {
	err := dao.DeleteConfigMap(args.ID)
	if err != nil {
		reply.ErrorCode = ecode.ServerException
		reply.Message = ""
		reply.Details = "删除配置表失败."

		return err
	}

	reply.ErrorCode = ecode.OK
	reply.Message = ""
	reply.Details = "成功删除配置表."

	return nil
}
