package testingutils

import (
	"os"

	"github.com/gaertnerl/translate-me.git/webserver/lib/env"
)

func GetDummyEnvVars() env.EnvT {
	return env.EnvT{DB_HOST: "dummy", DB_PORT: "dummy", DB_PASSWORD: "dummy", DB_NAME: "dummy", DB_USERNAME: "dummy", REGISTER_SIM_ENDPOINT_KEY: "dummy"}
}

func SetDummyEnvVars() {
	envs := GetDummyEnvVars()
	os.Setenv(env.VAR_DB_HOST, envs.DB_HOST)
	os.Setenv(env.VAR_DB_NAME, envs.DB_NAME)
	os.Setenv(env.VAR_DB_PASSWORD, envs.DB_PASSWORD)
	os.Setenv(env.VAR_DB_PORT, envs.DB_PORT)
	os.Setenv(env.VAR_DB_USERNAME, envs.DB_USERNAME)
	os.Setenv(env.VAR_REGISTER_SIM_ENDPOINT_KEY, envs.REGISTER_SIM_ENDPOINT_KEY)
}

/* prepares application (env vars, services) for tests */
func Setup() {
	SetDummyEnvVars()
	env.SetupEnv()
}
