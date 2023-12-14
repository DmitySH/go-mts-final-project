package app

type Config struct {
	GRPC    `yaml:"grpc"`
	HTTP    `yaml:"http"`
	Swagger `yaml:"swagger"`
	Mongo   `yaml:"mongo"`
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
