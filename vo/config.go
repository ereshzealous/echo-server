package vo

// ServerConfigurations is server configurations
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations is database properties
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}

// Configurations is a root model.
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Logging  Logging
}

// Person is a representation of a person
type Person struct {
	Name  string
	Phone string
}

// Logging struct
type Logging struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
}
