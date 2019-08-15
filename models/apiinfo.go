package models
/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: apiinfo.go
 => Version: X.X.X.X
 => Author: Jonathan McAllister
 => Purpose: 
    
      Data fetch for APIInfo API call. This is a model that reads
      the contents of ./public/html/version.txt and returns the 
      read contents in JSON format

*/


// Imports
// ----------------------------------
import (

	"bufio"
	"os"
	"strings"
	"microserviceframework/pkg/logging"

)

// STRUCT: [return item & format]
// ----------------------------------
type APIInfo struct {
	
	Version           string `json:"version"`
	Githash           string `json:"githash"`
	Compiledate       string `json:"compiledate"`
}

// FUNC: GetAPIInfo
// ----------------------------------
func GetAPIInfo()(*APIInfo) {

	var apidetails APIInfo
    
    logging.Info("Fetching API version information")

	// -- Open file 
	file, err := os.Open("./public/html/version.txt")
 	if err != nil {
		return nil
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
 			apidetails.Version = strings.TrimSuffix(line, "\n")
 		} 		

		// -- githash is line 2
 		if i == 1 {
 			apidetails.Githash = strings.TrimSuffix(line, "\n")
 		} 		

		// -- compiledate is line 3
 		if i == 2 {
 			apidetails.Compiledate = strings.TrimSuffix(line, "\n")
 		} 		

 		// -- nil check
 		if err != nil {
 			return nil
 		}
 	}
 	return &apidetails
}

