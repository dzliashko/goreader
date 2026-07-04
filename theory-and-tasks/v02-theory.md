# Урок 2. Теория — Интерфейсы и обработка ошибок
## Интерфейсы
Интерфейс описывает **поведение** — набор методов, которые должен реализовать тип. Сам тип при этом ничего не "декларирует" — если у него есть нужные методы, он **автоматически** реализует интерфейс.
```
type Stringer interface {
    String() string
}
```
Именно этот интерфейс (```fmt.Stringer```) ты уже реализовал на ```Feed``` и ```Article``` — поэтому ```fmt.Println``` и вызывает твой ```String()``` автоматически.
Зачем интерфейсы? Чтобы писать функции, которые работают **с любым типом** с нужным поведением:
```
func PrintAll(items []fmt.Stringer) {
    for _, item := range items {
        fmt.Println(item.String())
    }
}
// Можно передать и []Feed, и []Article — главное что у них есть String()
```

## Тип error

В Go error — это тоже интерфейс:
```
type error interface {
    Error() string
}
```
Поэтому любой тип с методом ```Error() string``` является ошибкой. Это значит, что можно создавать **свои типы ошибок** с дополнительными полями:
```
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("поле %s: %s", e.Field, e.Message)
}
```
Использование:
```
func validateURL(url string) error {
    if url == "" {
        return &ValidationError{Field: "URL", Message: "не может быть пустым"}
    }
    return nil
}

err := validateURL("")
if err != nil {
    fmt.Println(err) // "поле URL: не может быть пустым"
}
```
