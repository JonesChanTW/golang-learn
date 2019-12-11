package asynchronous

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const routingNum int = 5 ///畢竟只是要展示cancle的用法,所以gorouting 別放太多,不然cancle的時候有機會造成主程式跑完了 其他的卻還正在結束
type WorkInfo struct {
	workType int
	content  string
}

func tesk(ctx context.Context, searial int, wg *sync.WaitGroup) {
	var x int64 = 0
	for {
		x++ ///do anything...
		select {
		case <-ctx.Done():
			t := time.Now()
			timeStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
			fmt.Printf("%s task %d has been cancled\n", timeStr, searial)
			wg.Done()
			return
		default:
		}
	}
}

// ContextWithCancelTest cancle測試
func ContextWithCancelTest() {
	var wg sync.WaitGroup
	///先基本的要一個可以取消的context
	ctx, cancleFunc := context.WithCancel(context.Background())

	for i := 0; i < routingNum; i++ {
		go tesk(ctx, i, &wg)
	}

	time.Sleep(time.Second)
	cancleFunc()
	wg.Wait()
	fmt.Println("Finish...")
}

// ContextTimeoutTest timeout 測試,設定等候多久後逾時
func ContextTimeoutTest() {
	var wg sync.WaitGroup
	ctx, cancleFunc := context.WithTimeout(context.Background(), 5*time.Second) ///5秒逾時

	defer cancleFunc()
	for i := 0; i < routingNum; i++ {
		tesk(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("ContextTimeoutTest Finish...")
}

// ContextDeadlineTest 做到甚麼時候為止,給定一個時間,例如幾點幾分...,而不是多長的時間
func ContextDeadlineTest() {
	var wg sync.WaitGroup
	dl := time.Now().Add(30 * time.Second)
	ctx, cancleFunc := context.WithDeadline(context.Background(), dl)

	defer cancleFunc()
	t := dl
	timeStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("goroutings will be stop at ", timeStr)
	for i := 0; i < routingNum; i++ {
		go tesk(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("ContextDeadlineTest Finish...")
}

// ContextValueTest Value
func ContextValueTest() {
	type ctxKey string
	key := ctxKey("firstWork")
	info := WorkInfo{workType: 10, content: "This is Work info"}

	cancleCtx, cancleFunc := context.WithCancel(context.Background())
	ctx := context.WithValue(cancleCtx, key, info) ///WithValue 的回傳中不包含 cancle func
	var wg sync.WaitGroup

	for i := 0; i < routingNum; i++ {
		wg.Add(1)
		go func(ctx context.Context, searial int, wg *sync.WaitGroup) {
			var x int64 = 0
			for {
				x++ ///do anything...
				select {
				case <-ctx.Done():
					t := time.Now()
					timeStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
					fmt.Printf("%s task %d has been cancled\n", timeStr, searial)
					defer wg.Done()
					return
				default:
					fmt.Println("task ", searial, " get val = ", ctx.Value(key))
					time.Sleep(time.Second)
				}
			}
		}(ctx, i, &wg)
	}

	fmt.Println("Wait finish")
	time.Sleep(6 * time.Second) ///讓他跑六秒看輸出
	cancleFunc()
	wg.Wait()
	fmt.Println("Finish")
}
