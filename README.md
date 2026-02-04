# Microservice System for Task Processing and Monitoring

## Описание проекта
Этот проект представляет собой микросервисную систему для обработки задач и их мониторинга. В текущей версии используются **Go, Docker, Kubernetes, RabbitMQ и Redis**.

В дальнейшем планируется добавить **Prometheus** для сбора метрик, **Kafka** для обработки событий и **покрытие кода тестами**.

## Компоненты системы
- **API Gateway** – принимает задачи и отправляет их в очередь RabbitMQ.
- **Worker Service** – обрабатывает задачи и сохраняет результаты.
- **RabbitMQ** – брокер сообщений для очереди задач.
- **Redis** – кэш для хранения результатов.
- **Kubernetes** – оркестрация контейнеров (в процессе интеграции).
- **Prometheus (планируется)** – мониторинг метрик системы.
- **Kafka (планируется)** – асинхронная обработка событий.

## Запуск проекта
### Требования
Перед запуском убедитесь, что у вас установлены:
- **Docker** 
- **Go** 
- **Kubernetes (Minikube)** 

### Шаги для запуска
1. **Клонируйте репозиторий:**
   ```sh
   git clone https://github.com/femstuff/Microservice-system-for-task-processing-and-monitoring.git
   cd Microservice-system-for-task-processing-and-monitoring
   ```

2. **Соберите бинарные файлы (опционально для локального запуска):**
   ```sh
   make build
   ```

3. **Запустите сервисы (локально без контейнеров):**
   ```sh
   make run
   ```

4. **Запустите через Docker Compose:**
   ```sh
   make dc_up
   ```

5. **Посмотрите логи API Gateway:**
   ```sh
   make logs-api
   ```

6. **Посмотрите логи Worker Service:**
   ```sh
   make logs-worker
   ```

7. **Остановите контейнеры:**
   ```sh
   make dc_down
   ```

### Запуск в Kubernetes (Minikube)
1. **Запустите Minikube:**
   ```sh
   minikube start
   ```
2. **Примените манифесты для разворачивания сервисов:**
   ```sh
   kubectl apply -f deployments/rabbitmq-deployment.yaml
   kubectl apply -f deployments/gateway-deployment.yaml
   kubectl apply -f deployments/worker-deployment.yaml
   ```
3. **Проверьте состояние подов:**
   ```sh
   kubectl get pods
   ```
4. **Остановите Minikube при необходимости:**
   ```sh
   minikube stop
   ```

## Команды Makefile
| Команда          | Описание                          |
|-----------------|---------------------------------|
| `make build`   | Собирает бинарники сервисов      |
| `make run`     | Запускает оба сервиса локально   |
| `make stop`    | Останавливает локальные сервисы |
| `make dc_up`   | Запускает контейнеры через Docker Compose |
| `make dc_down` | Останавливает контейнеры        |
| `make logs-api` | Показывает логи API Gateway     |
| `make logs-worker` | Показывает логи Worker Service |

## Дальнейшие планы
-  **Интеграция Prometheus** – для сбора и мониторинга метрик.
-  **Добавление Kafka** – для улучшения асинхронной обработки задач.
-  **Покрытие кода тестами** – для повышения надежности системы.
-  **Доработка Kubernetes манифестов** – для автоматического масштабирования сервисов.

## Ссылка на репозиторий
[GitHub: Microservice System for Task Processing and Monitoring](https://github.com/femstuff/Microservice-system-for-task-processing-and-monitoring)

