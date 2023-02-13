package repository

import (
	"context"

	"github.com/itokun99/ms-go-boilerplate/core/config/log"

	"gorm.io/gorm"
)

type JsonXmlRepoStruct struct {
	db  *gorm.DB
	log *log.LogCustom
}

type JsonXmlRepoInterface interface {
	ParseJsonToXml(ctx context.Context) error
}

func NewJsonXmlRepository(db *gorm.DB, log *log.LogCustom) JsonXmlRepoInterface {
	return &JsonXmlRepoStruct{db: db, log: log}
}

func (b JsonXmlRepoStruct) ParseJsonToXml(ctx context.Context) error {

	log.NewElasticAPM(ctx, "Repo: JsonXmlRepo, Func: ParseJsonToXml", "Repo Handler")
	return nil
}
