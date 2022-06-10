package go_unit

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"runtime"
)

//封装的  可以截获恐慌的goroutine  避免子协程恐慌导致整个服务退出
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				stackInfo := fmt.Sprintf("%s", buf[:n])
				log.Println("get goroutinue err log：" + fmt.Errorf("%v", err).Error() + ",\"stackInfo\":" + stackInfo)
			}
		}()

		f()
	}()
}

//返回2个数之间的随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		return min
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
