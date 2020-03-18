test:
	go test -v ./... -coverprofile=coverage.out

cover-html:
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

cover-func:
	go test -v ./... -coverprofile=coverage.out && go tool cover -func=coverage.out