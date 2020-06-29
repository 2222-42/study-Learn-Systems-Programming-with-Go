# chapter 03 低レベルアクセスへの入口2： io.Reader

前章のio.Writerの対となるio.Readerを中心に、仲間のインターフェースを紹介。

書き込みに比べると読み込みのほうがトピックが多く、機能も多いため、本章の分量は多め

- io.Readerとその仲間たちの紹介(3.1)
- 少ないコード量でio.Reader からデータを効率よく読み込むための補助的な関数群の紹介(3.2)
- io.Readerとio.Writer以外の入出力インタフェースを紹介(3.3)
- io.Reader を満たす構造体で特に頻繁に使われるものの紹介（標準入力、ファイル、ソケット、メモリのバッファ）(3.4)
- バイナリ解析に便利な機能群の紹介(3.5)
- テキスト解析に便利な機能群の紹介(3.6)
- ちょっと抽象的なio.Readerの構造体の紹介(3.7)

## 3.1 io.Reader

出力先の種類（ファイルか、画面か、バッファか、ネットワークか）にかかわらず、
データを出力するという機能がGo言語のインタフェースという仕組みで抽象化されている

プログラムで外部からデータを読み込むための機能もGo 言語のインタフェースとして抽象化されている

io.Readerには、次のような形式のRead()メソッドが宣言

```
type Reader interface {
    func Read(p []byte) (n int, err error)
}
```

引数であるp は、読み込んだ内容を一時的に入れておくバッファ。あらかじめメモリを用意しておいて(make()を使うなどして)、それを使う。

読み込んだバイト数とエラーを返し、読み込んだ内容はpに格納する。

```
// 1024 バイトのバッファをmake で作る
buffer := make([]byte, 1024)
// size は実際に読み込んだバイト数、err はエラー
size, err := r.Read(buffer)
```

面倒なところ：
- バッファの管理をしつつ、何度もRead() メソッドを読んでデータを最後まで読み込むなど、読み込み処理を書くたびに同じようなコードを書かなければなりません

Goの対応：
- 低レベルなインタフェースだけでなく、それを簡単に扱うための機能も豊富に提供
- 次の節で、それらの補助機能を見る。

## 3.2 io.Readerの補助関数

io.Reader をそのまま使うのは多少不便なため、入力を扱うときは補助関数を使うことになります

Go言語では特別なもの以外はこのような外部のヘルパー関数を使って実現します。

### 3.2.1 読み込みの補助関数

`ioutil.ReadAll()`:

- これは終端記号にあたるまですべてのデータを読み込んで返します
  - メモリに収まらないかもしれないようなケースでは使えません
  
```
// すべて読み込む
buffer, err := ioutil.ReadAll(reader)
```

`io.ReadFull()`:

- 決まったバイト数だけ確実に読み込む
- 指定したバッファのサイズ分まで読み込めない場合にエラーが返る

```
// 4 バイト読み込めないとエラー
buffer := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)
```

(最低読み込みバイト数を指定しつつ、それ以上のデータも読む、`io.ReadAtLeast()` というものもある)

### 3.2.2 コピーの補助関数

`io.Reader`から`io.Writer`にそのままデータを渡したいときに使う

`io.Copy()`:

- すべてを読み尽くして書き込む

`io.CopyN()`:

- コピーするバイト数を指定できる

```
// すべてコピー
writeSize, err := io.Copy(writer, reader)
// 指定したサイズだけコピー
writeSize, err := io.CopyN(writer, reader, size)
```

`io.CopyBuffer()`:
- 自分で作った作業バッファを渡すことができる
　　- あらかじめコピーする量が決まっていて無駄なバッファを使いたくない場合や
　　- 何度もコピーするのでバッファを使いまわしたい場合

## 3.3 入出力に関するio.Writerとio.Reader以外のインタフェース

一般的な入出力関連インタフェース:

- io.Closerインタフェース
  - func Close() error メソッドを持ちます。
  - 使用し終わったファイルを閉じます。
- io.Seekerインタフェース
  - func Seek(offset int64, whence int) (int64, error) メソッドを持ちます。
  - 読み書き位置を移動します。
- io.ReaderAtインタフェース
  - func ReadAt(p []byte, off int64) (n int, err error) メソッドを持ちます。
  - 対象となるオブジェクトがランダムアクセスを行える場合に、好きな位置を自由にアクセスするときに使います

### 3.3.1 入出力関連の複合インタフェース

io.Closerやio.Seeker だけを満たした構造体を扱うことよりも、io.Readerやio.Writerを組み合わせたインタフェースを満たす構造体を使うことが多い

### 3.3.2 入出力関連インタフェースのキャスト

