# Задание 0 — Каркас проекта
Создай проект goreader со следующей структурой:
```
goreader/
├── go.mod
├── main.go
└── models/
    └── models.go
```
## Шаг 1. Создай папку и инициализируй модуль:
```
mkdir goreader && cd goreader
go mod init goreader
```

## Шаг 2. В файле models/models.go объяви два типа:
Feed (RSS-лента) с полями:
 - ID — int
- Title — string
- URL — string
- Description — string

Article (статья) с полями:

- ID — int
- FeedID — int (к какому фиду относится)
- Title — string
- URL — string
- Content — string
- IsRead — bool

## Шаг 3. В models/models.go напиши функцию:
```
func NewFeed(title, url, description string) Feed
```
Она принимает три строки и возвращает готовый Feed. Поле ID пока ставь в 0 — потом им займётся база данных.

## Шаг 4. В main.go напиши main, которая:

1. Создаёт фид через NewFeed (придумай любой RSS, например "Hacker News")
2. Создаёт вручную (через Article{...}) две любые статьи с FeedID = 0
3. Кладёт их в срез []Article
4. Выводит в консоль название фида и заголовок каждой статьи через цикл

**Запуск:**
```
go run main.go
```
