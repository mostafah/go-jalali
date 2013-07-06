package jalali

import (
	"fmt"
	"time"
)

func ExampleJtog() {
	d := Jtog(1392, 4, 15)
	fmt.Println(d.Format("2006-01-02 15:04:05.999999999"))
	// Output:
	// 2013-07-06 00:00:00
}

func ExampleJtog_hour() {
	d := Jtog(1392, 4, 15, 17)
	fmt.Println(d.Format("2006-01-02 15:04:05.999999999"))
	// Output:
	// 2013-07-06 17:00:00
}

func ExampleJtog_hourMin() {
	d := Jtog(1392, 4, 15, 17, 3)
	fmt.Println(d.Format("2006-01-02 15:04:05.999999999"))
	// Output:
	// 2013-07-06 17:03:00
}

func ExampleJtog_hourMinSec() {
	d := Jtog(1392, 4, 15, 17, 3, 32)
	fmt.Println(d.Format("2006-01-02 15:04:05.999999999"))
	// Output:
	// 2013-07-06 17:03:32
}

func ExampleJtog_hourMinSecNano() {
	d := Jtog(1392, 4, 15, 17, 3, 32, 0)
	fmt.Println(d.Format("2006-01-02 15:04:05.999999999"))
	// Output:
	// 2013-07-06 17:03:32
}

func ExampleGtoj() {
	year, month, day := Gtoj(time.Date(2013, 7, 3, 17, 14, 0, 0, time.Local))
	fmt.Printf("%d-%02d-%02d", year, month, day)
	// Output:
	// 1392-04-12
}
