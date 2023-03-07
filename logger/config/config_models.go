package config

type ConfigNodes struct {
	ConfigDB DB         `mapstructure:"config_db"`
	Server   ConfServer `mapstructure:"server"`
}

type ConfServer struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	RpcPort  string `mapstructure:"rpc_port"`
	GrpcProt string `mapstructure:"grpc_port"`
}

type DB struct {
	MongoUrl   string `mapstructure:"mongo_url"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Collection string `mapstructure:"collection"`
}
