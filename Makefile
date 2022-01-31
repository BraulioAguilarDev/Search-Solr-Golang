.PHONY: docker dc-up dc-down dc-stop dc-v

docker:
	docker build -t engine/golang .

dc-up:
	docker-compose -f docker-compose.yml up -d

dc-stop:
	docker-compose -f docker-compose.yml stop

dc-down:
	docker-compose -f docker-compose.yml down

dc-v:
	docker volume ls