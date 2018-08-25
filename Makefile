build-connector-service:
	@echo "=============== build connector-service start ==============="
	@cd connector-service && make build
	@echo "=============== build connector-service end ================="

build-lobby-service:
	@echo "=============== build lobby-service start ==============="
	@cd lobby-service && make build
	@echo "=============== build lobby-service end ================="

build-mjgame-service:
	@echo "=============== build mjgame-service start ==============="
	@cd mjgame-service && make build
	@echo "=============== build mjgame-service end ================="

build-nngame-service:
	@echo "=============== build nngame-service start ==============="
	@cd nngame-service && make build
	@echo "=============== build nngame-service end ================="

build-web-service:
	@echo "=============== build web-service start ==============="
	@cd web-service && make build
	@echo "=============== build web-service end ================="

build-all: build-connector-service build-lobby-service build-mjgame-service build-nngame-service build-web-service
