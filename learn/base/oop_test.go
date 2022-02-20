package base

import (
	"fmt"
	"testing"
)


type P struct {
	int
}

func (t P) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t *P) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func TestDemo1(t *testing.T) {
	t1 := P{1}
	t1.testT()
	t1.testP()
	t2 := &t1
	fmt.Printf("t2 is : %v\n", t2)
	t2.testT()
	t2.testP()
}


type Mover interface {
	move()
}

type dog struct {}

func (d *dog) move() {
	fmt.Println("狗会动")
}

func TestDemo2(t *testing.T) {
	var x Mover
	// var wangcai = dog{} // 旺财是dog类型
	// x = wangcai         // x不可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}

