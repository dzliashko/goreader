# Урок 0. Основы Go — теория
## Что такое Go?
Go — это язык от Google. Он компилируемый (код превращается в бинарник), строго типизированный (нельзя смешать число и строку), и намеренно простой — в нём мало конструкций, но каждая мощная.

## 1. Пакеты и точка входа
Каждый файл Go начинается с объявления пакета. Программа стартует с функции main в пакете main:
package main
```
import "fmt"

func main() {
    fmt.Println("Привет, Goreader!")
}
```
import подключает стандартные библиотеки. fmt — форматированный вывод (от format).

## 2. Переменные и типы
```
// Длинный способ — явный тип
var name string = "Goreader"
var count int = 42
var active bool = true

// Короткий способ — Go сам поймёт тип (только внутри функции)
name := "Goreader"
count := 42
active := true
Основные типы: string, int, float64, bool.
```

## 3. Структуры (struct)
Это главный инструмент Go для группировки данных — аналог "объекта" без лишнего.
```
type Article struct {
    Title   string
    URL     string
    Content string
    IsRead  bool
}
```
Создание и доступ к полям:
```
a := Article{
    Title:  "Первая статья",
    URL:    "https://example.com/1",
    IsRead: false,
}

fmt.Println(a.Title) // "Первая статья"
a.IsRead = true
```

## 4. Функции
```
func add(a int, b int) int {
    return a + b
}

// Если типы одинаковые — можно короче
func add(a, b int) int {
    return a + b
}
```
Go уникален тем, что функция может вернуть несколько значений. Это используется для ошибок:
```
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("нельзя делить на ноль")
    }
    return a / b, nil  // nil = "ошибки нет"
}
```

## 5. Обработка ошибок
В Go нет исключений. Ошибка — это просто второе возвращаемое значение:
```
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Ошибка:", err)
    return
}
fmt.Println(result)
```
Это выглядит многословно, но зато явно — ты никогда не пропустишь ошибку случайно.

## 6. Срезы (slice) и цикл for
Срез — это динамический массив:
```
articles := []Article{}  // пустой срез

// Добавить элемент:
articles = append(articles, Article{Title: "Новость"})

// Перебрать:
for i, a := range articles {
    fmt.Println(i, a.Title)
}
Если индекс не нужен — заменяем на _:
gofor _, a := range articles {
    fmt.Println(a.Title)
}
```
