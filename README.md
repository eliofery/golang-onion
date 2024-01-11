# Rest API с реализацией Луковичной архитектуры на Fiber фреймворке

Реализация CRUD Rest API на фреймворке Fiber с авторизацией по JWT токену.

В качестве базы данной была использована Postgres.

Основной уклон был сделан на построении Луковичной архитектуры и реализации ролей у пользователей.

Список доступного апи.

![Скриншот](screen.png)

## Команды проекта

Автоматическая сборка проекта при изменении файлов.

```bash
docker compose up -d

make watch
```

## Используемые пакеты

[fiber - фреймворк](https://github.com/gofiber/fiber)

```bash
go get -u github.com/gofiber/fiber/v3
```

[godotenv - переменные окружения](https://github.com/joho/godotenv)

```bash
go get github.com/joho/godotenv
```

[testify - тестирование](https://github.com/stretchr/testify)

```bash
go get github.com/stretchr/testify
```

[viper - yml конфигурация](https://github.com/spf13/viper)

```bash
go get github.com/spf13/viper
```

[sqlite3 - база данных sqlite](https://github.com/mattn/go-sqlite3)

```bash
go get github.com/mattn/go-sqlite3
```

[pgx - база данных postgres](https://github.com/jackc/pgx)
[pgerrcode - коды ошибок postgres](https://github.com/jackc/pgerrcode)

```bash
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgerrcode
```

[goose - миграция базы данных](https://github.com/pressly/goose)

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
go get github.com/pressly/goose/v3
```

[mysql - база данных mysql](https://github.com/go-sql-driver/mysql)

```bash
go get github.com/go-sql-driver/mysql
```

[jwt - токен](https://github.com/golang-jwt/jwt)

```bash
go get -u github.com/golang-jwt/jwt/v5
```

[modd - автоматическая компиляция при изменении файлов](https://github.com/cortesi/modd)

```bash
go install github.com/cortesi/modd/cmd/modd@latest
```

[validator - валидация данных](https://github.com/go-playground/validator)

```bash
go get github.com/go-playground/validator/v10
```

[bcrypt - шифрование](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

```bash
go get golang.org/x/crypto/bcrypt
```

## Миграция

### Создание миграции

```bash
# Создание миграции
goose -dir ./internal/database/migration create <имя миграции> sql

# Переименовывает миграции с формата даты создания в порядковый номер создания
# 20250104093011_<имя миграции>.sql -> 00001_<имя миграции>.sql
goose -dir ./internal/database/migration fix
```

### Проверка

```bash
# Вариант 1 (длинный)
goose -dir internal/database/migration postgres "postgresql://root:123456@127.0.0.1:5432/goan?sslmode=disable" status

# Вариант 2 (короткий)
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://root:123456@127.0.0.1:5432/goan?sslmode=disable

goose -dir internal/database/migration status
```

### Миграция

```bash
goose -dir internal/database/migration up
```

### Откат миграции

```bash
goose -dir internal/database/migration down
```
