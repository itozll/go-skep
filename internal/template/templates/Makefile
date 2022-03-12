TARGET := {{ .app_name }}

.PHONY: $(TARGET) gen vet test lint clean distclean jenkins

all: $(TARGET)

$(TARGET):
	go build -o bin/$@ main.go

gen:
	@find . -name "*_temp.go" -type f -exec rm -rf {} \;
	go generate ./...

vet:
	go vet ./...

test:
	go test -race -cover -coverprofile=cover.out ./...

lint:
	golangci-lint run

clean:
	rm -rf 'bin/*'

distclean: clean
	@find . -name "*_temp.go" -type f -exec rm -rf {} \;

jenkins: gen
	go build -a -o bin/$(TARGET) main.go