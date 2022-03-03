package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"user/model"
)

var (
	DB 					string
	DBHost 				string
	DBPort 				string
	DBUser 				string
	DBPassWord 			string
	DBName 				string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		DBUser, DBPassWord, DBHost, DBPort, DBName)
	model.Database(path)
}


func LoadMysqlData(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPassWord = file.Section("mysql").Key("DBPassWord").String()
	DBName = file.Section("mysql").Key("DBName").String()
}

