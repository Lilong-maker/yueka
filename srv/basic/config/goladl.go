package config

type AppConfig struct {
	Mysql
	Redis
}
type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}
