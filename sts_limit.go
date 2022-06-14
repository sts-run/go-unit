package go_unit

import (
	"fmt"
	"time"
)

//这里实现了一个简单的令牌桶限流

type limitFlow struct {
	limitNum  int           //桶的容量
	speed     time.Duration //令牌更新速率
	container chan byte     //桶容器
}

func NewLimitFlow(limitNum int, speed time.Duration) *limitFlow {
	l := &limitFlow{
		limitNum: limitNum,
		speed:    speed,
	}

	ch := make(chan byte, l.limitNum)
	l.container = ch
	l.upToken()

	return l
}

//更新令牌常驻任务
func (l *limitFlow) upToken() {
	go func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("update token error：%v", err))
		}

		for {
			diff := l.limitNum - len(l.container)
			l.push(diff)

			time.Sleep(l.speed)
		}
	}()
}

//获取令牌
func (l *limitFlow) GetToken() bool {
	if len(l.container) > 0 {
		<-l.container
		return true
	}
	return false
}

//添加令牌
func (l *limitFlow) push(diff int) {
	for {
		if diff <= 0 {
			return
		}
		l.container <- 1
		diff--
	}
}

//获取当前令牌数
func (l *limitFlow) GetCurrTokenNum() int {
	return len(l.container)
}
