
.PHONY: test
test:
	go test

.PHONY: build
build: test
	go build -o application 

.PHONY: run
run: build
	./application

.PHONY: docker-build
docker-build: build
	docker build -t ueber/application:1.0.0 .

.PHONY: docker-run
docker-run: docker-build
	docker run -it --rm --name application -p 8080:8080 ueber/application:1.0.0

.PHONY: docker-push
docker-push: docker-build
	docker push ueber/application:1.0.0

