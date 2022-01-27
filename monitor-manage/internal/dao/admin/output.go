package admin

import "database/sql"

type Output struct {
	db *sql.DB
}

func NewOutput(db *sql.DB) *Output {
	return &Output{
		db: db,
	}
}

func (o *Output) GetConfig() error {
	panic("implement me")
}

func (o *Output) AddNewConfig() error {
	panic("implement me")
}

func (o *Output) UpdateConfig() error {
	panic("implement me")
}

func (o *Output) DelConfig() error {
	panic("implement me")
}
