package main

type Config struct {
	Port string
}

func NewConfig(port string) *Config {
	return &Config{Port: port}
}
