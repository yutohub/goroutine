package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 複数の goroutine から読み書きするためにロック機構を入れる
	var mutex = &sync.Mutex{}
	// 結果を書き込むためのスライスを用意する
	var s []int
	// 全ての goroutine が終わるまで待つための機構を入れる
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		// goroutine のカウンタを一つ増やす
		wg.Add(1)
		// goroutine を一つ増やす
		go func(n int) {
			// 時間のかかる処理を行う
			time.Sleep(5 * time.Microsecond)
			// 書き込みする前にロックする
			mutex.Lock()
			// 結果の書き込みをする
			s = append(s, n*n)
			// 書き込み終わったのでアンロックする
			mutex.Unlock()
			// goroutine のカウンタを一つ減らす
			wg.Done()
		}(i)
	}

	// 全ての goroutine が終わるのを待つ
	wg.Wait()
	fmt.Print(s)
}
