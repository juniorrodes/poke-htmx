run-build:
	docker-compose up -d --build
	docker logs toy-project-app-1 -f
run:
	docker-compose up -d
	docker logs toy-project-app-1 -f