package v1
/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: ./api/v1/apiinfo.go
 => MVC Component: Controller
 => Package: v1
 => Version: 1.0.0
 => Author: Jonathan McAllister
 => Purpose: 

    This specific REST API returns [in JSON] the version information
    for the micro service. This version information is stored in 
    ./public/version.txt 

*/



// PKG Imports
// --------------------------------
import (

  "net/http"
  "github.com/gin-gonic/gin"
  "microserviceframework/models"
  "microserviceframework/pkg/app"
  "microserviceframework/pkg/e"

)

// @Summary Get APIInfo
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /apiinfo [get]
// ----------------------------------
func GetAPIInfo(c *gin.Context) {

  // -- Initialize Gin
  appG := app.Gin{C: c}

  // -- Get API Details via model call
  apidetails := models.GetAPIInfo()

  //apidetails := "{\"apiinfo\": { [\"version\": "+version+", \"githash\" : "+githash+", \"githash\" : "+compiledate+"]}}"

  // -- Perform formatting operations on data results. 

  // -- Return response to caller
  appG.Response(http.StatusOK, e.SUCCESS, apidetails)
}


