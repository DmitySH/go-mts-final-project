package app

type Config struct {
	GRPC    `yaml:"grpc"`
	HTTP    `yaml:"http"`
	Swagger `yaml:"swagger"`
	Mongo   `yaml:"mongo"`
	Kafka   `yaml:"kafka"`
}

type GRPC struct {
	Addr string `yaml:"addr"`
}

type HTTP struct {
	Addr string `yaml:"addr"`
}

type Swagger struct {
	Path string `yaml:"path"`
}

type Mongo struct {
	URI        string `yaml:"uri"`
	AuthSource string `yaml:"auth_source"`
}

type Kafka struct {
	GroupID            string   `yaml:"group_id"`
	Brokers            []string `yaml:"brokers"`
	DriverProduceTopic string   `yaml:"driver_produce_topic"`
	DriverConsumeTopic string   `yaml:"driver_consume_topic"`
}
