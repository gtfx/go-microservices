
all:
	GOOS=linux GOARCH=arm go build -o frontend/frontend ./frontend

build-frontend:
	docker build -t frontend:latest .

push-frontend:
	docker tag frontend:latest localhost:5000/frontend:latest
	docker push localhost:5000/frontend:latest
