package chapter2

import "fmt"

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func Calc(slice []int) (int, error) {
	// TODO Q1
	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	len := len(slice)

	var answer int
	switch {
	case len <= 0:
		return 0, fmt.Errorf("slice length must be greater than 0.")
	case len == 1:
		answer = slice[0]
	case len == 2:
		answer = slice[0] * slice[1]
	case len >= 3:
		sum := 0
		for _, x := range slice {
			sum += x
		}
		answer = sum
	}
	return answer, nil
}

type Number struct {
	index int
}

// 構造体Numberを3つの要素数から成るスライスにして返却
// 3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
func Numbers() []Number {
	// TODO Q2

	var slice []Number

	slice = append(slice, Number{index: 1})
	slice = append(slice, Number{index: 2})
	slice = append(slice, Number{index: 3})

	return slice
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
func CalcMap(m map[string]int) int {
	// TODO Q3

	sum := 0
	for k, v := range m {
		if k != "yon" {
			sum += v
		}
	}

	return sum
}

type Model struct {
	Value int
}

// 与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
func Add(models []Model) {
	// TODO  Q4

	for k, v := range models {
		models[k].Value = v.Value + 5
	}
}

// 引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
// 順序はスライスに格納されている順番のまま返却すること
// ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
func Unique(slice []int) []int {
	// TODO Q5

	var result []int

	memo := make(map[int]bool)
	for _, v := range slice {
		if !memo[v] {
			memo[v] = true
			result = append(result, v)
		}
	}

	return result
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func Fibonacci() func() int {
	// TODO Q6 オプション

	a := 0
	b := 1

	return func() int {
		result := a
		a = b
		b = result + b
		return result
	}
}
