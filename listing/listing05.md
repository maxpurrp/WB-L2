Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

В данном примере функция test() возвращает nil, но это указатель на customError. Хоть строкой переменная err объявлена как общий интерфейс `error`, после выполнения test() происходит преобразование типа и переменная `err` становится типом 'error'.  И так как перменная `err` теперь является указателем, условие истинно, выводим в stdout `error`.

```
