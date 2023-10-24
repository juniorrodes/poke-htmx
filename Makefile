run-build:
	docker-compose up -d --build
	docker logs toy-project-app-1 -f