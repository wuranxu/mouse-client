package main

type MouseConfig struct {
	Port  int   `json:"port"`
	Etcd  Etcd  `json:"etcd"`
	Kafka Kafka `json:"kafka"`
}

type Etcd struct {
	Endpoints []string `json:"endpoints"`
	Prefix    string   `json:"prefix"`
}

type Kafka struct {
	Endpoints []string `json:"endpoints"`
	Topic     string   `json:"topic"`
}
