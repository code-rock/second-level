Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
 2
 1

Defer вызывается после выполнения функции. defer работает после return.Это в частности позволяет менять из defer возвращаемое значение. Если их много складывается в стек и вызывается в обратном
порядке от поступления.
в первом случае создается замыкание на переменную возращаемую функией и при вызове дефер она будет
мутировать переменную и вернет ее в итоге.
Во втором случае х копируется в ретерн в значении 1. 
```
