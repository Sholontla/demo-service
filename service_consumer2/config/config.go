package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf = JsonConfigNodes()

var vp *viper.Viper

func JsonConfigNodes() ConfigNodes {
	//intial := JsonInitialConfig()
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

func DBConfig() (string, string, string, string, string, string) {
	data := JsonConfigNodes()
	postgresqlHost := data.ConfigDB.PostgreHost
	sslmode := data.ConfigDB.PostgreSSLMode
	postgresqlUsername := data.ConfigDB.PostgreUser
	postgresqlPassword := data.ConfigDB.PostgrePassword
	postgresqlPort := data.ConfigDB.PostgrePort
	postgresqlSchema := data.ConfigDB.PostgreSchema

	return sslmode, postgresqlUsername, postgresqlPassword, postgresqlHost, postgresqlPort, postgresqlSchema
}

func GRPCConfig() (string, string, bool) {
	data := JsonConfigNodes()
	grpc_host := data.GRPCServer.GRPCHost
	grpc_port := data.GRPCServer.GRPCPort
	grpc_tls := data.GRPCServer.GRPCTLS

	return grpc_host, grpc_port, grpc_tls
}

func GRPCPathsConfig() (string, string) {
	data := JsonConfigNodes()
	cert_File := data.GRPCPaths.CertFile
	key_File := data.GRPCPaths.CertFile

	return cert_File, key_File
}
