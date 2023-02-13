package model

var Header interface{}

type ServerConfig struct {
	Name             string `env:"APP_NAME"`
	Port             string `env:"APP_PORT"`
	Host             string `env:"APP_HOST"`
	ServiceHost      string `env:"APP_SERVICE_HOST"`
	Protocol         string `env:"APP_PROTOCOL_SERVER"`
	JWTSecret        string `env:"APP_SECRET"`
	EncKey           string `env:"APP_KEY_DECRYPT"`
	JSONPathFile     string `env:"APP_JSON_PATHFILE"`
	UseResolver      string `env:"DB_RESOLVER"`
	DBConfig         DBConfig
	DBResolverConfig DBResolverConfig
	ElasticConfig    ElasticConfig
}

// elastic search config
type ElasticConfig struct {
	Host     string `env:"ES_HOST"`
	Port     string `env:"ES_PORT"`
	User     string `env:"ES_USER"`
	Password string `env:"ES_PASS"`
	Index    string `env:"ES_INDEX"`
}

// db primary config
type DBConfig struct {
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
}

// db resolver / slave config
type DBResolverConfig struct {
	Name     string `env:"DB_SLAVE_NAME"`
	Host     string `env:"DB_SLAVE_HOST"`
	Port     string `env:"DB_SLAVE_PORT"`
	User     string `env:"DB_SLAVE_USER"`
	Password string `env:"DB_SLAVE_PASS"`
}
