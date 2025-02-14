GATEWAY_BIN = gateway-service
WORKER_BIN = worker-service

GATEWAY_DIR = ./cmd/gateway-service
WORKER_DIR = ./cmd/worker-service

BUILD_FLAGS = -ldflags="-s -w"

gateway:
	@echo "Компиляция Gateway Service..."
	go build $(BUILD_FLAGS) -o $(GATEWAY_BIN) $(GATEWAY_DIR)

worker:
	@echo "Компиляция Worker Service..."
	go build $(BUILD_FLAGS) -o $(WORKER_BIN) $(WORKER_DIR)

build: gateway worker

run-gateway:
	@echo "Запуск Gateway Service..."
	./$(GATEWAY_BIN)

run-worker:
	@echo "Запуск Worker Service..."
	./$(WORKER_BIN)

run: build
	@echo "Запуск всех сервисов..."
	@./$(GATEWAY_BIN) & ./$(WORKER_BIN)

stop:
	@echo "Остановка всех сервисов..."
	@pkill -f $(GATEWAY_BIN) || true
	@pkill -f $(WORKER_BIN) || true
