# 1. Go言語で覗くシステムプログラミングの世界

## 1.1 システムプログラミングとは

「OSの提供する機能を使ったプログラミング」をシステムプログラミングの定義として話を進める。

### 1.1.1 OSの機能について

次のようなものが「OSの機能」だといえるでしょう。
- メモリの管理
- プロセスの管理
- プロセス間通信
- ファイルシステム
- ネットワーク
- ユーザー管理（権限など）
- タイマー

本書の目的の下位のレイヤーを知る = これらのOSの機能を学ぶ

これらのOSの機能を、もっとプログラマーになじみやすい視点で、普段の開発にもフィードバックできるように見ていく

## 1.2 Go言語

Go言語の特徴：
- 多くのOSの機能を直接扱えて、
- なおかつ少ない行数で動くアプリケーションが作れる
- ライブラリの収集が簡単
- コンパイルで多くのエラーが見つかる
- 実行時のエラーもわかりやすい
- ガベージコレクションのおかげでメモリ管理を注意深く設計する必要がありません
- コンパイルも速い

Go言語は、スクリプト系言語よりははるかに高速であるとはいえ、C/C++ と比べると性能は落ちますし、バイナリサイズもかなり大きくなるから、すぐにC/C++と置き換わるわけではない。

OSを書くような言語にはならないものの、OSの機能を気軽に使え、高度に抽象化されすぎない簡単に書けるC言語という前提で、Go言語を取り扱っていく

## 1.3 Go言語のインストールと準備

### 1.3.1 Visual Studio CodeでGoが使えるようにする

### 1.3.2 はじめてのGoプロジェクト

## 1.4 デバッガーを使ってシステムコールを「見る」

先ほど書いたGo言語の"Hello World!"プログラムの、さらに下のレイヤーでは、いったい何が起きているのでしょうか。

Go言語であれば、処理系に全環境用のソースコードもすべてバンドルされています。そのため、デバッガーで処理を追いかけていくだけで、
OSの機能を直接呼び出すコードまで簡単に見ることができます

### 1.4.1 Visual Studio Codeでデバッガーを有効にする

### 1.4.2 デバッガーを使って"Hello World!"の裏側を覗く

デバッガーにバグを修正する機能はありません。
デバッガーに備わっているのは、プログラムが実行されるようすをプログラマーが観察する機能だけ

1. ブレークポイントの設定
2. Debug
3. ブレークポイントとして選択した行に処理がくると、そこでプログラムの実行がいったん中断

デバッグでやること：
- 継続実行（Continue）：次のブレークポイントに到達するまで処理を継続させる
- ステップオーバー（Step Over）：見えているソースコードの次の行に移動する。カーソル位置の関数の中は実行され、終了するところまで処理が進む
- ステップイン（Step Into）：関数呼び出しの中に飛び込む。下のレイヤーに降りていくときに使う
- ステップアウト（Step Out）：いま実行している関数が終了するところまで処理を進める
- 再スタート（Restart）：一度終了して再度実行を開始する
- 停止（Stop）：一度終了する

Unix:
このsyscall.Write() がシステムコールと呼ばれる関数です。
システムコールにはいくつも種類があり、アプリケーションのプログラム単体では達成できない仕事をOSのカーネルに依頼するために使います。
ここでは、プログラムの外のターミナルに対して文字列を出力するという仕事を依頼しています

Unix以外のWindows:
コンソールの場合はfd.writeConsole()メソッドを呼び出しています。最終的にはsyscall.Write() を呼ぶ。

Unix:
1. main()
2. fmt: Println() -> Fprintln()
3. os/file: File.Write() -> 
4. os/file : os/file_unix.go: File.write()
5. internal/poll/fd_unix.go FD.Write()
6. syscall: Write()

Windows:
1. main()
2. fmt: Println() -> Fprintln()
3. os/file: File.Write() -> 
4. os/file : os/file_windows.go: File.write()
5. internal/poll/fd_windows.go FD.Write() -> syscall: Write()
5'. internal/poll/fd_windows.go FD.Write() -> FD.wirteConsole() -> syscall: WriteConsole()


## 1.5 本章のまとめと次章予告

Go言語の"Hello World!" プログラムの実行時に下のレイヤーでどんなシステムコールが呼び出されているのか、デバッガーを使ってレイヤーを降りていくことで実際に確かめる

Go言語の基本的な機能だけで、下のレイヤーで起こることを意外なほどあっさりと確かめられる

以下の機能を知ることで、ファイルやインターネット通信のためのソケットなど、OSから提供されている機能をGoから活用しやすくなります。:
- 第2 章と第3 章では、システムコールよりも少しだけ高いレイヤーに戻り、io.Writer、io.Readerを例にGo言語の「インタフェース」という抽象化の仕組みについて説明し
- 第4章「低レベルアクセスへの入口3：チャネル」では、やはりGo言語における抽象化の仕組みのひとつであるチャネルについて説明

## 1.6 問題
