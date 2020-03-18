.PHONY: coverprofile test

coverprofile:
	go tool cover -html=coverage.txt -o coverage.html
	
test:
	mkdir -p ./coverage && \
		go test -v -coverprofile=./coverage/profile.txt -covermode=atomic ./... && \
	 		go tool cover -func=./coverage/profile.txt &&\
			 	go tool cover -html=./coverage/profile.txt -o ./coverage/profile.html