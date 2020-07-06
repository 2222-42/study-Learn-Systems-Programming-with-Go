package main

//使っていいのはio パッケージの内容＋基本文法のみです
import (
	"bytes"
	"io"
	"os"
	"strings"
)

//次の3つの文字列を3つの入力ストリーム（io.Reader）とし、下記に示すmain()
//関数のコメント部にコードを追加して、最後のio.Copy() で「ASCII」の文字列が
//出力されるようにしてみてください。

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

//文字列リテラルを使用してはいけません。
// コメント部以外を変更してはいけません。当然、import するパッケージを増や
//してはいけません。
func main() {
	var stream io.Reader

	ra := io.NewSectionReader(programming, 5, 1)
	rs := io.NewSectionReader(system, 0, 1)
	rc := io.NewSectionReader(computer, 0, 1)
	ri := io.NewSectionReader(programming, 8, 1)

	var buffer bytes.Buffer
	ri2 := io.TeeReader(ri, &buffer)
	stream = io.MultiReader(ra, rs, rc, ri2, &buffer)

	//stream = io.MultiReader(ra, rs, rc, ri, ri)
	//ASCIってなっちゃう。
	// -> 同一のReaderだからReaderからなくなったから

	// MultiWriterとPipeを使うケースも考えられる。

	io.Copy(os.Stdout, stream)
}
