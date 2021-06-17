package main

import (
	"strconv"
)

type strArrayToInt struct {
	data  []string
	index int
	next  int
	err   error
}

func (s *strArrayToInt) init(sl []string) {
	s.data = sl
	s.index = 0
}

func (s *strArrayToInt) Scan() bool {
	if len(s.data) <= s.index {
		return false
	}
	d := s.data[s.index]

	i, err := strconv.Atoi(d)
	if err != nil {
		s.err = err
		return false
	}

	s.index += 1
	s.next = i
	return true
}

func (s *strArrayToInt) Int() int {
	return s.next
}

func (s *strArrayToInt) Err() error {
	return s.err
}

func main() {
	s := strArrayToInt{}
	strA := []string{"2", "4", "8", "16"}

	s.init(strA)
	for s.Scan() {
		println(s.Int())
	}

	if err := s.Err(); err != nil {
		println(err.Error())
	}

	println("===================")

	strB := []string{"2", "4", "GO", "16"}

	s.init(strB)
	for s.Scan() {
		println(s.Int())
	}

	if err := s.Err(); err != nil {
		println(err.Error())
	}
}
