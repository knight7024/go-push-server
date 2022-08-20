package config

var Config = new(ViperConfig)

type ViperConfig struct {
	Core       *Core       `yaml:"core"`
	Datasource *Datasource `yaml:"datasource"`
	Redis      *Redis      `yaml:"redis"`
}

type Core struct {
	AppName   string `mapstructure:"app_name"`
	SecretKey string `mapstructure:"secret_key"`
}

type Datasource struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	DBName   string `mapstructure:"dbname"`
}

type Redis struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	DBName   int    `mapstructure:"dbname"`
}
