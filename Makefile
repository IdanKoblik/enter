.PHONY: all clean linux windows darwin

all: linux windows darwin

linux:
	GOOS=linux GOARCH=amd64 go build -o enter-linux .

windows:
	GOOS=windows GOARCH=amd64 go build -o enter-windows.exe .

darwin:
	GOOS=darwin GOARCH=amd64 go build -o enter-macos .

clean:
	rm -f enter-linux enter-windows.exe enter-macos
