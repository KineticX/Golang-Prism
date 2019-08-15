package logging

/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: ./api/v1/apiinfo.go
 => Package: v1
 => Version: 1.0.0
 => Author: Jonathan McAllister
 => Purpose: 
    Returs the version information for the API

-=-==-=---==-=-=-=-=-----=-=-=-=---=----------=-=-=-=-=-=-=-=----
*/



// Package Imports
// -------------------------------
import (

	"fmt"
	"time"
	"microserviceframework/config"
)

// getLogFilePath get the log file save path
// -------------------------------
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", config.AppSetting.RuntimeRootPath+config.ServerSetting.RunMode+"_", config.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
// -------------------------------
func getLogFileName() string {

	fmt.Sprintf("Hello")

	return fmt.Sprintf("%s%s.%s",
		config.AppSetting.LogSaveName,
		time.Now().Format(config.AppSetting.TimeFormat),
		config.AppSetting.LogFileExt,
	)
}
