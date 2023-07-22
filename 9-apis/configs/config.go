package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var config *Config

// Config is the configuration for the application. It is loaded from a file. The file is in JSON format. The file is loaded using the LoadConfig function. The configuration is stored in the config variable. The config variable is a pointer to a Config struct. The file .env inject the values for the Config struct.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExperesIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

// init is called before main function is executed by Go runtime environment and is used to initialize global variables.

// func init() {
// 	config = &Config{
// 		DBDriver:      "mysql",
// 		DBHost:        "localhost",
// 		DBPort:        "3306",
// 	}
// }

// LoadConfig loads the configuration from a file.
func LoadConfig(path string) (*Config, error) {
	// i can many config files in the same time and viper will merge them together and override the values if they are duplicated in the files and the order of the priority is the order of the files
	viper.SetConfigName("app_config")
	// which is the type of configuration file we are using (JSON, YAML, TOML, etc.)
	viper.SetConfigType("env")
	// the path to the directory where the configuration file is located
	viper.AddConfigPath(path)

	// which is the name of the environment variable that contains the path to the configuration file
	viper.SetConfigFile(".env")
	// viper will do automactically the unmarshalling of the configuration file into the Config struct
	viper.AutomaticEnv()

	// read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		// i dont want that my application up and running if the configuration file is not loaded
		panic(err)
	}

	// unmarshal the configuration file into the Config struct
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	// initialize the JWTAuth
	// crate a instance of JWTAuth
	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	return config, nil
}
