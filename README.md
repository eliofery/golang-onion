# Название проекта

Краткое писание проекта.

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
