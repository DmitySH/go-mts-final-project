package app

type Config struct {
	GRPC     `yaml:"grpc"`
	HTTP     `yaml:"http"`
	Swagger  `yaml:"swagger"`
	Postgres `yaml:"postgres"`
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

type Postgres struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DBName string `yaml:"db_name"`
	SSL    string `yaml:"ssl"`
}
