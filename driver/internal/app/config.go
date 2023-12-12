package app

type Config struct {
	GRPC    `yaml:"grpc"`
	HTTP    `yaml:"http"`
	Swagger `yaml:"swagger"`
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
