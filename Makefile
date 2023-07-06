run: build
	./bin/qchat

build: main.go
	go build -o ./bin/qchat

watch:
	gowatch

.PHONY: build run
.SILENT: build run