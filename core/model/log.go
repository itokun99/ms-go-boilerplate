package model

import "gorm.io/gorm"

type Logs struct {
	gorm.Model
	Level             string `json:"level" gorm:"column:level"`
	ServiceName       string `json:"service_name" gorm:"column:service_name"`
	ServiceEndpoint   string `json:"service_endpoint" gorm:"column:service_endpoint"`
	ServiceExternalId string `json:"service_external_id" gorm:"column:service_external_id"`
	JsonData          string `json:"json_data" gorm:"column:json_data"`
	XmlData           string `json:"xml_data" gorm:"column:xml_data"`
}

func (Logs) TableName() string {
	return "logs"
}
