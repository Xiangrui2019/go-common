package services

import (
	"context"
	"go-common/app/infra/config/internel/dao"
	"go-common/app/infra/config/internel/models"
	"go-common/app/infra/config/reply"
	"go-common/app/infra/config/request"
	"go-common/library/ecode"
)

type ConfigEntryService struct {
}

func (service *ConfigEntryService) CreateConfigEntry(ctx context.Context, args *request.CreateConfigEntryRequest, rep *reply.CreateConfigEntryReply) error {
	config_entry, err := dao.CreateConfigEntry(&models.ConfigEntry{
		Key:         args.Key,
		Value:       args.Value,
		ConfigMapId: args.ConfigMapId,
	})

	if err != nil {
		rep.ErrorCode = ecode.ServerException
		rep.Message = "Create config_entry had an exception."
		rep.Details = "创建配置项时出现一个错误."
		rep.Data = nil
		return err
	}

	rep.ErrorCode = ecode.OK
	rep.Message = "Create OK."
	rep.Details = "创建成功."
	rep.Data = &reply.CreateConfigEntryDto{
		ID:        config_entry.ID,
		Key:       config_entry.Key,
		Value:     config_entry.Value,
		CreatedAt: config_entry.CreatedAt,
		UpdatedAt: config_entry.UpdatedAt,
	}

	return nil
}

func (service *ConfigEntryService) GetConfigEntrysByMapId(ctx context.Context, args *request.GetConfigEntrysByMapIdRequest, rep *reply.GetConfigEntrysByMapIdReply) error {
	var config_entrys_dto []*reply.GetConfigEntryDto

	entrys, err := dao.ConfigEntrysByMapId(args.ConfigMapId)
	if err != nil {
		rep.ErrorCode = ecode.ServerException
		rep.Message = "Getting ConfigEntrys had an exception."
		rep.Details = "获取配置项失败."
		rep.Datas = nil
		return err
	}

	for _, v := range entrys {
		config_entrys_dto = append(config_entrys_dto, &reply.GetConfigEntryDto{
			ID:        v.ID,
			Key:       v.Key,
			Value:     v.Value,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	rep.ErrorCode = ecode.OK
	rep.Message = "OK."
	rep.Details = "获取配置项成功."
	rep.Datas = config_entrys_dto

	return nil
}

/**
func (service *ConfigEntryService) UpdateConfigEntry(ctx context.Context, args *request.UpdateConfigEntryRequest, rep *reply.UpdateConfigEntryRequest) error {
}

func (service *ConfigEntryService) DeleteConfigEntry(ctx context.Context, args *request.DeleteConfigEntryRequest, rep *reply.DeleteConfigEntryReply) error {
}
**/
