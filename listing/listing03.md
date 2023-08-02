Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Создается интерфейс в которм храняться данные и ссылки на тип, методанный для возможности динамической типизации. И не смотря на то что там нуль, внутренняя структура остается прежней просто в качестве и типа и данных устанавливается nil, что заведомо делает егон не равным простому nil.

```