all: receive send pingpong

send: send.go util.go
	go build $^

receive: receive.go util.go
	go build $^

pingpong: pingpong.go util.go
	go build $^
