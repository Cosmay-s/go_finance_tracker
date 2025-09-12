💼 **Проект**: Finance Tracker API  

## 📌 Описание

Сервис позволяет:
- Регистрироваться и авторизоваться
- Вести учет доходов и расходов
- Просматривать баланс
- Получать список транзакций
- Получать, обновлять и удалять транзакции по ID

## 🛠️ Технологии

| Компонент             | Технология                     |
|-----------------------|--------------------------------|
| Язык                  | Go 1.24+                      |
| Фреймворк             | Echo                          |
| База данных           | PostgreSQL                    |
| ORM                   | GORM                          |
| Миграции              | golang-migrate                |
| Аутентификация        | JWT                           |
| Контейнеризация       | Docker + Docker Compose       |
| Переменные среды      | github.com/joho/godotenv      |

## 📁 Эндпоинты

### 🧑 Аутентификация
- `POST /register` — Регистрация пользователя
- `POST /login` — Авторизация, возвращает JWT-токен

### 💵 Транзакции
- `POST /transactions` — Создание транзакции (тип: income/expense)
- `GET /transactions` — Получение списка транзакций
- `GET /balance` — Получение баланса (доходы - расходы)
- `GET /transactions/:id` — Получение транзакции по ID
- `PUT /transactions/:id` — Обновление транзакции
- `DELETE /transactions/:id` — Удаление транзакции

## 📂 Структура проекта

```
finance-tracker/
├── cmd/
│   └── main.go
├── config/
│   └── config.go         # Чтение .env
├── internal/
│   ├── models/           # GORM модели
│   ├── handlers/         # HTTP-обработчики
│   ├── services/         # Бизнес-логика
│   ├── repositories/     # Работа с БД
│   ├── middleware/       # JWT и логгеры
│   └── database/
│       └── db.go         # Подключение к PostgreSQL
├── migrations/           # SQL миграции
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.markdown
```


## 🚀 Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone <repository-url>
   cd finance-tracker
   ```

2. Настройте `.env`:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=secret
   DB_NAME=finance_tracker
   JWT_SECRET=your_jwt_secret_key
   PORT=8080
   ```

3. Инициализируйте зависимости:
   ```bash
   go mod tidy
   ```

4. Запуск с Docker:
   ```bash
   docker-compose up --build
   ```

5. Локальный запуск:
   - Убедитесь, что PostgreSQL запущен
   - Выполните миграции:
     ```bash
     migrate -path migrations -database "postgresql://postgres:secret@localhost:5432/finance_tracker?sslmode=disable" up
     ```
   - Запустите приложение:
     ```bash
     go run cmd/main.go
     ```

## 🧪 Примеры curl-запросов

### Регистрация
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com", "password":"secret"}'
```

### Логин
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com", "password":"secret"}'
```

### Получение баланса
```bash
curl -X GET http://localhost:8080/balance \
  -H "Authorization: Bearer <your_token>"
```

### Создание транзакции
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"amount": 1500, "type": "income", "timestamp": "2025-09-12T12:00:00Z"}'
```

### Обновление транзакции
```bash
curl -X PUT http://localhost:8080/transactions/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"amount": 2000, "type": "expense", "timestamp": "2025-09-12T12:00:00Z"}'
```

### Удаление транзакции
```bash
curl -X DELETE http://localhost:8080/transactions/1 \
  -H "Authorization: Bearer <your_token>"
```

### Получение списка транзакций
```bash
curl -X GET http://localhost:8080/transactions \
  -H "Authorization: Bearer <your_token>"
```

### Получение транзакции по ID
```bash
curl -X GET http://localhost:8080/transactions/1 \
  -H "Authorization: Bearer <your_token>"
```