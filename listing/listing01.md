Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```


```
Ответ:

[77 78 79]


Вывод будет таким, потому что в переменную b сохраняется срез с 1 по 4 индекс не включая его, то есть 1,2,3 с значениями 77 78 79
```



