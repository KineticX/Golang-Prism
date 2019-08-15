# 
# __________        .__                
# \______   \_______|__| ______ _____  
#  |     ___/\_  __ \  |/  ___//     \ 
#  |    |     |  | \/  |\___ \|  Y Y  \
#  |____|     |__|  |__/____  >__|_|  /
#                          \/      \/ 
#
# ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------
#    => File: Makefile
#    => Version: X.X.X.
#    => Author: Jonathan McAllister
#    => Purpose: 
#       This is the Makefile for the go microservice
#       Targets are all listed below

# -- Targets
.PHONY: build clean run docgen toolinstall lint help
VERSION=1.2.3.4

# Target: ALL
# ---------------------
all: clean build docgen

# Target: RUN
# ---------------------
run: 
	go run -ldflags "-X main.version=1.0.0.0" main.go;


# Target: VERSION
# ---------------------
version:
	rm ./public/html/version.txt
	echo "${VERSION}" >> ./public/html/version.txt
	git rev-parse HEAD >> ./public/html/version.txt
	date  >> ./public/html/version.txt



# Target: BUILD
# ---------------------
build:

	go build -ldflags "-X main.version=1.0.0.0" main.go;
	#swag init -g config/config.go

# Target: DOCGEN
# ---------------------
docgen: 
	swag init -g config/config.go;

# Target: TOOLINSTALL
# ---------------------
toolinstall:
	go get;

# Target: CLEAN
# ---------------------
clean:
	rm -rf go-gin-example
	go clean -i .

# Target: HELP
# ---------------------
help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"

