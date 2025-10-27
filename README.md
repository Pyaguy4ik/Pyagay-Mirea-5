<img width="969" height="725" alt="image" src="https://github.com/user-attachments/assets/3dd7eb53-1019-4bdf-9f55-7053e0868c66" />## Практическое задание 5
 
## Подключение к PostgreSQL через database/sql. Выполнение простых запросов (INSERT, SELECT)

Студент группы *ЭФМО-02-25 Пягай Даниил Игоревич*

## Описание

**Цели:**

1.	Установить и настроить PostgreSQL локально.
2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
3.	Выполнить параметризованные запросы INSERT и SELECT.
4.	Корректно работать с context, пулом соединений и обработкой ошибок.

## Открываем postgres и создаём базу данных (Linux)

![create_db](img/create_db.png)

### Создаём таблицу

![create_table](img/create_table.png)

### Выполняем быстрый тест

![insert](img/insert.png)

## Инициализация проекта

```bash
mkdir pz5-db && cd pz5-db
go mod init pz5-db
go get github.com/jackc/pgx/v5/stdlib
go get github.com/joho/godotenv
```

### Создаём структуру файлов

![structure](img/structure.png)

### Содержимое db.go
![db.go](img/db.go.png)

### Содержимое repository.go
![repository.go](img/repository.go.png)

### Содержимое main.go (фрагмент)
![main.go](img/main.go(1).png)
![main.go](img/main.go(2).png)

### Запуск и проверка
![go.run](img/run.go.png)
## Ожидаемый лог был получен

### Результаты выполнения проверочный заданий
![run.go](img/run.go.2.png)

### Обновить (статус первого задания меняем с false на true)
![update](img/update.png)

### Проверяем изменение
![run.go.3](img/run.go.3.png)

### Полученные результаты
![Answers](img/Answers.png)
