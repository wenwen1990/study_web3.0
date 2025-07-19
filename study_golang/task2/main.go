package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var num1 = 1
	fmt.Println(question1(&num1))
	var arr2 = []int{1, 2, 3}
	question2(&arr2)
	fmt.Println(arr2)
	question3()
	question4()
	question5()
	question6()
	question7()
	question8()
	question9()
	question10()
}

/*
1、题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func question1(idx *int) int {
	*idx += 10
	return *idx
}

/*
2、题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func question2(slice *[]int) {
	for idx := range *slice {
		(*slice)[idx] *= 2
	}
}

/*
3、题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func question3() {
	fmt.Println("启动两个协程")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数：", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数：", i)
		}
	}()
	wg.Wait()
	fmt.Println("结束两个协程")
}

/*
4、题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func question4() {
	fmt.Println("初始化调度器")
	scheduler := &Scheduler{}
	task1 := func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务1执行完毕！")
	}
	task2 := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务2执行完毕！")
	}
	task3 := func() {
		time.Sleep(3 * time.Second)
		fmt.Println("任务3执行完毕！")
	}
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.Run()
	fmt.Println("定义3个任务，并发送给调度器执行,打印每个任务执行时间，执行完毕")

}

// 任务类型：Task 定义了一个无返回值的函数类型。
type Task func()

// 调度器：Scheduler 是任务调度器，它有一个任务列表 tasks，并且包含一个 sync.WaitGroup 用于等待所有任务完成。
type Scheduler struct {
	tasks []Task
	wg    sync.WaitGroup
}

// 添加任务：AddTask 方法用于将任务添加到调度器中。
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// 运行任务：方法遍历任务列表，为每个任务启动一个新的 goroutine，并且记录每个任务的执行时间。
func (s *Scheduler) Run() {
	for _, task := range s.tasks {
		s.wg.Add(1)
		go func(task Task) {
			defer s.wg.Done()
			startTime := time.Now()
			task()
			costTime := time.Since(startTime)
			fmt.Printf("c花费时间: %v秒\n", costTime.Seconds())

		}(task)
	}
	s.wg.Wait()
}

/*
5、题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
func question5() {
	var r Rectangle = Rectangle{width: 3, height: 5}
	fmt.Println("矩形面积方法:", r.Area())
	fmt.Println("矩形周长方法:", r.Perimeter())
	var c Circle = Circle{radius: 2}
	fmt.Println("圆形面积方法:", c.Area())
	fmt.Println("圆形周长方法:", c.Perimeter())
}

// 定义一个 Shape 接口
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 //计算周长
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Area() float64 {
	var res = r.height * r.width
	return res
}

func (r Rectangle) Perimeter() float64 {
	var res = (r.height + r.width) * 2
	return res
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	var res = math.Pi * math.Pow(c.radius, 2)
	return res
}

func (c Circle) Perimeter() float64 {
	var res = 2 * math.Pi * c.radius
	return res

}

/*
6、题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
func question6() {
	var employee = Employee{
		EmployeeID: 101,
		Person: Person{
			Name: "张三",
			Age:  18,
		},
	}
	employee.PrintInfo()
}

type Person struct {
	Name string
	Age  uint
}
type Employee struct {
	EmployeeID uint
	Person     Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("打印员工信息：员工id：%v,姓名：%v,年龄：%v", e.EmployeeID, e.Person.Name, e.Person.Age)
}

/*
7、题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func question7() {
	fmt.Println("初始化无缓冲的通道")
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("发送：", i)
		}
		close(ch)
		fmt.Println("关闭通道！")
	}()
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("接收：", value)
		}
	}()
	wg.Wait()
	fmt.Println("结束！")
}

/*
8、题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func question8() {
	fmt.Println("初始化带缓冲的通道")
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)
	go func(ch chan<- int) {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("发送数据：", i)
			time.Sleep(50 * time.Millisecond)
		}
		close(ch)
		fmt.Println("关闭通道！")
	}(ch)
	go func(ch <-chan int) {
		defer wg.Done()
		for value := range ch {
			fmt.Println("接收数据：", value)
		}
	}(ch)
	wg.Wait()
	fmt.Println("结束！")
}

/*
9、题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func question9() {
	var wg sync.WaitGroup
	counter := Counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Add()
			}
		}()
	}
	wg.Wait()
	fmt.Println("有锁计数器最后值：", counter.GetNum())
}

type Counter struct {
	lock sync.Mutex
	num  uint
}

// 写成Counter有误
func (c *Counter) Add() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.num++
}
func (c *Counter) GetNum() uint {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.num
}

/*
10、题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func question10() {
	var num uint64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddUint64(&num, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("无锁计数器最后值：", num)
}
