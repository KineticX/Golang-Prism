package config
/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: config.go
 => Version: X.X.X.X
 => Author: Jonathan McAllister
 => Purpose: 
    
      Bindings for handling the config.ini file and the
      structures it contains. This file NEEDs to be kept in 
      lockstep with the config.ini file.

      The format of this file is: 
      Struct
      Instantiation
      One struct for each section in the config

*/


// Imports
// ----------------------------------
import (

	"log"
	"time"
	"github.com/go-ini/ini"
	
)


// Struct: APP
// ----------------------------------
type App struct {
	
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

}

var AppSetting = &App{}


// STRUCT: Server
// ----------------------------------
type Server struct {

	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

}

var ServerSetting = &Server{}


// STRUCT: Database
// ----------------------------------
type Database struct {

	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string

}

var DatabaseSetting = &Database{}



// Main Script Functionality
// ----------------------------------
var cfg *ini.File

// FUNC: Setup initialize the configuration instance
// ----------------------------------
func Setup() {

	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

}

// FUNC: mapTo map section
// ----------------------------------
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
