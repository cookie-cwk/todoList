package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"mq-server/model"
)

var (
	DB 					string
	DBHost				string
	DBPort				string
	DBUser				string
	DBPassword			string
	DBName				string

	RabbitMQ 			string
	RabbitMQUser		string
	RabbitMQPassword	string
	RabbitMQHost		string
	RabbitMQPort		string
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：%v",err)
	}
	LoadMySQL(file)
	pathMySQL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)
	model.Database(pathMySQL)
	LoadRabbitMQ(file)
	pathRabbitMQ := fmt.Sprintf("%s://%s:%s@%s:%s/",
		RabbitMQ, RabbitMQUser, RabbitMQPassword, RabbitMQHost, RabbitMQPort)
	model.RabbitMQ(pathRabbitMQ)
}

func LoadMySQL(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPassword = file.Section("mysql").Key("DBPassword").String()
	DBName = file.Section("mysql").Key("DBName").String()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassword = file.Section("rabbitmq").Key("RabbitMQPassword").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}