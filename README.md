# Task Manager API

RESTful API для управления задачами. Проект реализован на Go с использованием Gin, GORM и PostgreSQL. Сервис поддерживает создание, получение, обновление и удаление задач.

## 🔧 Стек технологий

- **Go (Golang)**
- **Gin** — HTTP фреймворк
- **GORM** — ORM для работы с БД
- **PostgreSQL** — СУБД
- **Docker** — контейнеризация

---

## 📦 Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/your-username/Task-Manager-API.git
cd Task-Manager-API
```

### 2. Запуск с Docker

```bash
docker-compose up --build
```

## 📄 Пример JSON-запроса
```bash
{
  "title": "Сделать API",
  "description": "Реализовать все CRUD методы",
  "completed": false
}
```

## Методы
| Метод  | Путь       | Описание              |
| ------ | ---------- | --------------------- |
| GET    | /tasks     | Получить все задачи   |
| POST   | /tasks     | Создать новую задачу  |
| GET    | /tasks/:id | Получить задачу по ID |
| PUT    | /tasks/:id | Обновить задачу       |
| DELETE | /tasks/:id | Удалить задачу        |

## 🧪 Миграции
GORM автоматически выполняет миграции таблиц при старте приложения (AutoMigrate в main.go).

## 🚀 Планы на будущее
- Аутентификация пользователей
- Авторизация (JWT)
- Swagger-документация
- Фильтрация и пагинация
