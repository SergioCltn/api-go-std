build:
	docker build -t api-go-std .

run: build
	docker-compose up