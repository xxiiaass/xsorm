package xsorm

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Model *gorm.DB

var models map[string]map[string]*gorm.DB

type XConfig struct {
	MaxIdleConns int
	MaxOpenConns int
	Config       mysql.Config
	Debug        bool
}

var xconfigs []XConfig

type Log interface {
	Print(v ...interface{})
}

type defaultLog struct{}

func (defaultLog) Print(v ...interface{}) {}

var log Log

const (
	// 数据库的读写性质
	ReadOnly = "readonly"
	Write    = "write"
	Proxy    = "proxy"
)

var DefaultCon = "default" // 默认使用连接的key名

func init() {
	log = defaultLog{} // 初始化日志接口为无操作
}

func AddConnect(config XConfig) {
	xconfigs = append(xconfigs, config)
	if len(xconfigs) == 1 {
		DefaultCon = xconfigs[0].Config.DBName
	}
}

func SetLogger(logger Log) {
	log = logger
}

func Init() {
	models = make(map[string]map[string]*gorm.DB)

	for _, item := range xconfigs {
		model, err := gorm.Open("mysql", item.Config.FormatDSN())
		if err != nil {
			panic(err)
		}
		model.DB().SetMaxIdleConns(item.MaxIdleConns)
		model.DB().SetMaxOpenConns(item.MaxOpenConns)
		model.LogMode(item.Debug)
		model.SetLogger(log)

		if _, ok := models[item.Config.DBName]; !ok {
			models[item.Config.DBName] = make(map[string]*gorm.DB)
		}
		models[item.Config.DBName][Proxy] = model
	}
	for _, v := range models {
		if _, ok := v[Write]; !ok {
			v[Write] = v[Proxy]
		}
	}
	Model = models[DefaultCon][Proxy]
}
