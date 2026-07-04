# Задание 1 — Методы и String()
Добавь в models/models.go три метода:

**1.** На Feed — метод String() string, который возвращает строку вида:
```
[Feed] Golang Weekly — https://golangweekly.com/rss/
```
**2.** На Article — метод String() string, который возвращает:
```
[Article] Gin: 12 years... (read: false)
```
**3.** На Article — метод MarkAsRead(), который устанавливает IsRead = true. Подумай: нужен ли здесь указатель-получатель?

В main.go замени fmt.Printf(...) на вызовы fmt.Println(feed) и fmt.Println(article) — Go автоматически вызовет String() если он есть. И добавь вызов MarkAsRead() на первой статье, затем выведи её снова, чтобы убедиться что IsRead изменился.

**Подсказка по форматированию строк:** fmt.Sprintf("[Feed] %s — %s", f.Title, f.URL) работает как Printf, но возвращает строку вместо печати.
