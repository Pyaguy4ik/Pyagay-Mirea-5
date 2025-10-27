## Подключение к PostgreSQL через database/sql. Выполнение простых запросов (INSERT, SELECT)

Студент группы *ЭФМО-02-25 Пягай Даниил Игоревич*

## Описание

**Цели:**

1.	Установить и настроить PostgreSQL локально.
2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
3.	Выполнить параметризованные запросы INSERT и SELECT.
4.	Корректно работать с context, пулом соединений и обработкой ошибок.

## Открываем postgres и создаём базу данных (Linux)

![create_db](img/create_db.JPG)

### Создаём таблицу

![create_table](img/create_table.JPG)

### Выполняем быстрый тест

![insert](img/insert.JPG)

## Инициализация проекта

```bash
mkdir pz5-db && cd pz5-db
go mod init pz5-db
go get github.com/jackc/pgx/v5/stdlib
go get github.com/joho/godotenv
```

### Создаём структуру файлов

![structure](img/structure.JPG)

### Содержимое db.go
![db.go](img/db.go.JPG)

### Содержимое repository.go
![repository.go](img/repository.go.JPG)

### Содержимое main.go (фрагмент)
![main.go](img/main.go(1).JPG)
![main.go](img/main.go(2).JPG)

### Запуск и проверка
![go.run](img/run.go.JPG)
## Ожидаемый лог был получен

### Результаты выполнения проверочный заданий
![run.go](img/run.go.JPG)

### Обновить (статус первого задания меняем с false на true)
![update](img/update.JPG)

### Проверяем изменение
![run.go.3](img/run.go_3.JPG)
