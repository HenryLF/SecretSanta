build:
	go mod tidy
	GOOS=linux go build -o SecretSanta server.go
	GOOS=windows go build -o SecretSanta.exe server.go
	GOOS=darwin go build -o SecretSanta.app server.go

run :
	go run .
push :
	git add .
	git commit -a
	git push
