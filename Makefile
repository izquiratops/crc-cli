SOURCE_FILE = main.go
BINARY_NAME = json-crc32
GO_BUILD = go build -ldflags "-w"

all: build

build:
	$(GO_BUILD) -o $(BINARY_NAME) $(SOURCE_FILE)

clean:
	rm -f $(BINARY_NAME)

rebuild: clean build

.PHONY: all build clean rebuild