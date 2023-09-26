export PATH=$PATH:$(go env GOPATH)/bin
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g api/*