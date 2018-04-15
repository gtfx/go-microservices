FRONTEND_APP=frontend
BACKEND_APP=backend

all:
	build-frontend
	push-frontend


frontend-build:
	docker build --build-arg app=${FRONTEND_APP} -t frontend:latest .

frontend-push:
	docker tag frontend:latest localhost:5000/frontend:latest
	docker push localhost:5000/frontend:latest

backend-build:
	docker build --build-arg app=${BACKEND_APP} -t backend:latest .

backend-push:
	docker tag backend:latest localhost:5000/backend:latest
	docker push localhost:5000/backend:latest
