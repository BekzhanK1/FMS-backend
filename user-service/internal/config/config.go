package config

type DBConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
    DBName string
}

// TODO:
// Change to .env
func Load() DBConfig {
    return DBConfig{
        DBHost:     "localhost",
        DBPort:     "5432",
        DBUser:     "postgres",
        DBPassword: "1234",
        DBName: "fms",
    }
}