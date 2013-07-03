package clj

import "unicode/utf8"

type Reader interface{
	Read() rune
	Unread(v rune)
}


type stringReader struct{
	val []byte
	i int
}

func (this *stringReader) Read() rune{
	r, s := utf8.DecodeRune(this.val[this.i:])
	print(r)
	this.i += s
	return r
}

func (this *stringReader) Unread(v rune){
	this.i -= utf8.RuneLen(v)
}

func NewStringReader(v string) Reader{
	return &stringReader{val: []byte(v)}
}


