package env

import "os"

const (
	VAR_REGISTER_SIM_ENDPOINT_KEY = "REGISTER_SIM_ENDPOINT_KEY"
	VAR_DB_HOST                   = "DB_HOST"
	VAR_DB_PORT                   = "DB_PORT"
	VAR_DB_USERNAME               = "DB_USERNAME"
	VAR_DB_PASSWORD               = "DB_PASSWORD"
	VAR_DB_NAME                   = "DB_NAME"
)

type EnvT struct {
	REGISTER_SIM_ENDPOINT_KEY string
	DB_HOST                   string
	DB_PORT                   string
	DB_PASSWORD               string
	DB_USERNAME               string
	DB_NAME                   string
}

var Env *EnvT = new(EnvT)

func panicIfVarNotSet(set bool, var_name string) {
	if !set {
		panic("not set: " + var_name)
	}
}

func SetupEnv() {

	reg_sim_endp_key, ok := os.LookupEnv(VAR_REGISTER_SIM_ENDPOINT_KEY)
	panicIfVarNotSet(ok, VAR_REGISTER_SIM_ENDPOINT_KEY)
	Env.REGISTER_SIM_ENDPOINT_KEY = reg_sim_endp_key

	db_password, ok := os.LookupEnv(VAR_DB_PASSWORD)
	panicIfVarNotSet(ok, VAR_DB_PASSWORD)
	Env.DB_PASSWORD = db_password

	db_username, ok := os.LookupEnv(VAR_DB_USERNAME)
	panicIfVarNotSet(ok, VAR_DB_USERNAME)
	Env.DB_USERNAME = db_username

	db_port, ok := os.LookupEnv(VAR_DB_PORT)
	panicIfVarNotSet(ok, VAR_DB_PORT)
	Env.DB_PORT = db_port

	db_host, ok := os.LookupEnv(VAR_DB_HOST)
	panicIfVarNotSet(ok, VAR_DB_HOST)
	Env.DB_HOST = db_host

	db_name, ok := os.LookupEnv(VAR_DB_NAME)
	panicIfVarNotSet(ok, VAR_DB_NAME)
	Env.DB_NAME = db_name
}
