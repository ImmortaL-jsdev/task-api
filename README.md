```markdown
# Task API

Простой REST API для управления задачами на Go с использованием PostgreSQL.

## Возможности

- Получение списка всех задач (GET /tasks)
- Получение одной задачи по ID (GET /tasks/{id})
- Создание задачи (POST /tasks)
- Обновление задачи (PUT /tasks/{id})
- Удаление задачи (DELETE /tasks/{id})

## Стек технологий

- Go 1.25+
- PostgreSQL 15
- sqlx – удобная работа с БД
- lib/pq – драйвер PostgreSQL
- Docker Compose – для локального запуска БД

## Требования

- Установленный Docker и Docker Compose
- Go 1.25+ (для запуска без Docker)

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/ImmortaL-jsdev/task-api.git
cd task-api
```

### 2. Запуск PostgreSQL через Docker Compose

```bash
docker-compose up -d
```

Контейнер `task-api-db` будет запущен, база данных и таблица `tasks` создадутся автоматически из `sql/init.sql`.

### 3. Настройка переменных окружения (опционально)

Создайте файл `.env` или экспортируйте переменные:

```bash
export DATABASE_URL="postgres://taskuser:taskpass@localhost:5432/tasksdb?sslmode=disable"
export SERVER_PORT="8080"
```

Если переменные не заданы, используются значения по умолчанию (указаны выше).

### 4. Запуск сервера

```bash
go run cmd/api/main.go
```

Сервер будет доступен на `http://localhost:8080`.

## Примеры запросов (curl)

**Создать задачу**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Изучить Go","description":"Пройти курс"}'
```

**Получить все задачи**
```bash
curl http://localhost:8080/tasks
```

**Получить задачу по ID**
```bash
curl http://localhost:8080/tasks/1
```

**Обновить задачу**
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Новый заголовок"}'
```

**Удалить задачу**
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Структура проекта

```
task-api/
├── cmd/api/main.go            # точка входа
├── internal/
│   ├── database/              # работа с БД (репозиторий)
│   ├── handlers/              # обработчики HTTP
│   └── models/                # модели данных
├── sql/init.sql               # инициализация БД
├── docker-compose.yml         # запуск PostgreSQL
├── go.mod
└── README.md
```

## Команды для разработки

**Запуск тестов (пока нет, но можно добавить)**
```bash
go test ./...
```

**Остановка контейнера**
```bash
docker-compose down
```

## Лицензия

MIT
```

