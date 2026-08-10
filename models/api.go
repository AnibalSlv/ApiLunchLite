package models

type ApiConfig struct {
	Id         int
	Name       string
	Host       string
	Port       int
	PathFolder string
	State      string
	Pid        int
}
