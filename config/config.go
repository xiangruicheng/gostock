package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type ConfigType struct {
	App struct {
		LogPath string `yaml:"log_path"`
	} `yaml:"app"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"mysql"`
	DDL struct {
		Path   string   `yaml:"path"`
		Tables []string `yaml:"tables"`
	} `yaml:"ddl"`
	Xueqiu struct {
		InitNum int64  `yaml:"init-num"`
		IncrNum int64  `yaml:"incr-num"`
		Cookie  string `yaml:"cookie"`
	} `yaml:"xueqiu"`
	Etf []struct {
		Name   string `yaml:"name"`
		Market string `yaml:"market"`
		Code   string `yaml:"code"`
	} `yaml:"etf"`
	Index []struct {
		Name   string `yaml:"name"`
		Market string `yaml:"market"`
		Code   string `yaml:"code"`
	} `yaml:"index"`
}

var Data ConfigType

// InitConfig 初始化配置
func InitConfig() {
	filePath := "config/config.yaml"
	// 读取YAML文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("datainit config error %s: %v", filePath, err)
	}
	// 解析YAML内容到Config结构体
	err = yaml.Unmarshal(data, &Data)
	if err != nil {
		log.Fatalf("datainit config Error parsing YAML: %v", err)
	}
	log.Println("datainit config success")
}
