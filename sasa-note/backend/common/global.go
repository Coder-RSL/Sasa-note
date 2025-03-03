package common

import (
	"backend/config"
	"database/sql"
	"go.uber.org/zap"
)

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server
)

var (
	Mysql *sql.DB
)
