package chapter5

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var delimiter = flag.String("d", ",", "区切り文字を指定してください")
var fields = flag.Int("f", 1, "フィールドの何番目を取り出すか指定してください")

func Validate(args []string, fields int) error {
	if len(args) == 0 {
		return fmt.Errorf("ファイルパスを指定してください。")
	}
	if fields <= 0 {
		return fmt.Errorf("-f は1以上である必要があります")
	}
	return nil
}

func Cut(src io.Reader, dst io.Writer, delimiter string, fields int) error {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		text := scanner.Text()
		sb := strings.Split(text, delimiter)
		if len(sb) < fields {
			return fmt.Errorf("-fの値に該当するデータがありません")
		}
		s := sb[fields-1]
		fmt.Fprintln(dst, s)
	}
	return scanner.Err()
}

func cut() {
	flag.Parse()

	// このValidationを関数1つ目に切り出す ---------
	// ヒント：flagの内容を渡してやって、バリデーションし、エラーがあれば返すような関数にできる
	if err := Validate(flag.Args(), *fields); err != nil {
		log.Fatal(err)
	}
	// ---------------------------------------

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// この部分をCutコマンドとして関数2つ目に切り出す------
	// ヒント：NewScannerにfileを渡しているが、NewScannerはio.Readerであれば何でも良い
	// また、出力も現在fmt.Println(s)にしているが、io.Writerを使って書き出す先を指定できるようにしてやる
	// 関数の引数で読み出すio.Readerと、
	// 書き出すio.Writer (本関数からはos.Stdout, テストからはbyte.Bufferなどへ)を指定できるようにすると良い
	if err := Cut(file, os.Stdout, *delimiter, *fields); err != nil {
		log.Fatal(err)
	}
	// ------------------------------------------------
}
