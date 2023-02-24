package main

type MouseConfig struct {
	Port     int      `json:"port"`
	Etcd     Etcd     `json:"etcd"`
	Kafka    Kafka    `json:"kafka"`
	Influxdb Influxdb `json:"influxdb"`
}

type Etcd struct {
	Endpoints []string `json:"endpoints"`
	Prefix    string   `json:"prefix"`
}

type Kafka struct {
	Endpoints []string `json:"endpoints"`
	Topic     string   `json:"topic"`
}

type Influxdb struct {
	Org    string `json:"org"`
	Addr   string `json:"addr"`
	Bucket string `json:"bucket"`
	Token  string `json:"token"`
}
