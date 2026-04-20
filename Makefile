APP_NAME=bmad-demo
VERSION=0.0.1
UI_PORT?=3000
API_PORT?=8080

.PHONY: frontend backend build clean

frontend:
	cd frontend && npm run build

backend:
	go build -o $(APP_NAME) .

build: frontend backend

clean:
	rm -rf dist frontend/dist
