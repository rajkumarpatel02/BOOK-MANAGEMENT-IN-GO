package config

// Dummy functions to avoid DB connection crash when testing without MySQL

func Connect() {}

func GetDB() interface{} {
	return nil
}
