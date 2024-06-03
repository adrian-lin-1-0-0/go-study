package peterson

import (
	"sync"
	"testing"
)

/*
p0:
flag[0] = true //1
turn = 1 //2
for flag[1] && turn == 1 {} //3
// critical section
fmt.Println("p0") //4

flag[0] = false //5

p1:
flag[1] = true //6
turn = 0 //7
for flag[0] && turn == 0 {} //8
// critical section
fmt.Println("p1") //9
flag[1] = false //10
*/

var flag []bool
var turn int
var criticalSection []int

func p0_1() {
	flag[0] = true
}
func p0_2() {
	turn = 1
}
func p0_3() {
	for flag[1] && turn == 1 {
		//busy wait
	}
}
func p0_4() {
	criticalSection = append(criticalSection, 0)
}
func p0_5() {
	flag[0] = false
}

func p1_6() {
	flag[1] = true
}
func p1_7() {
	turn = 0
}
func p1_8() {
	for flag[0] && turn == 0 {
		//busy wait
	}
}
func p1_9() {
	criticalSection = append(criticalSection, 1)
}
func p1_10() {
	flag[1] = false
}

func Test_123678(t *testing.T) {
	flag = make([]bool, 2)
	turn = -1
	criticalSection = []int{}
	wg := sync.WaitGroup{}

	p0_1()
	p0_2()
	wg.Add(1)
	go func() {
		wgIn := sync.WaitGroup{}
		wgIn.Add(1)
		go func() {
			p0_3()
			wgIn.Done()
		}()
		wgIn.Wait()
		p0_4()
		p0_5()
		wg.Done()
	}()
	p1_6()
	p1_7()
	wg.Add(1)
	go func() {
		wgIn := sync.WaitGroup{}
		wgIn.Add(1)
		go func() {
			p1_8()
			wgIn.Done()
		}()
		wgIn.Wait()
		p1_9()
		p1_10()
		wg.Done()
	}()

	wg.Wait()
	if criticalSection[0] != 0 {
		t.Errorf("critical section 0 is not executed")
	}
	if criticalSection[1] != 1 {
		t.Errorf("critical section 1 is not executed")
	}
}

func Test_1623(t *testing.T) {
	flag = make([]bool, 2)
	turn = -1
	criticalSection = []int{}
	wg := sync.WaitGroup{}

	p0_1()
	p1_6()
	p0_2()
	wg.Add(1)
	go func() {
		wgIn := sync.WaitGroup{}
		wgIn.Add(1)
		go func() {
			p0_3()
			wgIn.Done()
		}()
		wgIn.Wait()
		p0_4()
		p0_5()
		wg.Done()
	}()

	p1_7()
	wg.Add(1)
	go func() {
		wgIn := sync.WaitGroup{}
		wgIn.Add(1)
		go func() {
			p1_8()
			wgIn.Done()
		}()
		wgIn.Wait()
		p1_9()
		p1_10()
		wg.Done()
	}()

	wg.Wait()
	if criticalSection[0] != 0 {
		t.Errorf("critical section 0 is not executed")
	}
	if criticalSection[1] != 1 {
		t.Errorf("critical section 1 is not executed")
	}
}

func Test_16278(t *testing.T) {
	flag = make([]bool, 2)
	turn = -1
	criticalSection = []int{}
	wg := sync.WaitGroup{}

	p0_1()
	p1_6()
	p0_2()
	p1_7()
	wg.Add(1)
	go func() {
		wgIn := sync.WaitGroup{}
		wgIn.Add(1)
		go func() {
			p1_8()
			wgIn.Done()
		}()
		wgIn.Wait()
		p1_9()
		p1_10()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		wgIn2 := sync.WaitGroup{}
		wgIn2.Add(1)
		go func() {
			p0_3()
			wgIn2.Done()
		}()
		wgIn2.Wait()
		p0_4()
		p0_5()
		wg.Done()
	}()
	wg.Wait()

	if criticalSection[0] != 0 {
		t.Errorf("critical section 0 is not executed")
	}
	if criticalSection[1] != 1 {
		t.Errorf("critical section 1 is not executed")
	}
}
