package env

import "os"

const (
	var_name_register_sim_endpoint_key = "REGISTER_SIM_ENDPOINT_KEY"
	var_name_db_host                   = "DB_HOST"
	var_name_db_port                   = "DB_PORT"
	var_name_db_username               = "DB_USERNAME"
	var_name_db_password               = "DB_PASSWORD"
	var_name_db_name                   = "DB_NAME"
)

type env struct {
	REGISTER_SIM_ENDPOINT_KEY string
	DB_HOST                   string
	DB_PORT                   string
	DB_PASSWORD               string
	DB_USERNAME               string
	DB_NAME                   string
}

var Env *env = new(env)

func panicIfVarNotSet(set bool, var_name string) {
	if !set {
		panic(var_name)
	}
}

func SetupEnv() {

	reg_sim_endp_key, ok := os.LookupEnv(var_name_register_sim_endpoint_key)
	panicIfVarNotSet(ok, var_name_register_sim_endpoint_key)
	Env.REGISTER_SIM_ENDPOINT_KEY = reg_sim_endp_key

	db_password, ok := os.LookupEnv(var_name_db_password)
	panicIfVarNotSet(ok, var_name_db_password)
	Env.DB_PASSWORD = db_password

	db_username, ok := os.LookupEnv(var_name_db_username)
	panicIfVarNotSet(ok, var_name_db_username)
	Env.DB_USERNAME = db_username

	db_port, ok := os.LookupEnv(var_name_db_port)
	panicIfVarNotSet(ok, var_name_db_port)
	Env.DB_PORT = db_port

	db_host, ok := os.LookupEnv(var_name_db_host)
	panicIfVarNotSet(ok, var_name_db_host)
	Env.DB_HOST = db_host

	db_name, ok := os.LookupEnv(var_name_db_name)
	panicIfVarNotSet(ok, var_name_db_name)
	Env.DB_NAME = db_name
}
