IMAGE_NAME=ll931217/portfolio:latest

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -d -p 23234:23234 $(IMAGE_NAME)

push:
	docker push $(IMAGE_NAME)
