help: ## Displays description of the goals - everything that is written after a double sharp (##) separated by a space
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

build-docker-image: ## Собирает docker образ для линтера
build-docker-image:
	docker build --tag openapi-linter:latest .

push-docker-image: ## Публикует docker образ для линтера
push-docker-image: build-docker-image
	docker push openapi-linter:latest

lint: ## Запускает линтер
lint: build-docker-image
	docker run --rm --volume $(PWD):/data --entrypoint lint-openapi openapi-linter:latest --verbose --report_statistics openapi.yaml
