package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/nbjahan/go-jalali/jalali"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

// IsTerminal returns true if the given file descriptor is a terminal.
// Stolen from http://code.google.com/p/go/source/browse/ssh/terminal/util.go?repo=crypto#31
func IsTerminal(fd int) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), syscall.TIOCGETA, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Converts Gregorian Dates to Jalali\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [20130623|2013/6/23|2013-6-23]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tWithout arguments it prints current date in Gregorian Calendar\n\n")
	}

	flag.Parse()

	gdate := ""

	if !IsTerminal(syscall.Stdin) {
		in := []byte{}
		in, _, _ = bufio.NewReader(os.Stdin).ReadLine()
		gdate = string(in)
	} else {
		args := flag.Args()
		ln := len(args)
		if ln > 0 {
			gdate = args[0]
			if ln > 1 {
				gdate = strings.Join(args, " ")
			}
		}

	}
	if gdate == "" {
		fmt.Println(jalali.Strftime("%Y%m%d", time.Now()))
		return
	}

	// parse the input
	rx := regexp.MustCompile(`\A((?:[1-9]\d)?(?:\d\d))(\D*)(0[1-9]|[1-9]|1[0-2])\D*([0-9]|0[1-9]|[1-2][0-9]|3[01])\z`)
	m := rx.FindStringSubmatch(gdate)
	// fmt.Println(match)
	if len(m) == 5 {
		// m := match[0]
		year, _ := strconv.Atoi(m[1])
		sep := m[2]
		month, _ := strconv.Atoi(m[3])
		day, _ := strconv.Atoi(m[4])
		jy, jm, jd := jalali.Gtoj(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local))
		fmt.Printf("%d%s%02d%s%02d\n", jy, sep, jm, sep, jd)
	} else {
		fmt.Fprintf(os.Stderr, "Invalid Gregorian Date: '%s'\n", gdate)
	}

}
