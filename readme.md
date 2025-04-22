# Структура проекта
```
.
├── cmd
│   ├── app
│   │   └── main.go
│   ├── migrate.go
│   ├── root.go
│   └── serve.go
├── config
│   ├── config.go
│   └── util_transform.go
├── docker-compose.yml
├── dockerfile
├── entrypoint.sh
├── go.mod
├── go.sum
├── internal
│   ├── delivery
│   │   └── httpdelivery
│   │       ├── middleware
│   │       │   └── middleware.go
│   │       ├── order_handler.go
│   │       ├── router.go
│   │       └── user_handler.go
│   ├── domain
│   │   ├── order.go
│   │   ├── product.go
│   │   └── user.go
│   ├── models
│   │   ├── order_request.go
│   │   ├── responses.go
│   │   └── user_request.go
│   ├── observability
│   │   └── trace.go
│   ├── repository
│   │   └── postgresrepo
│   │       ├── db.go
│   │       ├── migrations
│   │       │   ├── 1_init.down.sql
│   │       │   ├── 1_init.up.sql
│   │       │   └── 2_seed_products.up.sql
│   │       ├── order_repo.go
│   │       └── user_repo.go
│   ├── usecase
│   │   ├── order_async.go
│   │   ├── order.go
│   │   └── user.go
│   └── utils
│       └── retry.go
├── makefile
├── other
│   ├── docker_for_local
│   │   └── docker-compose.yml
│   └── Golang 2.pdf
├── readme.md
├── tests
│   ├── curls.txt
│   ├── order_usecase_test.go
│   └── user_usecase_test.go
└── wait-for-it.sh

19 directories, 40 files
```

#### [Техническое задание](other/Golang%202.pdf)

# Варианты запуска

1. Сборка приложения в докере
`docker-compose up --build`

2. Локальный запуск
`make env & make run` (для удобства так же создал доп докер /other/docker_for_local)

# Ручки

### Регистрация нового пользователя.
`POST /users/register`

```
curl -X POST http://localhost:8080/users/register -H "Content-Type: application/json" -d '{"first_name":"Alice","last_name":"Johnson","age":25,"is_married":true,"password":"strongpass123"}'
```

### Создание заказа пользователем (асинхронно).
`POST /orders`

```
curl -X POST http://localhost:8080/users/register -H "Content-Type: application/json" -d '{"first_name":"Bob","last_name":"Short","age":30,"is_married":false,"password":"123"}'
```
