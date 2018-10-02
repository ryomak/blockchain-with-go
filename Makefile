PORT=5000
deps:
	dep ensure
run:
	go run main.go -p $(PORT)

