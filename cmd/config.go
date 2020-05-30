package main

type Opts struct {
	Redis RedisGroup `group:"redis" namespace:"redis" env-namespace:"REDIS"`
}

type RedisGroup struct {
	Addr   string `long:"addr" env:"ADDR" description:"Redis <host:port> address"`
	Passwd string `long:"passwd" env:"PASSWD" description:"Redis password"`
	DB     int    `long:"db" env:"DB" default:"0" description:"Redis database to be selected"`
}
