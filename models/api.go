package models

type ApiConfig struct {
	Id         int
	Name       string
	Host       string
	Port       int
	Path       string
	PathFolder string
	State      string
	Pid        int
}
