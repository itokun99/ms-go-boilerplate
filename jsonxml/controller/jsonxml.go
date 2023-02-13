package controller

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"

	"github.com/itokun99/ms-go-boilerplate/core/config/log"
	"github.com/itokun99/ms-go-boilerplate/core/model"
	"github.com/itokun99/ms-go-boilerplate/jsonxml/usecase"
	dynamicstruct "github.com/ompluscator/dynamic-struct"

	"github.com/gin-gonic/gin"
)

type JsonXmlControllerStruct struct {
	log     *log.LogCustom
	usecase usecase.JsonXmlUseCaseInterface
	config  model.ServerConfig
}

func NewJsonXmlController(router *gin.RouterGroup, log *log.LogCustom, uc usecase.JsonXmlUseCaseInterface, config model.ServerConfig) {
	handler := JsonXmlControllerStruct{
		log:     log,
		usecase: uc,
		config:  config,
	}

	router.POST("/json-xml", handler.JsonToXml)
	router.POST("/xml-json", handler.XmlToJson)
}

type exampleXml struct {
	XmlName xml.Name `xml:"xml"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}

func (p *JsonXmlControllerStruct) JsonToXml(c *gin.Context) {
	var reqTemp map[string]interface{}

	instance := dynamicstruct.NewStruct()

	log.NewElasticAPM(c, "Controller JsonToXml", "request")

	c.ShouldBindJSON(&reqTemp)

	// fmt.Println("reqTemp ===>", reqTemp)

	for key1, value1 := range reqTemp {
		// fmt.Println("i ==>", key1, "s ==>", reflect.ValueOf(value1).Kind())

		if reflect.ValueOf(value1).Kind() == reflect.Map {
			// for key2, value2 := range value1.(map[string]interface{}) {
			// 	if reflect.ValueOf(value2).Kind() == reflect.Map {

			// 	} else {

			// 	}
			// }
		} else {
			fmt.Println("i ==>", key1, "s ==>", reflect.ValueOf(value1).Kind())
			instance.AddField(key1, value1, `json:"`+key1+`"`)
		}
	}

	ins := instance.Build().New()

	fmt.Println("instance:", ins)
	c.JSON(http.StatusOK, ins)
}

func (p *JsonXmlControllerStruct) XmlToJson(c *gin.Context) {

	log.NewElasticAPM(c, "Controller XmlToJson", "request")
	c.JSON(http.StatusOK, gin.H{
		"code":    "00",
		"message": "Response json successfully",
	})
}
