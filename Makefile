dev:
	docker-compose up -d;
	watchexec --restart --exts "go" --watch . "docker-compose restart app";

stop:
	docker-compose down;
