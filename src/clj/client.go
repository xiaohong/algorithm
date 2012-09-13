package main

import (
	"fmt"
	"strings"
	"io"
)

func main(){
	r := strings.NewReader("(def a (123 aa (345 d34)))")
	fmt.Println(read(r))
}

type MacroReader  func (rune int,r io.RuneScanner) interface{}
var macros []MacroReader = make([]MacroReader, 256)
func init(){
	macros[int('(')] = seqReader 
}

func seqReader(rune int, r io.RuneScanner) interface{}{
	seq := make([]interface{},0)
	for {
		nrune,_,err := r.ReadRune()
		for string(nrune) == " " {
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

func read(r io.RuneScanner) interface{}{
	rune,_,err := r.ReadRune()
	for {
		if err != nil {
			return nil
		}
		for rune == ' ' {
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

func readNumber(b int, r io.RuneScanner) string{
	s := string(b)
	rune,_,err := r.ReadRune()
	for ; err == nil && string(rune) != " " && string(rune) != "(" && string(rune) != ")"  ; {
		s = s + string(rune)
		rune,_,err = r.ReadRune()
	}
	if err == nil {
		r.UnreadRune()
	}

	return s
}

func readObj(b int, r io.RuneScanner) string{
	
	return readNumber(b,r)
}
