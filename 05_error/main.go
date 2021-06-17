package main

import (
	"errors"
	"fmt"
	"strconv"
)

type strArrayToInt struct {
	data  []string
	index int
	next  int
	err   error
}

type NumLimitErr struct{}

func (n NumLimitErr) Error() string {
	return "要素は6つまでです！"
}

func (s *strArrayToInt) init(sl []string) {
	s.data = sl
	s.index = 0
}

func (s *strArrayToInt) Scan() bool {
	if len(s.data) > 6 {
		// エラーA
		s.err = NumLimitErr{}
	}

	if len(s.data) <= s.index {
		return false
	}
	d := s.data[s.index]

	i, err := strconv.Atoi(d)
	if err != nil {
		// エラーB
		s.err = fmt.Errorf("StrArrayToInt: INTへの変換に失敗しました %w", err)
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

		if errors.As(&strconv.NumError{}, &err) {
			println("数字の変換エラーです")
		}
	}

	println("===================")

	strC := []string{"2", "4", "8", "16", "2", "4", "8", "16"}

	s.init(strC)
	for s.Scan() {
		println(s.Int())
	}

	if err := s.Err(); err != nil {
		println(err.Error())

		if errors.As(NumLimitErr{}, &err) {
			// エラーはエラーでも 要素が多かった時にしか実行されない
			println(err.Error())
			// 配列の要素を減らす処理とか？
		}

		if errors.As(&strconv.NumError{}, &err) {
			// エラーはエラーでも 数字にテキストが混じっていた時にしか実行されない
			println("数字の変換エラーです")
			// パニックするとか？
		}
	}
}
