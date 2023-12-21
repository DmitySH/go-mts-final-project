package app

type Config struct {
	GRPC     `yaml:"grpc" mapstructure:",squash"`
	HTTP     `yaml:"http" mapstructure:",squash"`
	Swagger  `yaml:"swagger" mapstructure:",squash"`
	Mongo    `yaml:"mongo" mapstructure:",squash"`
	Kafka    `yaml:"kafka" mapstructure:",squash"`
	Driver   `yaml:"driver" mapstructure:",squash"`
	Location `yaml:"location" mapstructure:",squash"`
}

type GRPC struct {
	Addr string `yaml:"addr" mapstructure:"GRPC_ADDR"`
}

type HTTP struct {
	Addr string `yaml:"addr" mapstructure:"HTTP_ADDR"`
}

type Swagger struct {
	Path string `yaml:"path" mapstructure:"SWAGGER_PATH"`
}

type Mongo struct {
	URI        string `yaml:"uri" mapstructure:"MONGO_URI"`
	AuthSource string `yaml:"auth_source" mapstructure:"MONGO_AUTH_SOURCE"`
}

type Kafka struct {
	GroupID            string   `yaml:"group_id" mapstructure:"KAFKA_GROUP_ID"`
	Brokers            []string `yaml:"brokers" mapstructure:"KAFKA_BROKERS"`
	DriverProduceTopic string   `yaml:"driver_produce_topic" mapstructure:"KAFKA_DRIVER_PRODUCE_TOPIC"`
	DriverConsumeTopic string   `yaml:"driver_consume_topic" mapstructure:"KAFKA_DRIVER_CONSUME_TOPIC"`
}

type Driver struct {
	OfferRadius float64 `yaml:"offer_radius" mapstructure:"DRIVER_OFFER_RADIUS"`
}

type Location struct {
	Addr string `yaml:"addr" mapstructure:"LOCATION_ADDR"`
}
