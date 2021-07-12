all:

build:
	cd cmd && go build -o redis_test *.go

test:
