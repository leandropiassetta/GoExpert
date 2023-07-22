package configs

var config *Config

type Config struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExperesIn  int
}

// init is called before main function is executed by Go runtime environment and is used to initialize global variables.

// func init() {
// }
