package consts

const (
	DBHostEnv     = "DB_HOST"
	DBPortEnv     = "DB_PORT"
	DBUserEnv     = "DB_USER"
	DBPasswordEnv = "DB_PASSWORD"
	DBNameEnv     = "DB_NAME"

	MySQLDSNFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)
