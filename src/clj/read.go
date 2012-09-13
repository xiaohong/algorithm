package main

import (
	"fmt"
	"strings"
	"io"
)

func main(){
	r := strings.NewReader("(def a 12.33)")
	fmt.Println(read(r))
}

type MacroReader  func (rune int,r io.RuneScanner) interface{}
var macros []MacroReader = make([]MacroReader, 256)
func init(){
	macros['('] = seqReader 
	macros['"']      = stringReader
}

func isWhiteSpace(rune int) bool{
	return rune == ' ' || rune == ','
}

func seqReader(rune int, r io.RuneScanner) interface{}{
	seq := make([]interface{},0)
	for {
		nrune,_,err := r.ReadRune()
		for isWhiteSpace(nrune) {
			nrune,_,err = r.ReadRune()
		}
		if err != nil || string(nrune) == ")"{
		    return seq
		}
		r.UnreadRune()
		fmt.Println(seq)
		seq = append(seq,read(r))
	}
	return seq
}

func stringReader(flag int, r io.RuneScanner) interface{}{
	var s string = ""
	for {
		rune,_,err := r.ReadRune()
		if err != nil {
			panic("err")
		}
		if rune == '"' {
			return s
		}
		s = s + string(rune)
	}
	return s
}

func read(r io.RuneScanner) interface{}{
	rune,_,err := r.ReadRune()
	for {
		if err != nil {
			return nil
		}
		for ; isWhiteSpace(rune); {
			rune,_,err = r.ReadRune()
		}
		if isDigit(rune) {
			return readNumber(rune, r)
		}
		if rune < 256 {

			m := macros[int(rune)]
			if m != nil {
				return m(rune,r)
			}
			
		}
		return	readObj(rune, r)
	}
	return nil
}

func isDigit(i int) bool {
	return i < int('9') && i > int('0')
}

func readRune(r io.RuneScanner) int {
	rune,_,err := r.ReadRune()
	if err != nil {
		return -1
	}
	return rune
}
func readNumber(b int, r io.RuneScanner) interface{}{
	s := string(b)

	for {
		rune := readRune(r)
		if (rune == -1 || isWhiteSpace(rune) || rune == ')'){
			r.UnreadRune()
			return s
		}
		s = s + string(rune)
	}
	
	return matchNumer(s)
}

func matchNumer(s string) interface{}{
	return s
}

func readObj(b int, r io.RuneScanner) interface{}{
	
	return readNumber(b,r)
}
