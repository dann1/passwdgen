build:
	go build -o bin/passwdgen src/main.go

clean:
	rm -f bin/passwdgen

all: clean build
