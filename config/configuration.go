package config

import (
	"echoserver/vo"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var errLog *log.Logger

// Data is data
var Data *vo.Configurations

func init() {
	profile := flag.String("profile", "local", "Environment profile, something similar to spring profiles")
	flag.Parse()
	viper.Set("profile", *profile)
}

// New function is initialize method
func New() (c *Configuration) {
	c = &Configuration{
		ConfigData: new(vo.Configurations),
	}
	c.ConfigData = LoadApplicationProperties()
	CreateLoggerFile(c.ConfigData)
	return
}

// CreateLoggerFile creates Logger File
func CreateLoggerFile(configData *vo.Configurations) {
	file, err := os.OpenFile("info-user.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	errLog = log.New(file, "", log.Ldate|log.Ltime)
	l := &lumberjack.Logger{
		Filename:   "info-user.log",
		MaxSize:    configData.Logging.MaxSize,    // megabytes after which new file is created
		MaxBackups: configData.Logging.MaxBackups, // number of backups
		MaxAge:     configData.Logging.MaxAge,     //days
	}
	errLog.SetOutput(l)
	log.SetOutput(l)
}

// LoadApplicationProperties => Loads application properties
func LoadApplicationProperties() *vo.Configurations {
	configFileName := fmt.Sprintf("%s%s%s", "config", "-", viper.GetString("profile"))
	viper.SetConfigName(configFileName)

	// Set the path to look for the configurations file
	viper.AddConfigPath("./configs")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.String())
	})
	var configuration vo.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err1 := viper.Unmarshal(&configuration)
	if err1 != nil {
		fmt.Printf("Unable to decode into struct, %v", err1)
	}
	return &configuration
}

type (
	// Configuration Model
	Configuration struct {
		ConfigData *vo.Configurations
	}
)
