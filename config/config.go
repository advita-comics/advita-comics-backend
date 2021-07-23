package config

// Config - конфиги приложения
type Config struct {
	HTTP     HTTP
	Log      Log
	DB       DBConfig
	FilePath string
}

// HTTP - конфиги сервера
type HTTP struct {
	Port int
}

// Log - конфиги логера
type Log struct {
	Level int64
	Path  string
}

// DBConfig - конфиги подключения к базе
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}
