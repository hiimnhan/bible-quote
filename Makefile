mainfile = cmd/main.go
binpath = bin/main

build:
	@go build -o ${binpath} ${mainfile}

run:
	./${binpath}

all: build run


