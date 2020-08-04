package makeCpu

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"othello/cpu"
	"othello/simulate"
)

var trynum = 0
var chnum = 0

func climb(now []int, ch <-chan bool, rec chan<- []int) {
	for {
		select {
		case <-ch:
			rec <- now
			return
		default:
			trynum++
			newCpu := now
			//for i := 0; i < cpu.Pow3[10]; i++ {
			pos := rand.Intn(cpu.Pow3[16])
			val := rand.Intn(1e15)
			pm := rand.Intn(2)
			if pm == 1 {
				val *= -1
			}

			newCpu[pos] = val
			//}

			res := simulate.MakeCpuSimulate(now, newCpu)
			if res > 0 {
				chnum++
				now = newCpu
			}
		}
	}
}

func MakeCpu() {
	ch := make(chan bool)
	receiver := make(chan []int)
	rand.Seed(time.Now().UnixNano())
	f, err := os.Open("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	for i := 0; i < 16; i++ {
		now := make([]int, cpu.Pow3[16])
		for i := 0; i < cpu.Pow3[16]; i++ {
			scanner.Scan()
			now[i], err = strconv.Atoi(scanner.Text())
			if err != nil {
				now[i] = rand.Intn(1e15)
				if rand.Intn(2) == 1 {
					now[i] *= -1
				}
			}
		}
		fmt.Printf("%dth read finished\n", i+1)
		go climb(now, ch, receiver)
	}

	err = os.Remove("result.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(20 * time.Second)
		fmt.Printf("%d%% finished  current trynum is %d chnum, chnum is %d\n", (i+1)*10, trynum, chnum)
	}

	for i := 0; i < 16; i++ {
		ch <- true
	}

	time.Sleep(10 * time.Second)
	fmt.Println("goroutines closed")

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	//flag := false
	//mx := make([]int, cpu.Pow3[16])
	for i := 0; i < 16; i++ {
		now := <-receiver
		//if flag {
		//	cnt := 0
		//	for j := 0; j < 5; j++ {
		//		res := simulate.MakeCpuSimulate(mx, now)
		//		if res > 0 {
		//			cnt++
		//		}
		//	}
		//	if cnt >= 3 {
		//		mx = now
		//	}
		//} else {
		//	mx = now
		//	flag = true
		//}
		for i := 0; i < cpu.Pow3[16]; i++ {
			str := strconv.Itoa(now[i])
			_, err := writer.WriteString(str)
			if err != nil {
				log.Fatal(err)
			}
			_, err2 := writer.WriteString("\n")
			if err2 != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("%dth write finished\n", i+1)
	}
}
