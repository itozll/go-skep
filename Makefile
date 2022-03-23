TARGET := go-skep

.PHONY: $(TARGET)

all: $(TARGET)

$(TARGET):
	go build -o $@ main.go
