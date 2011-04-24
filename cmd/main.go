package main

import (
	"flag"
	"fmt"
	"github.com/dchest/captcha"
	"io"
	"log"
	"os"
)

var (
	flagImage = flag.Bool("i", true, "output image captcha")
	flagAudio = flag.Bool("a", false, "output audio captcha")
	flagLen   = flag.Int("len", captcha.StdLength, "length of captcha")
	flagImgW  = flag.Int("width", captcha.StdWidth, "image captcha width")
	flagImgH  = flag.Int("height", captcha.StdHeight, "image captcha height")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: captcha [flags] filename\n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	fname := flag.Arg(0)
	if fname == "" {
		usage()
		os.Exit(1)
	}
	f, err := os.Create(fname)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer f.Close()
	var w io.WriterTo
	d := captcha.RandomDigits(*flagLen)
	switch {
	case *flagAudio:
		w = captcha.NewAudio(d)
	case *flagImage:
		w = captcha.NewImage(d, *flagImgW, *flagImgH)
	}
	_, err = w.WriteTo(f)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(d)
}