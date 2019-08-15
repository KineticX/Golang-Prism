package v1

/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: ./api/v1/sample.go
 => Package: v1
 => MVC Component: Controller
 => Version: 1.0.0
 => Author: Jonathan McAllister
 => Purpose: 
      This is a basic sample of a REST API contoller in PRISM
	  Its aim is to example how the solution would be implemented
*/



// Package Imports 
// ----------------------------------
import (

	"fmt"
	"net/http"

	"microserviceframework/models"
	"microserviceframework/pkg/app"
	"microserviceframework/pkg/e"

	"github.com/gin-gonic/gin"

)

// @Summary Get a single sample
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/sample [get]
func GetSample(c *gin.Context) {

	appG := app.Gin{C: c}

	sample,err := models.GetSample(1)
	if err != nil {
	fmt.Printf("ERROR:")

	}
	appG.Response(http.StatusOK, e.SUCCESS, sample)
}


