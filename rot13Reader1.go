package main

import (
	"io"
	"os"
	"strings"
)

var rot13 map[string]string = {
	"a": "m",
	"b": "n",
	"c": "o",
	"d": "p",
	"e": "q",
	"f": "r",
	"g": "s",
	"h": "t",
	"i": "u",
	"j": "v",
	"k": "w",
	"l": "x",
	"m": "y",
	"n": "z",
	"o": "a",
	"p": "b",
	"q": "c",
	"r": "d",
	"s": "e",
	"t": "f",
	"u": "g",
	"v": "h",
	"w": "i",
	"x": "j",
	"y": "k",
	"z": "l",
	"A": "M",
	"B": "N",
	"C": "O",
	"D": "P",
	"E": "Q",
	"F": "R",
	"G": "S",
	"H": "T",
	"I": "U",
	"J": "V",
	"K": "W",
	"L": "X",
	"M": "Y",
	"N": "Z",
	"O": "A",
	"P": "B",
	"Q": "C",
	"R": "D",
	"S": "E",
	"T": "F",
	"U": "G",
	"V": "H",
	"W": "I",
	"X": "J",
	"Y": "K",
	"Z": "L",
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

