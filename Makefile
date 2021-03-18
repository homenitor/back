APP_NAME := homenitor-back

.PHONY: run
run: run-dependencies run-app

.PHONY: run-app
run-app:
	@echo "+ $@"
	go run main.go

.PHONY: run-dependencies
run-dependencies:
	@echo "+ $@"
	@docker-compose -p $(APP_NAME) -f containers/docker-compose.yml down || true;
	@docker-compose -p $(APP_NAME) -f containers/docker-compose.yml pull;
	@docker-compose -p $(APP_NAME) -f containers/docker-compose.yml up -d --build

.PHONY: test
test:
	@echo "+ $@"
	go test -v ./...
