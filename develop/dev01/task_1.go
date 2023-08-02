package main

// Создать программу печатающую точное время с использованием NTP -
// библиотеки. Инициализировать как go module. Использовать
// библиотеку github.com/beevik/ntp. Написать программу печатающую
// текущее время
// точное время с использованием этой библиотеки.
// Требования:
// Программа должна быть оформлена как go module
// Программа должна корректно обрабатывать ошибки библиотеки:
// выводить их в STDERR и возвращать ненулевой код выхода в OS

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

type facade interface {
	getCurrentTime()
}

func getCurrentTime(address string) string {
	ntpTime, err := ntp.Time(address)
	if err != nil {
		// Код ошибки еще надо откудото взять
		// возвращать ненулевой код выхода в OS
		fmt.Fprintf(os.Stderr, "log msg: %s", err)
	}
	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Fprintf(os.Stdout, "%s\n", ntpTimeFormatted)
	return ntpTimeFormatted
}

func main() {
	for {
		getCurrentTime("pool.ntp.org")
		time.Sleep(1 * time.Second)
	}
}
