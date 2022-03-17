TARGET := go-skep

.PHONY: $(TARGET) gen vet test lint clean distclean jenkins

all: $(TARGET)

$(TARGET):
	go build -o $@ main.go
