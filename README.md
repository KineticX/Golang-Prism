# Prism Micro Service Solutions

Prism MicroService Framework aims to implement an enterprise ready full featured Micro Service solution in GO. There are a few of these already out there but none seemed enterprise ready and often lacked features/solutions that one would expect in something enterprise grade [service discovery, versioning, etc.]. This framework aims to be simple, convention based, encapsulated and clean. 

The current implementation of the framework was developed using GIN, REST, MVC, etc. and generally works fairly well. It offers the concept of views because some microservices might be UI driven where as others would be data/api driven.

### Note: 

While golang programmers generally prefer to implement solutions from scratch, doing so with a solution that is enterprise ready can lead to varying implementations and half baked ideas with potential security holes, and gaps. Not always is re-inventing a wheel a good idea.

# Features

Currently the framework has implemented support for the following concepts:

    * MVC Structure [Model View Controller] [GIN+Custom]
    * Restful APIs [versioning on the URL with filesystem abstractions]
    * Database Support for MySQL
    * Swagger doc generation
    * Automated Testing Structures
    * Basic RESTful API implemnented
    * Logging [debug logs, release logs] [Warn, Fatal, Error, Info]
    * Easy editing ini config files [db connections, mode etc.] 
    * JSON parsing
    * Versioning baked in per makefile

 Swagger-UI is at http://localhost:8000/swagger/index.html 

 GIN [https://github.com/gin-gonic/gin] is A high-productivity web framework for the [Go language](http://www.golang.org/).

# Getting Started:

To get set up you will need the following installed on your system:

* golang [#> brew install go] 
* gin [#> go get]
* All other libs [#> go get ]
* MySQL [define connection in ./config/app.ini]

Then start normally using [#> go run main.go] and enjoy.

## Code Layout

The directory structure of a generated GIN application:


    config/           Configuration directory
        app.ini       Main app configuration file

    api/v1            API Controller Code
        sample.go     Simple Controller

    public/           Public static assets    
        html/         html files
        css/          CSS files
        js/           Javascript files
        images/       Image files

    routes/           Route definitions folder
        routes.go     Routes file

    logs/             Log output folders
        debug_logs    Debug log files
        release_logs  Release log files

    tests/            Test suites

    cli/              Beginnings of the CLI

    models/           Data connectivity / models

## Help

Help is a google search away, but here are some fundamentals:

  * The [Getting Started with GIN](https://github.com/gin-gonic/gin) docs.
  * Getting started with Swagger notation [ swag - Automatically generate RESTful API documentation with Swagger 2.0 for Go.]
  * MVC Patterns [https://www.tutorialspoint.com/design_pattern/mvc_pattern]
  * INI Config [https://en.wikipedia.org/wiki/INI_file]

## Versioning

The versioning constructs of this micro service are embedded into the following file:

  * ./publi/html/version.txt

It contains the following 3 lines:

  * Version number
  * Git Commit Hash
  * Compilation Date Time

This file is read at start up and is then delivered into 2 locations:

  * The API on the URL at: http://<APIURL>/api/v1/apiinfo [JSON]
  * The execution of the GIN command line [make run]

The reason versioning was done this way is simple: 

As the micro services concept drives forward, we will need a way to perform service registry / discovery, and debug/trace defects. 

Jenkins can create this file at compile time [version number, git hash, dt] and the file will then act as a marker for debugging issues that may have arisen from that version. This also helps with service discovery whereas the service discovery microservice can pull this information via a simple API call and register what services are in the wild. 

## Todo

  * Implement the automation of CLI design via the rest API doc contents
  * Build service discovery connectivity & elastic scaling through K8
  * Fill out the example API and get JSON results for calls.
  * Heartbeat for Service Discovery / alive / autoscaling

