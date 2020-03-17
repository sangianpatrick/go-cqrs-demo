test:
	go test -v ./... -coverprofile=coverage.out

coverage:
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out