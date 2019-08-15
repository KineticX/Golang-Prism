package main
/* 
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: main.go
 => Version: X.X.X.X
 => Author: Jonathan McAllister
 => Purpose: 
    This is the Micro Service entry point. It is based on MVC

    Model => ./models [Interact with data and do queries]
    Views => ./public [Served to the public .html etc.]
    APIControllers => ./api/v1 [Perform logic and business logic]

-=-==-=---==-=-=-=-=-----=-=-=-=---=----------=-=-=-=-=-=-=-=----
*/

// Package Imports 
// ----------------------------------
import (

	"fmt"
	"log"
	"net/http"
	"bufio"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"microserviceframework/models"
	"microserviceframework/pkg/logging"
	"microserviceframework/config"
	"microserviceframework/routes"

)

// Init
// ----------------------------------
func init() {

	config.Setup()
	models.Setup()
	logging.Setup()
	
}

// Version Information
// ----------------------------------
var version     string
var githash     string
var compiledate string


// @title Micro Service API [MVC] [GIN]
// @version 1.0
// @description Micro Service Framework using MVC and GIN
// ----------------------------------
func main() {

	// -- Get Versioning info
	version, githash, compiledate := GetVersion(`./public/html/version.txt`)

	// -- Display Header:
	fmt.Printf("\033[H\033[2J")
	fmt.Printf("__________        .__                       \n")         
	fmt.Printf("\\______   \\_______|__| ______ _____       \n")
 	fmt.Printf("|     ___/\\_  __ \\  |/  ___//     \\      \n")
 	fmt.Printf("|    |     |  | \\/  |\\___ \\|  Y Y  \\.   \n")
 	fmt.Printf("|____|     |__|  |__/____  >__|_|  /.       \n")
    fmt.Printf(".                        \\/      \\/                                \n") 

	fmt.Printf("---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=------- \n")
	fmt.Printf(" => Micro Service: Prism Micro Service Example\n")
    fmt.Printf(" => Version: %s \n", version)
    fmt.Printf(" => Git Commit: %s \n", githash)
    fmt.Printf(" => Compilation Date: %s \n", compiledate)
    fmt.Printf(" => Runtime Mode: %s \n", config.ServerSetting.RunMode)
    fmt.Printf(" => Author: Jonathan McAllister \n")
    fmt.Printf(" => Purpose: Micro Service for Prism \n\n")

	gin.SetMode(config.ServerSetting.RunMode)

	// - Variable Declarations
	routesInit     := routes.InitRouter()
	readTimeout    := config.ServerSetting.ReadTimeout
	writeTimeout   := config.ServerSetting.WriteTimeout
	endPoint       := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	// -- Specify Server Detail
	server         := &http.Server{

		Addr:           endPoint,
		Handler:        routesInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,

	}

	// -- Logging Examples
	logging.Info("This is an info message")
	logging.Warn("This is a warning")
	logging.Error("This is a non fatal error")
	
	//logging.Fatal("Error here")

	// -- Listen and Serve
	log.Printf("\n[INFO] Started http server [listening at %s]", endPoint)


	// -- Start the actual server
	server.ListenAndServe()

}

// Get Version
// ----------------------------------
func GetVersion(filename string) (version string, githash string, compiledate string) {


	// -- Open file 
	file, err := os.Open(filename)
 	if err != nil {
		return "0.0.0.0","Unknown","Today"
 	}

 	// -- Read File
 	reader := bufio.NewReader(file)

 	// -- Iterate lines in file [expected 3]
    for i := range [3]int{} {

    	// -- Read currentline
 		line, err := reader.ReadString('\n')

 		// -- Could be in a switch case type statement worth it?
		// -- version number is line 1
 		if i == 0 {
 			version = strings.TrimSuffix(line, "\n")
 		} 		

		// -- githash is line 2
 		if i == 1 {
 			githash = strings.TrimSuffix(line, "\n")
 		} 		

		// -- compiledate is line 3
 		if i == 2 {
 			compiledate = strings.TrimSuffix(line, "\n")
 		} 		

 		// -- nil check
 		if err != nil {
 			return "0.0.0.0","Unknown","Today"
 		}
 	}
 	return version, githash, compiledate
}

