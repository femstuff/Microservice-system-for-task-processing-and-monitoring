GATEWAY_BIN = gateway-service
WORKER_BIN = worker-service
DOCKER_COMPOSE = docker compose
K8S_DIR = ./k8s

GATEWAY_DIR = ./api-gateway/cmd/gateway-service
WORKER_DIR = ./worker-service/cmd/worker

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

dc_build:
	$(DOCKER_COMPOSE) build

dc_up:
	$(DOCKER_COMPOSE) up -d

dc_down:
	$(DOCKER_COMPOSE) down

logs-api:
	$(DOCKER_COMPOSE) logs -f $(GATEWAY_BIN)

logs-worker:
	$(DOCKER_COMPOSE) logs -f $(WORKER_BIN)

dc_clean:
	$(DOCKER_COMPOSE) down -v --rmi all --remove-orphans

dc_restart: down up

dc_ps:
	$(DOCKER_COMPOSE) ps

k8s_apply:
	@echo "Применяем манифесты в Kubernetes..."
	kubectl apply -f $(K8S_DIR)

k8s_delete:
	@echo "Удаляем ресурсы из Kubernetes..."
	kubectl delete -f $(K8S_DIR)

k8s_logs_redis:
	@echo "Логи Redis..."
	kubectl logs -l app=redis

k8s_get_pods:
	@echo "Получение списка подов..."
	kubectl get pods