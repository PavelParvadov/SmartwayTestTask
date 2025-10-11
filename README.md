# Smartway Employee Service

Сервис представляет собой REST API для управления сотрудниками компании, предоставляющий следующие функции:

- Добавление сотрудников
- Удаление сотрудников по ID
- Вывод списка сотрудников для указанной компании
- Вывод списка сотрудников для указанного отдела компании
- Изменение сотрудника по его ID (изменяются только указанные поля)



### Запуск через Docker

1. Клонируйте репозиторий:
```bash
git clone https://github.com/PavelParvadov/SmartwayTestTask.git
```

2. Создайте файл `.env` на основе примера:
```bash
cp .env.example .env
```

3. Запустите приложение:
```bash
docker compose up --build -d
```

### Запуск из консоли

1. Клонируйте репозиторий:
```bash
git clone https://github.com/PavelParvadov/SmartwayTestTask.git

```

2. Выполните миграции базы данных:
```bash
# Через Task
task migrate-up

# Или напрямую
go run ./cmd/migrator --migrations-path=./migrations --migrations-table=migrations --db-url=postgres://postgres:postgres@localhost:5432/smartway_employees?sslmode=disable
```

3. Запустите приложение:
```bash
# Через Task
task run

# Или напрямую
go run cmd/employee-service/main.go --config=./config/config.yaml
```



##  API Endpoints

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/employees` | Создать сотрудника |
| `GET` | `/employees/company/:companyID` | Получить сотрудников компании |
| `GET` | `/employees/company/:companyID/department/:departmentID` | Получить сотрудников по отделу |
| `PATCH` | `/employees/:id` | Обновить данные сотрудника |
| `DELETE` | `/employees/:id` | Удалить сотрудника |

##  Swagger документация

После запуска приложения Swagger документация доступна по адресу:
```
http://localhost:8081/swagger/
```