引数にio.ReadCloser が要求されているが、今あるオブジェクトはio.Readerしか満たしていない、ということもたまにあります

その場合はioutil.NopCloser() 関数を使うと、
ダミーのClose() メソッドを持ってio.ReadCloser のフリをする（ただしClose()しても何も起きない）ラッパーオブジェクトを得る。

```
var reader io.Reader = strings.NewReader(" テストデータ")
var readCloser io.ReadCloser = ioutil.NopCloser(reader)
```

バッファリングが入ってしまいますが、bufio.NewReadWriter() 関数を使うと、
個別のio.Reader とio.Writer をつなげて、io.ReadWriter 型のオブジェクトを作ることができます

```
var readWriter io.ReadWriter = bufio.NewReadWriter(reader, writer)
```

## 3.4 io.Readerを満たす構造体で、よく使うもの

io.Reader インタフェースを満たす構造体

構造体の種類ごとに、Read() メソッドの使い方を説明

実はGo言語の構造体の多くは、読みと書きの両方のインタフェースを満たしています。さらに、前項で説明した入出力関連の他のインタフェースを満たす場合もあります。
そのため、前章のio.Writerの説明に登場した構造体の多くは、io.Readerインタフェースも満たしています

### 3.4.1 標準入力

os.Stdin:
- io.Reader
- io.Closer

標準入力に対応するオブジェクトがos.Stdin です。このプログラムをそのまま実行すると入力待ちになり、以降はEnter が押されるたびに結果が返ってきます

プログラムを単体で実行すると、入力待ちでブロックしてしまいます。つまり、入力がくるまで実行が完全停止してしまう
- Go言語のRead() はタイムアウトのような仕組みもなく、このブロックを避けられません
- net.Connにはタイムアウトがある

Go言語の場合は並列処理機構が便利に使えるので、それを使ってノンブロッキングな処理を書きます

読み込んだ文字列を処理するコードにはチャネルという仕組みを使って渡すのが定石

os.Stdin の入力がキーボードに接続されているのか、上記の例のように他のプロセスに接続されているのかを判定する方法は、
プロセス周辺のトピックとして第11章「プロセスの役割とGo言語による操作」で紹介

### 3.4.2 ファイル入力

os.File:
- io.Reader
- io.Writer
- io.Seeker
- io.Closer

- os.Create() 関数: ファイルの新規作成
- os.Open() 関数: 既存のファイルを開く

