# go-unsafe-sample

[GCに思いを馳せてunsafeなコードを書く \- Qiita](https://qiita.com/taxio/items/b49629d3f44f1c933f6d)の写経


最初の実行結果。
そういえば`runtime.GC()`って初めて使ったな。これでGC呼べるのか。

```zsh
❯ go run main.go
before: &{123 abc}
after: &{123 abc}
```

`-m`フラグでビルド。

> Use -gcflags -m to observe the result of escape analysis and inlining decisions for the gc toolchain.

`-m`フラグをつけるとescape解析と、gcツールのインライン決定の結果が見れるらしい。

```sh
❯ go build -gcflags="-m" main.go
# command-line-arguments
./main.go:19:12: inlining call to fmt.Printf
./main.go:23:6: can inline newXAsIntPtr
./main.go:31:21: inlining call to newXAsIntPtr
./main.go:35:12: inlining call to fmt.Printf
./main.go:17:6: &X literal escapes to heap
./main.go:19:13: x escapes to heap
./main.go:19:12: NewX []interface {} literal does not escape
./main.go:19:12: io.Writer(os.Stdout) escapes to heap
./main.go:35:13: x escapes to heap
./main.go:35:12: main []interface {} literal does not escape
./main.go:35:12: io.Writer(os.Stdout) escapes to heap
<autogenerated>:1: (*File).close .this does not escape
```

`-l`でインライン化をdisableできるので、追加

```sh
❯ go build -gcflags="-m -l" main.go
# command-line-arguments
./main.go:17:6: &X literal escapes to heap
./main.go:19:12: NewX ... argument does not escape
./main.go:19:13: x escapes to heap
./main.go:35:12: main ... argument does not escape
./main.go:35:13: x escapes to heap
```

## References
- [GCに思いを馳せてunsafeなコードを書く \- Qiita](https://qiita.com/taxio/items/b49629d3f44f1c933f6d)
- [CompilerOptimizations · golang/go Wiki](https://github.com/golang/go/wiki/CompilerOptimizations#escape-analysis)
- [go build \-gcflags '\-N \-l' とは何か \- The King's Museum](https://www.thekingsmuseum.info/entry/2017/09/21/113002)