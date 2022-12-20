ping:
	echo "hotel reservation API"

build:
	go build -v -ldflags="-X 'main.BuildInfo.Version=v1.0.0' -X 'main.BuildInfo.User=$(id -u -n)' -X 'main.BuildInfo.Time=$(date)'"

run:
	go run main.go --migrate

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

generate-swagger:
	swag init --parseDependency cmd\main.go -o docs/swagger

test:
	echo "Testing ..."
	go test ./...
	echo "Done!"

build-docker-compose:
	docker-compose up -d

start-vault-server:
	 vault server -dev -dev-root-token-id RhcmkiLCJhZGRyZXNzIjoiIiwicGhvbmVfbnVtYmVyIjoiOTk5OTk5OSIsInRlbmFudF9pZCI6IkkwS0ZsL