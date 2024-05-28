package mongo

import "fmt"

type Config struct {
	Host     string `envconfig:"MONGODB_HOST"`
	Port     string `envconfig:"MONGODB_PORT"`
	Username string `envconfig:"MONGODB_USERNAME"`
	Password string `envconfig:"MONGODB_PASSWORD"`
	Database string `envconfig:"MONGODB_DATABASE"`
}

func (c Config) GetConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
}
