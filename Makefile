all:
	go build -o qasite.exe .
	./qasite.exe
install:
	go mod download
clean:
	rm qasite.exe