package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

var last []byte

var index int

var colors = [][2]color.Attribute{
	{color.BgHiBlack, color.FgHiWhite},
	{color.BgHiRed, color.FgHiWhite},
	{color.BgHiGreen, color.FgHiBlack},
	{color.BgHiYellow, color.FgHiBlack},
	{color.BgHiBlue, color.FgHiBlack},
	{color.BgHiMagenta, color.FgHiWhite},
	{color.BgHiCyan, color.FgHiBlack},
	{color.BgHiWhite, color.FgHiBlack},
}

func main() {
	for _, f := range os.Args[1:] {
		if err := process(f); err != nil {
			log.Println(err)
		}
	}
}

func process(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, f); err != nil {
		return err
	}

	sum := hash.Sum(nil)

	// If this sum is different than the last, advance to the next color pair.
	if last != nil && !bytes.Equal(last, sum) {
		index++
		if index >= len(colors) {
			index = 0
		}
	}

	last = sum

	fmt.Printf(
		"%s : %s\n",
		color.New(colors[index][0], colors[index][1]).Sprintf("%x", sum),
		path,
	)

	return nil
}
