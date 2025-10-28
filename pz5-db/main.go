package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
)

func main() {
    // Загружаем переменные окружения из .env файла
    _ = godotenv.Load()

    // Получаем DSN из переменных окружения или используем fallback
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        // Fallback - прямой DSN в коде (только для учебного стенда!)
        dsn = "postgres://postgres:your_password@localhost:5432/todo?sslmode=disable"
    }

    // Подключаемся к базе данных
    db, err := openDB(dsn)
    if err != nil {
        log.Fatalf("openDB error: %v", err)
    }
    defer db.Close()

    repo := NewRepo(db)

    // 1) Вставляем несколько задач
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    titles := []string{"Сделать ПЗ №5", "Купить кофе", "Проверить отчёты"}
    for _, title := range titles {
        id, err := repo.CreateTask(ctx, title)
        if err != nil {
            log.Fatalf("CreateTask error: %v", err)
        }
        log.Printf("Inserted task id=%d (%s)", id, title)
    }

    // 2) Прочитаем список всех задач
    ctxList, cancelList := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancelList()

    tasks, err := repo.ListTasks(ctxList)
    if err != nil {
        log.Fatalf("ListTasks error: %v", err)
    }

    // 3) Напечатаем все задачи
    fmt.Println("=== Все задачи ===")
    for _, t := range tasks {
        fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
            t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
    }

    // ========== ПРОВЕРОЧНЫЕ ЗАДАНИЯ ==========

    // Задание 1: ListDone - фильтрация по статусу выполнения
    fmt.Println("\n=== Задание 1: Фильтрация по статусу выполнения ===")
    
    // Помечаем одну задачу как выполненную для демонстрации
    err = repo.MarkDone(ctxList, 1)
    if err != nil {
        log.Printf("MarkDone error: %v", err)
    } else {
        fmt.Println("Задача #1 помечена как выполненная")
    }

    // Показываем невыполненные задачи
    undoneTasks, err := repo.ListDone(ctxList, false)
    if err != nil {
        log.Printf("ListDone error: %v", err)
    } else {
        fmt.Println("Невыполненные задачи:")
        for _, t := range undoneTasks {
            fmt.Printf("  #%d | %s\n", t.ID, t.Title)
        }
    }

    // Показываем выполненные задачи
    doneTasks, err := repo.ListDone(ctxList, true)
    if err != nil {
        log.Printf("ListDone error: %v", err)
    } else {
        fmt.Println("Выполненные задачи:")
        if len(doneTasks) == 0 {
            fmt.Println("  Пока нет выполненных задач")
        } else {
            for _, t := range doneTasks {
                fmt.Printf("  #%d | %s\n", t.ID, t.Title)
            }
        }
    }

    // Задание 2: FindByID - поиск по ID
    fmt.Println("\n=== Задание 2: Поиск задачи по ID ===")
    
    // Ищем задачу с ID=2
    taskID := 2
    foundTask, err := repo.FindByID(ctxList, taskID)
    if err != nil {
        log.Printf("FindByID error: %v", err)
    } else {
        fmt.Printf("Найдена задача #%d:\n", taskID)
        fmt.Printf("  Заголовок: %s\n", foundTask.Title)
        fmt.Printf("  Выполнена: %v\n", foundTask.Done)
        fmt.Printf("  Создана: %s\n", foundTask.CreatedAt.Format(time.RFC3339))
    }

    // Пробуем найти несуществующую задачу
    nonExistentID := 999
    _, err = repo.FindByID(ctxList, nonExistentID)
    if err != nil {
        fmt.Printf("Задача #%d не найдена (ожидаемо): %v\n", nonExistentID, err)
    }

    // Задание 3: CreateMany - массовая вставка через транзакцию
    fmt.Println("\n=== Задание 3: Массовая вставка задач ===")
    batchTitles := []string{
        "Массовая задача 1",
        "Массовая задача 2", 
        "Массовая задача 3",
    }

    ctxBatch, cancelBatch := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancelBatch()

    err = repo.CreateMany(ctxBatch, batchTitles)
    if err != nil {
        log.Printf("CreateMany error: %v", err)
    } else {
        fmt.Println("Успешно добавлено 3 задачи одной транзакцией")
        
        // Показываем обновленный список
        finalTasks, err := repo.ListTasks(ctxList)
        if err != nil {
            log.Printf("ListTasks error: %v", err)
        } else {
            fmt.Println("\n=== Финальный список всех задач ===")
            for _, t := range finalTasks {
                fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
                    t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
            }
        }
    }

    fmt.Println("  потенциальных проблем с сетью или БД, при этом не слишком часто")
}
