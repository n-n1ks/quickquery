package config

func GetConfigLevel() string {
	return Env("LOG_LEVEL", "info")
}
