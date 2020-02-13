package main

import (
	"errors"
	"fmt"
)

type iter struct {
	arr     []string
	current int
}

func newIter(arr []string) *iter {
	return &iter{arr: arr, current: 0}
}

func (i *iter) next() (string, error) {
	if !i.hasNext() {
		return "", errors.New("iteration has no more elements")
	}
	r := i.arr[i.current]
	i.current++
	return r, nil
}

func (i *iter) hasNext() bool {
	return i.current < len(i.arr)
}

func main() {

	fmt.Println("=========regular array traversing==========")

	sa := [4]int{}

	sa[0] = 1
	sa[1] = 1
	sa[2] = 1
	sa[3] = 1

	fmt.Printf("%v %T \n", sa, sa)

	for i := range sa {
		fmt.Println(i, sa[i])
	}

	for i, el := range sa {
		fmt.Println(i, el)
	}

	var asa [5]int

	asa[0] = 1
	asa[4] = 3

	fmt.Printf("%v %T \n", asa, asa)

	fmt.Println("=========regular slice traversing==========")

	ss := make([]int, 4)

	fmt.Printf("%v %T \n", ss, ss)

	ss[0] = 1
	ss[1] = 2
	ss[2] = 3
	ss[3] = 4

	fmt.Printf("%v %T len: %d\n", ss, ss, len(ss))

	ss = append(ss, 5)

	fmt.Printf("%v %T len: %d\n", ss, ss, len(ss))

	for i := range ss {
		fmt.Println(i, ss[i])
	}

	ss2 := []int{}

	ss2 = append(ss2, 1)

	fmt.Printf("%v %T \n", ss2, ss2)

	for i := range ss2 {
		fmt.Println(i, ss2[i])
	}

	fmt.Println("=========regular map traversing==========")

	sm := map[string]string{}

	fmt.Printf("%v %T \n", sm, sm)

	sm["one"] = "foo"
	sm["two"] = "bar"
	sm["three"] = "baz"

	fmt.Printf("%v %T \n", sm, sm)

	for k := range sm {
		fmt.Println(k, sm[k])
	}

	for k, v := range sm {
		fmt.Println(k, v)
	}

	for _, v := range sm {
		fmt.Println(v)
	}

	fmt.Println("=========custom iterator traversing==========")

	i := newIter([]string{"foo", "bar", "baz"})

	for i.hasNext() {
		c, _ := i.next()
		fmt.Println(c)
	}

	_, err := i.next()

	fmt.Println(err.Error())

}
