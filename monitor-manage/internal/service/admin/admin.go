package admin

import (
	"database/sql"
	"tx_gdut_monitor/internal/dao/admin"
)

type Admin struct {
	outputAdmin admin.IOutput
	checkAdmin admin.ICheck
}

func Register(db *sql.DB) *Admin {
	return &Admin{
		outputAdmin: admin.NewOutput(db),
	}
}
