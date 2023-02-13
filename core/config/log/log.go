package log

import (
	"fmt"
	"sync"
	"time"

	"database/sql"

	"github.com/itokun99/ms-go-boilerplate/core/model"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"gorm.io/gorm"

	pglogrus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
)

type LogCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
	Db     *gorm.DB
}

type iAm struct {
	Name string
	Host string
	Port string
}

var instance *LogCustom
var once sync.Once

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom(configServer model.ServerConfig, db *gorm.DB) *LogCustom {
	var log *logrus.Logger

	configElstc := configServer.ElasticConfig
	sqlDb, err := db.DB()

	// setup logrus
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// setup elasticsearch logging
	client, err := elastic.NewClient(elastic.SetURL(
		fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(configElstc.User, configElstc.Password))
	if err != nil {
		selfLogError(err, "core/config/log: elastic client", log)
	} else {
		hook, err := elogrus.NewAsyncElasticHook(
			client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
		if err != nil {
			selfLogError(err, "core/config/log: elastic client", log)
		}
		log.Hooks.Add(hook)
	}

	// setup db logger
	hook2 := pglogrus.NewAsyncHook(sqlDb, map[string]interface{}{})
	hook2.InsertFunc = func(sqlDb *sql.Tx, entry *logrus.Entry) error {
		level := entry.Level.String()

		if level == "info" {
			level = "success"
		}

		err = db.Debug().Exec("INSERT INTO logs(level, service_name, service_endpoint, service_external_id, json_data, xml_data, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?);", level, entry.Data["service_name"], entry.Data["service_endpoint"], entry.Data["service_external_id"], entry.Data["json_data"], entry.Data["xml_data"], time.Now(), time.Now()).Error

		if err != nil {
			selfLogError(err, "core/config/log: error execute db log", log)

			err := sqlDb.Rollback()
			if err != nil {
				selfLogError(err, "core/config/log: error roleback db log", log)
				return err
			}
		}

		return err
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: configServer.Name,
				Host: configServer.Host,
				Port: configServer.Port,
			},
		}
	})
	return instance
}

func (l *LogCustom) Success(description, serviceName string, serviceEndpoint string, serviceExId string, jsonData interface{}, xmlData interface{}) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":              l.WhoAmI,
		"service_name":        serviceName,
		"service_endpoint":    serviceEndpoint,
		"service_external_id": serviceExId,
		"json_data":           jsonData,
		"xml_data":            xmlData,
	}).Info("SUCCESS")
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Info(description string, data ...interface{}) {
	l.Logrus.WithFields(logrus.Fields{
		"whoami":  l.WhoAmI,
		"message": data,
	}).Info(description)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Error(err error, description string, serviceName string, serviceEndpoint string, serviceExId string, jsonData interface{}, xmlData interface{}) {

	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":              l.WhoAmI,
		"error_cause":         stFormat,
		"error_message":       err.Error(),
		"service_name":        serviceName,
		"service_endpoint":    serviceEndpoint,
		"service_external_id": serviceExId,
		"json_data":           jsonData,
		"xml_data":            xmlData,
	}).Error(description)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Fatal(err error, description string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

// for description please use format for example
// "usecase: sync data"
func SelfLogFatal(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}
