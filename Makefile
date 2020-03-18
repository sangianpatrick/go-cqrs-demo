.PHONY: coverprofile test

coverprofile:
	go tool cover -html=coverage.txt -o coverage.html
	
test:
	mkdir -p ./coverage && \
		go test -v -coverprofile=./coverage/profile.txt -covermode=atomic -race ./... && \
	 		go tool cover -func=./coverage/profile.txt