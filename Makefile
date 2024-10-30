all: build bin

build:
	go build -o $@ cmd/mc-server/main.go

bin:
	cp mc-server /opt/bin/mc-server

clean:
	rm /opt/bin/mc-server

