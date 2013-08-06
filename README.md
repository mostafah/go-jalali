go-jalali
=======

go-jalali provides some tools to work with Jalali Calendar in golang

There's also `j2g` and `g2j` commands to convert Jalali to Gregorian and vice versa.

## Documentation

To view `go doc` style documentation of this project use GoDoc site
[here](http://godoc.org/github.com/nbjahan/go-jalali/jalali)

You can also view the documentation locally once the package is installed with
the `godoc` tool by running `godoc -http=":6060"` and pointing your browser to
http://localhost:6060/pkg/github.com/nbjahan/go-jalali/jalali

## Installation

```bash
$ go get github.com/nbjahan/go-jalali/jalali/...
```

## Quick Start

To create `time.Time` from Jalali date:

```Go
var d time.Time
d = Jtog(1392, 4, 15)               // "2013-07-06"
```

Optionally you can specify hour, min, sec, nsec:

```Go
d = Jtog(1392, 4, 15, 17, 3, 32, 0) // "2013-07-06 17:03:32"
```

To convert Gregorian date to Jalali date:

```Go
year, month, day := Gtoj(time.Now())
```

There is also `Strftime` so you can convert to Jalali and format:

```Go
jstring := jalali.Strftime("Printed on %Y/%m/%d", time.Now()) // "Printed on 1392/04/02"
```

[format flags](http://godoc.org/github.com/nbjahan/go-jalali/jalali#Strftime)

Checking the Leap year:

```Go
isLeap := jalali.IsLeap(time.Now())
```
