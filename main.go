package main

import "fmt"

func main() {
	TestQueue()
}

func TestQueue() {
	queue := NewQueue() // var queue Queue shunday elon qilsak ham boladi
	queue.AddQueue(101)
	fmt.Println("queue:", queue)
	queue.AddQueue(24).AddQueue(77)
	fmt.Println(queue)
	fmt.Println("queue boshmi? ", queue.IsEmpty())

	var result, _ = queue.Peak()
	fmt.Println("Navbatda: ", result)
	fmt.Println("Ushbu element o'chirildi")
	result, _ = queue.DeleteQueue()
	fmt.Print(result)
	fmt.Println("queue boshmi? ", queue.IsEmpty())
	fmt.Println("Navbatda turganlar soni: ", queue.LengthQueue())

}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Peak() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Qator bosh")
	}
	return q.data[0], nil
}

func (q *Queue) AddQueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

func (q *Queue) DeleteQueue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Qator bosh")
	}
	element := q.data[0]
	q.data = q.data[1:] //birinchi elementdan boshlab barcha elementlarni qaytaradi(o'zlashtiradi)
	return element, nil
}

func (q *Queue) LengthQueue() int {
	l := len(q.data)
	return l
}
