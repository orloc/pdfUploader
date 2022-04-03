GOCMD=go
GOBUILD=${GOCMD} build
MAIN_FILE=main.go

all: clean build

build:
	GOOS=linux ${GOBUILD} -o ./build/run -ldflags="-s -w" ${MAIN_FILE}
	cp .env ./build/.env
clean:
	rm -f ./build/run
