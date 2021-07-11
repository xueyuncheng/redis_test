package config

type Config struct {
	Cache    *Cache
	Database *Database
	Http     *Http
}

type Cache struct {
	Host     string
	Port     string
	UserName string
	Password string
}

type Database struct {
	Host         string
	Port         string
	UserName     string
	Password     string
	DatabaseName string
}

type Http struct {
	Host string
	Port string
}
