Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
timeout running program
1
2
4
6
8
3
5
0
0
0
0
...

В функцию передаются числовые значения. Создается канал и в горутине перебираются переданные в функцию
 числа. На каждей итерации в канал пишется 1 число и засыпает на рендомное количество милисекунд. По 
 прохождению цикла канал закрывается. go playground убивает процесс раньше чем следует и последнее значение иногда теряется, как 7. В отдельной горутине был создан канал для чтения, который возвращал числа из двух 
 каналов пока те были открыты, чтение из закрытого канала сперва вернуло то что было записано но не 
 успело быть прочитано, затем зациклилось возвращая 0.
```
