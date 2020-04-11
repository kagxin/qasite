all:
	go build -o qasite cmd/main.go
clean:
	rm qasite