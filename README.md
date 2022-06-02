# Goroutine

複数の条件で重いアルゴリズムを動かした結果を集計する場合を想定したテンプレートを作成する。

## Concurrency

[Concurrency is not parallelism](https://go.dev/blog/waza-talk) でも紹介されているように、プログラミングでは、並行処理とは独立して実行されるプロセスの`構成`であり、並列処理とは（おそらく関連する）計算を同時に`実行`することです。

## Difficulty

`goroutine` を利用した場合、コードの実行順が予測できない。そのため実行の順序に影響のあるような処理はには使うべきでない。

また、共有メモリに正しくアクセスしないといけず、ロックなどを用いて排他制御を行う必要があることに注意する。

サブの `goroutine` が終わる前にメインの `goroutine` が終わってしまうことがないよう気をつける。

## sync.WaitGroup

`goroutine` を複数使って並列で処理を行って、それがすべて完了したら次に進みたい場合は、全員が完了するまで待つ処理が必要になる

`sync.WaitGroup` は複数の `goroutine` の完了を待つためのカウンタのようなもので、`goroutine` を1つ実行させる前にカウンタを1つ増やし、処理が終了したらカウンタを1つ減らすという処理をする。全員が完了した後に実行したい処理の前に、カウンタが0になるまで待つ処理を書くことで実現できる。

> 参考資料: 
> [Rui Ueyama | sync.WaitGroupの正しい使い方](https://qiita.com/ruiu/items/dba58f7b03a9a2ffad65)

## sync.Mutex

チャネルやロックなどを使わずに複数の `goroutine` から共有変数を読み書きしていけない。安全であるである実装をするためには、読み書きする前にロックして、終わったらアンロックすることを忘れないようにする。

## Sample Code

```Go
func main() {
	var mutex = &sync.Mutex{}
	var s []int
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			mutex.Lock()
			s = append(s, n)
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Print(s)
}
```

