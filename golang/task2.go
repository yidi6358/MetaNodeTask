/*指针题目1开始*/
func addTen(thisValue *int) int {
	var result int
	result = *thisValue + 10
	return result
}
/*指针题目1结束*/

/*指针题目2开始*/
func multiTwo(ptr *[]int) []int {
	slice := *ptr
	for i := range *ptr {
		slice[i] = slice[i] * 2
	}
	return *ptr
}
/*指针题目2结束*/


/*Goroutine题目1 开始*/
func printQi() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
func printWo() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
/*Goroutine题目1 结束*/


/*Goroutine题目2 开始*/
func main() {
	fmt.Println("123")
	// 模拟一些示例任务
	tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Task A done")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task B done")
		},
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Task C done")
		},
	}

	Scheduler(tasks)

}

type Task func()

// 调度器：接收任务列表，并并发执行，统计耗时
func Scheduler(tasks []Task) {
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, t Task) {
			defer wg.Done()

			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)

			fmt.Printf("Task #%d finished in %v\n", i+1, duration)
		}(i, task)
	}

	wg.Wait() // 等待所有任务完成
}
/*Goroutine题目2 结束*/


/*面向对象题目1 开始*/
func (rec Rectangle) Area() float64 {
	return rec.chang * rec.kuan
}
func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.chang + rec.kuan)
}

func (c Circle) Area() float64 {
	return 3.14 * (c.r * c.r)
}

func (c Circle) Perimeter() float64 {
	return 2 * (3.14 * c.r)
}

func main() {
	cectangle := Rectangle{chang: 10, kuan: 20}
	circle := &Circle{r: 10}

	fmt.Println(cectangle.Area())
	fmt.Println(cectangle.Perimeter())
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())
}
/*面向对象题目1 结束*/


/*面向对象题目2 开始*/
type Person struct {
	Name string
	age  int
}

type Employee struct {
	person    Person
	EmplyeeId int
}

func (e Employee) PrintInfo() {
	fmt.Println("此员工的名字=", e.person.Name, " 年龄=", e.person.age, " 员工编号=", e.EmplyeeId)
}

func main() {
	person := Person{"sam", 20}

	employee := Employee{person, 10001}
	employee.PrintInfo()
}
/*面向对象题目2 结束*/


/*Channel 题目1 开始*/
func sendNumber(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("发送:", i)
	}
	close(ch)
}

func receivNumber(ch <-chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	defer wg.Done()
	for v := range ch {
		time.Sleep(time.Second * 1)
		fmt.Println("接收: ", v)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1) // 消费者
	fmt.Println(125)
	go sendNumber(ch)
	go receivNumber(ch, &wg)
	wg.Wait()
}
/*Channel 题目1 结束*/


/*Channel 题目2 开始*/
func sendNumber(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("发送:", i)
	}
	close(ch)
}

func receivNumber(ch <-chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	defer wg.Done()
	for v := range ch {
		time.Sleep(time.Second * 1)
		fmt.Println("接收: ", v)
	}
}

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1) // 消费者
	fmt.Println(125)
	go sendNumber(ch)
	go receivNumber(ch, &wg)
	wg.Wait()
}
/*Channel 题目2 结束*/


/*锁机制 题目1 开始*/
func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var counter int

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter的最终值：", counter)
}
/*锁机制 题目1 结束*/


/*锁机制 题目2 开始*/
func main() {
	var wg sync.WaitGroup
	var counter int32

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter的最终值：", counter)
}
/*锁机制 题目2 结束*/
