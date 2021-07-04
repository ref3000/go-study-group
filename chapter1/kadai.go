package chapter1

import (
	"bufio"
	"fmt"
	"github.com/apbgo/go-study-group/chapter1/lib"
	"math"
	"os"
	"strconv"
)

// Calc opには+,-,×,÷の4つが渡ってくることを想定してxとyについて計算して返却(正常時はerrorはnilでよい)
// 想定していないopが渡って来た時には0とerrorを返却
func Calc(op string, x, y int) (int, error) {

	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	result := 0
	var err error = nil

	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "×":
		result = x * y
	case "÷":
		result = x / y
	default:
		err = fmt.Errorf("invalid op=%s", op)
	}

	return result, err
}

// StringEncode 引数strの長さが5以下の時キャメルケースにして返却、それ以外であればスネークケースにして返却
func StringEncode(str string) string {
	// ヒント：長さ(バイト長)はlen(str)で取得できる
	// chapter1/libのToCamelとToSnakeを使うこと

	result := ""
	if len(str) <= 5 {
		result = lib.ToCamel(str)
	} else {
		result = lib.ToSnake(str)
	}

	return result
}

// Sqrt 数値xが与えられたときにz²が最もxに近い数値zを返却
func Sqrt(x float64) float64 {

	z := 1.0
	const EPS = 1e-12

	for {
		tmp := z - (z*z-x)/(2*z)
		if math.Abs(z-tmp) < EPS {
			return z
		}
		z = tmp
	}

}

// Pyramid x段のピラミッドを文字列にして返却
// 期待する戻り値の例：x=5のとき "1\n12\n123\n1234\n12345"
// （x<=0の時は"error"を返却）
func Pyramid(x int) string {
	// ヒント：string <-> intにはstrconvを使う
	// int -> stringはstrconv.Ioa() https://golang.org/pkg/strconv/#Itoa

	result := ""

	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			result += strconv.Itoa(j)
		}
		if i < x {
			result += "\n"
		}
	}

	return result
}

// StringSum x,yをintにキャストし合計値を返却 (正常終了時、errorはnilでよい)
// キャスト時にエラーがあれば0とエラーを返却
func StringSum(x, y string) (int, error) {

	// ヒント：string <-> intにはstrconvを使う
	// string -> intはstrconv.Atoi() https://golang.org/pkg/strconv/#Atoi

	xInt, err := strconv.Atoi(x)
	if err != nil {
		return 0, fmt.Errorf("failed to parse. x: %s", x)
	}

	yInt, err := strconv.Atoi(y)
	if err != nil {
		return 0, fmt.Errorf("failed to parse. y: %s", y)
	}

	return xInt + yInt, nil
}

// SumFromFileNumber ファイルを開いてそこに記載のある数字の和を返却
func SumFromFileNumber(filePath string) (int, error) {
	// ヒント：ファイルの扱い：os.Open()/os.Close()
	// bufio.Scannerなどで１行ずつ読み込むと良い

	fp, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file. filePath: %s", filePath)
	}
	defer fp.Close()

	sum := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("failed to parse.")
		}
		sum += num
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sum, nil
}
