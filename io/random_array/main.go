package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

func main() {
	// arr := generateSample(10000000, 500)
	// if err := save("big_array", arr); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(arr)
	fmt.Println(read("big_array"))
}

// generates a random array of size n
func generateSample(size int, seed int64) []int64 {
	rand.Seed(seed)
	arr := make([]int64, size)
	for i := 0; i < size; i++ {
		arr[i] = int64(rand.Intn(10 * size))
	}
	return arr
}

// saves an array in the fs
// the first 8 bytes position contains the size of the array
func save(filename string, arr []int64) error {
	arr = append([]int64{int64(len(arr))}, arr...)
	buf := new(bytes.Buffer)
	for _, v := range arr {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0664)
}

// reads an array from the fs
// expects the first 8 bytes position to contain the size of the array
func read(filename string) ([]int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	br := bufio.NewReader(f)

	lb := make([]byte, 8)
	if _, err = br.Read(lb); err != nil {
		return nil, err
	}
	l := int64(binary.LittleEndian.Uint64(lb))
	arr := make([]int64, l)
	if err := binary.Read(br, binary.LittleEndian, &arr); err != nil {
		return nil, err
	}
	return arr, nil
}
