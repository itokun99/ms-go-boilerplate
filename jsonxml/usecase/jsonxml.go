package usecase

import (
	"context"

	"github.com/itokun99/ms-go-boilerplate/core/config/log"
	"github.com/itokun99/ms-go-boilerplate/core/model"
	"github.com/itokun99/ms-go-boilerplate/jsonxml/repository"
)

type JsonXmlUseCaseInterface interface {
	JsonToXml(requestBody interface{}, requestHeader model.BaseHeaderRequest, ctx context.Context) (response interface{}, err error)
}

type JsonXmlUseCase struct {
	log         *log.LogCustom
	jsonXmlRepo repository.JsonXmlRepoInterface
	config      model.ServerConfig
}

func NewJsonXmlUseCase(jsonXmlRepo repository.JsonXmlRepoInterface, log *log.LogCustom, config model.ServerConfig) JsonXmlUseCaseInterface {
	return &JsonXmlUseCase{
		jsonXmlRepo: jsonXmlRepo, log: log, config: config}
}

func (outgoingUc JsonXmlUseCase) JsonToXml(requestBody interface{}, requestHeader model.BaseHeaderRequest, ctx context.Context) (response interface{}, err error) {
	return response, nil
}
