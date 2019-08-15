package app

/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: response.go
 => Package: v1
 => Version: 1.0.0
 => Author: Jonathan McAllister
 => Purpose: 

    Returs the version information for the API.

*/

import (

	"github.com/gin-gonic/gin"
	"microserviceframework/pkg/e"

)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
