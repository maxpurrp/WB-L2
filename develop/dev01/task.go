package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===
Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и go lint.
*/

// the entry point of program
func main() {
	// extract current time from ntp server
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// handle error if not possible get time from ntp
		// and print error in stderr with code err 1
		fmt.Fprintln(os.Stderr, "Error receiving time from NTP :", err)
		os.Exit(1)
	}
	// parse hour, min, seconds and nanosecond from ntpTime
	hour, min, second, nanosecond := ntpTime.Hour(), ntpTime.Minute(), ntpTime.Second(), ntpTime.Nanosecond()
	// Print the current time in the desired format.
	fmt.Printf("Current time : %02d:%02d:%02d.%09d", hour, min, second, nanosecond)
}
