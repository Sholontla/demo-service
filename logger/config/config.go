package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf = JsonConfigNodes()

var vp *viper.Viper

func JsonConfigNodes() ConfigNodes {
	vp = viper.New()
	var config ConfigNodes

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config_files")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("Error while reading config.json")
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println("Error Unmarrsall", err)
	}
	return config
}

func DBConfig() (string, string, string) {
	var (
		data     = JsonConfigNodes()
		mongoUrl = data.ConfigDB.MongoUrl
		user     = data.ConfigDB.MongoUrl
		pass     = data.ConfigDB.MongoUrl
	)
	return mongoUrl, user, pass
}

func DBCollection() string {

	data := JsonConfigNodes()
	collection := data.ConfigDB.Collection

	return collection
}

func ServerConfig() (string, string) {
	data := JsonConfigNodes()
	port := data.Server.Port
	rpcPort := data.Server.RpcPort

	return port, rpcPort
}

func GrpcConf() string {
	data := JsonConfigNodes()

	port := data.Server.GrpcProt

	return port
}
