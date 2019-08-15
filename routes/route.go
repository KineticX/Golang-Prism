package routes

/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: routes.go
 => Version: X.X.X.X
 => Author: Jonathan McAllister
 => Purpose: 
    This set of routes configured for the micro-service  This file defines all   
	 application routes (Higher priority routes first)

 => Notes:

 		Defines the standards for basic CRUD operations in REST format
 		HTTP      Verb	HTTP Return Options
 		----      ----    ----------
 		POST	  Create	201 (Created)
 		GET	      Read	    200 (OK), 404 (Not Found), if ID not found or invalid.
 		PUT       Update	200 (OK), 405 (Method Not Allowed), 404 (Not found)
 		DELETE	  Delete	200 (OK), 405 (Method Not Allowed), 404 (Not Found)

-=-==-=---==-=-=-=-=-----=-=-=-=---=----------=-=-=-=-=-=-=-=----


*/


import (

	"github.com/gin-gonic/gin"

	_ "microserviceframework/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"microserviceframework/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// -- Serve Static VERSION Files on / [public/version.txt]	
	r.StaticFile("/version.txt", "./public/html/version.txt")

	// -- Serve Static Files on / [public/?]	
	r.StaticFile("/", "./public/html/index.html")
	
	// -- Serve Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// -- Public Access Routing Group for V1 API [not behind Auth]
	api := r.Group("/api/v1")
	{

	  api.GET("/apiinfo", v1.GetAPIInfo)
	
	}

		//		   apiv1.GET("/sample",          v1.GetSample)
		//		   apiv1.GET("/auth",            v1.GetAuth)
		//		   apiv1.GET("/apiinfo/*any",    v1.GetAPIInfo)

		//#		   apiv1.GET("/articles/:id",    v1.GetArticle)
		//#		  apiv1.POST("/articles",        v1.AddArticle)
		//#		   apiv1.PUT("/articles/:id",    v1.EditArticle)
		//#		apiv1.DELETE("/articles/:id",    v1.DeleteArticle)


	return r
}
