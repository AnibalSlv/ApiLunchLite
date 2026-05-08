package models

type ApiType = ApiConfig

type Db interface {
	InitDb() error
	Save(Api ApiConfig)
	UpdatePID(pid int, id int) error
	UpdateState(state string, id int) error
	GetAll() ([]ApiType, error)
	GetName(name string) (ApiType, error)
	GetId(id int) (ApiType, error)
	Delete(id int) error
}
