all: receive send

send: send.go util.go
	go build $^

receive: receive.go util.go
	go build $^
