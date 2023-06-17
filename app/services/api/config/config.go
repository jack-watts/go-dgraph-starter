package config

import (
	"net/http"
	"time"
)

type Config struct {
	Web   http.Server
	DgURL string
}

func (c *Config) Initialize() {
	c.Web.ReadTimeout = time.Second * 5
	c.Web.WriteTimeout = time.Second * 10
	c.Web.IdleTimeout = time.Second * 120
}
