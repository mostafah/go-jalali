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
		fmt.Fprintf(os.Stderr, "Converts Jalali Dates to Gregorian\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [13920204|1392/2/4|1392-02-04]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tWithout arguments it prints current date in Jalali Calendar\n\n")
	}

	flag.Parse()

	jdate := ""

	if !IsTerminal(syscall.Stdin) {
		in := []byte{}
		in, _, _ = bufio.NewReader(os.Stdin).ReadLine()
		jdate = string(in)
	} else {
		args := flag.Args()
		ln := len(args)
		if ln > 0 {
			jdate = args[0]
			if ln > 1 {
				jdate = strings.Join(args, " ")
			}
		}

	}

	if jdate == "" {
		now := time.Now()
		fmt.Printf("%d%02d%02d\n", now.Year(), int(now.Month()), now.Day())
		return
	}

	// parse the input
	rx := regexp.MustCompile(`\A((?:[1-9]\d)?(?:\d\d))(\D*)(0[1-9]|[1-9]|1[0-2])\D*([0-9]|0[1-9]|[1-2][0-9]|3[01])\z`)
	m := rx.FindStringSubmatch(jdate)
	if len(m) == 5 {
		year, _ := strconv.Atoi(m[1])
		sep := m[2]
		month, _ := strconv.Atoi(m[3])
		day, _ := strconv.Atoi(m[4])
		gdate := jalali.Jtog(year, month, day)
		fmt.Printf("%d%s%02d%s%02d\n", gdate.Year(), sep, int(gdate.Month()), sep, gdate.Day())
	} else {
		fmt.Fprintf(os.Stderr, "Invalid Jalali Date: '%s'\n", jdate)
	}

}