フラグ違いのエイリアスで、同じシステムコールが呼ばれている：
```
func Open(name string) (*File, error) {
    return OpenFile(name, O_RDONLY, 0)
}
func Create(name string) (*File, error) {
    return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

ファイルを一度開いたらClose() する必要がある。

defer は、現在のスコープが終了したら、その後ろに書かれている行の処理を実行するので、確実に行う後処理を実行できる。

### 3.4.3 ネットワーク通信の読み込み

net.Conn:
- io.Reader
- io.Writer
- io.Closer

インターネット上でのデータのやり取り
- 送信データを送信者側から見ると書き込みで、
- 受信者側から見ると読み込み

HTTP を読み込むプログラムを開発するたびにRFC に従ってパース処理を実装するのは効率的ではありません

HTTPのレスポンスをパースするhttp.ReadResponse()関数を使おう
1. bufio.NewReader() 関数で、bufio.Reader でラップ
2. bufio.Reader でラップしたnet.Conn を渡すと、
3. http.Response構造体のオブジェクトが返されます

### 3.4.4 メモリに蓄えた内容をio.Readerとして読み出すバッファ

| 構造体 | io.Reader | io.Writer | io.Seeker | io.Closer | io.ReaderAt |
| --- | --- | --- | --- | --- | --- |
|bytes.Buffer | O | O | | | |
| bytes.Reader | O |  | O | | O |
| strings.Reader | O | | O | | O |

io.Reader としても使える
- 書き込まれた内容をメモリに保持しておくbytes.Buffer
- bytes.Reader
- strings.Reader

ただし、バイナリデータの解析に使うio.SectionReader だけは、
io.Reader ではなく
io.ReaderAt というちょっと違うインタフェースのReader を必要とする。

初期化の方法(初期データが必要かどうか、初期化データの型の違い)：
```
// 空のバッファ -> ポインタではなく実体なので、引数に渡すときは`&buffer1`のようにポインタ値を取り出す必要がある。
var buffer1 bytes.Buffer
// バイト列で初期化
buffer2 := bytes.NewBuffer([]{byte{0x10, 0x20, 0x30})
// 文字列で初期化
buffer3 := bytes.NewBufferString(" 初期文字列")
```

```
// bytes.Reader はbytes.NewReader で作成
bReader1 := bytes.NewReader([]byte{0x10, 0x20, 0x30})
bReader2 := bytes.NewReader([]byte(" 文字列をバイト配列にキャストして設定")
// strings.Reader はstrings.NewReader() 関数で作成
sReader := strings.NewReader("Reader の出力内容は文字列で渡す")
```

## 3.5 バイナリ解析用のio.Reader関連機能

io.Reader から出てくるデータは:
- テキストデータのこともあれば
- バイナリデータのこともあります。

まずはバイナリデータを読み込むときに便利な機能を見る

### 3.5.1 必要な部位を切り出すio.LimitReader／io.SectionReader

io.LimitReader を使うと、データがたくさん入っていても、先頭の一部だけしか読み込めないようにブロック

使用ケース:

- たとえばファイルの先頭にヘッダー領域があって、そこだけを解析するルーチンに処理を渡したい

```
// たくさんデータがあっても先頭の16 バイトしか読み込めないようにする
lReader := io.LimitReader(reader, 16)
```

io.SectionReader:

- io.Reader が使えず、代わりにio.ReaderAt を使う
― 使える型、使えない型について
  - os.File 型はio.ReaderAt を満たしますが、
  - それ以外のio.Reader を満たす型からio.SectionReader で直接に読み込むことはできません
    - strings.Reader やbytes.Reader でラップしてから渡す

使用ケース：

- PNGファイルやOpenType フォントなど、バイナリファイル内がいくつかのチャンク（データの塊）に分かれている場合は、
チャンクごとにReader を分けて読み込ませる

### 3.5.2 エンディアン変換

バイナリ解析ではエンディアン変換が必要。なぜなら、ネットワークで受け取ったデータをリトルエンディアンに修正する必要があるから。

- CPU はリトルエンディアン
  - 小さい桁からメモリに格納
- ネットワーク上で転送されるデータの多くは、ビッグエンディアン
  - 大きい桁からメモリに格納される

encoding/binary パッケージのbinary.Read()メソッドに、
io.Reader とデータのエンディアン、それに変換結果を格納する変数のポインタを渡せば、
エンディアンが修正されたデータが得られる

### 3.5.3 PNGファイルを分析してみる

PNGファイルはバイナリフォーマットです。
- 先頭の8バイトがシグニチャ（固定のバイト列）となっています。
- それ以降は図3.3 のようなチャンクのブロックで構成

```
| Signature | Chunck{Length | Type | Data | CRC} | Chunck{Length | Type | Data | CRC} | ...
```

例：各チャンクとその長さを列挙

```
chunk 'IHDR' (13 bytes)
chunk 'sRGB' (1 bytes)
chunk 'IDAT' (473761 bytes)
chunk 'IEND' (0 bytes)
```

### 3.5.4 PNG画像に秘密のテキストを入れてみる

PNGに埋め込まれる(画像としては表示されない)テキストのチャンク：

- テキストを追加するためのtEXt
- また、それに圧縮をかけたzTXt

chunkの書き込み、テキストチャンクの中は複雑に見えますが、パーツごとに見れば、
それぞれio.Writerの書き込みのみで構成されているから実際単純。
- binary.Write()による長さの書き込み、
- 次にチャンク名の書き込み、
- 本体の書き込み、
- 最後にCRCの計算と、binary.Write() による書き込みです

```
chunk 'IHDR' (13 bytes)
chunk 'tEXt' (19 bytes)
ASCII PROGRAMMING++
chunk 'sRGB' (1 bytes)
chunk 'IDAT' (473761 bytes)
chunk 'IEND' (0 bytes)
```

## 3.6 テキスト解析用のio.Reader関連機能

### 3.6.1 改行／単語で区切る

### 3.6.2 データ型を指定して解析

### 3.6.3 その他の形式の決まったフォーマットの文字列の解析

## 3.7 io.Reader／io.Writerでストリームを自由に操る

## 3.8 本章のまとめと次章予告

本章ではio.Readerの仲間たちと、io.Reader と一緒に使う補助関数、具体的なサンプルをいくつか紹介

書き込みと比べると読み込みのほうが複雑な機能が求められる。知っておくべきメソッドの数は読み込みに関するもののほうが多い:

- サイズやデータの種類を推定しながら読み込んだり、
- セクションごとに読み込み処理を切り替えたり、
- 可変長データを少ないメモリでうまく読み込む必要があったり、

だから、読み込みについては補助関数の機能や種類も豊富

次章はio.Writerやio.Readerとともに、低レベルAPI として使われるチャネルについて説明

## 3.9 問題
