run-build:
	docker-compose up -d --build
	docker-compose logs app -f -t