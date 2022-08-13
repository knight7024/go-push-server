package config

var Config = new(ViperConfig)

type ViperConfig struct {
	Core       *Core       `yaml:"core"`
	APIs       *APIs       `yaml:"apis"`
	Datasource *Datasource `yaml:"datasource"`
	Redis      *Redis      `yaml:"redis"`
}

type Core struct {
	AppName   string `mapstructure:"app_name"`
	Port      int    `mapstructure:"port"`
	Env       string `mapstructure:"env"`
	SecretKey string `mapstructure:"secret_key"`
}

type APIs struct {
	TopicSubscribeURI   string `mapstructure:"topic_subscribe"`
	TopicUnsubscribeURI string `mapstructure:"topic_unsubscribe"`
	PushMessageURI      string `mapstructure:"push_message_uri"`
	PushMulticastURI    string `mapstructure:"push_multicast_uri"`
	ProjectURI          string `mapstructure:"project_uri"`
	ProjectAllURI       string `mapstructure:"project_all_uri"`
	UserLoginURI        string `mapstructure:"user_login_uri"`
	UserLogoutURI       string `mapstructure:"user_logout_uri"`
	UserSignupURI       string `mapstructure:"user_signup_uri"`
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
