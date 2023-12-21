package app

type Config struct {
	GRPC     `yaml:"grpc" mapstructure:",squash"`
	HTTP     `yaml:"http" mapstructure:",squash"`
	Swagger  `yaml:"swagger" mapstructure:",squash"`
	Postgres `yaml:"postgres" mapstructure:",squash"`
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

type Postgres struct {
	Host   string `yaml:"host" mapstructure:"POSTGRES_HOST"`
	Port   string `yaml:"port" mapstructure:"POSTGRES_PORT"`
	DBName string `yaml:"db_name" mapstructure:"POSTGRES_DB_NAME"`
	SSL    string `yaml:"ssl" mapstructure:"POSTGRES_SSL"`
}
