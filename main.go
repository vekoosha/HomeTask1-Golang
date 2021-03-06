package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("stdin.txt") // just pass the file name
	check(err)
	//fmt.Println(b)   // print the content as 'bytes'
	str := string(b) // convert content to a 'string'

	fmt.Println("Содержимое файла stdin.txt:")
	fmt.Printf(str)

	var sl []string
	var slInt []int
	sl = strings.Split(str, " ")
	for i := 0; i < len(sl); i++ {
		n, err := strconv.Atoi(sl[i])
		if err != nil {
			fmt.Print(err)
		} else {
			slInt = append(slInt, n)
			//fmt.Println(sl[i])
		}
	}
	insertsort(slInt)

	file, err := os.Create("stdout.txt")
	check(err)
	defer file.Close()

	writeFile(slInt, file, err)
}

func check(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func writeFile(p []int, file *os.File, err error) {

	/*	var file, err = os.OpenFile("stdout.txt", os.O_RDWR, 0644)
		check(err)
		defer file.Close()*/

	//запись
	for i := 0; i < len(p); i++ {
		_, err = file.WriteString(fmt.Sprint(p[i], " "))
		check(err)
	}
	check(err)

	//сохраяем изменения
	err = file.Sync()
	check(err)

	fmt.Println("\nСодержимое файла stdout.txt:")
	cont, err := ioutil.ReadFile("stdout.txt")
	check(err)
	fmt.Printf("%s", cont)
}

func insertsort(s []int) {
	k := len(s)
	j := 0
	for i := 1; i < k; i++ {
		key := s[i] //первый позиционный элемент в данном обходе
		j = i - 1   // запоминаем индекс предыдущего элемента массива
		for j > -1 && s[j] > key {
			s[j+1] = s[j] // перестановка элементов массива
			s[j] = key
			j--
		}
	}
	//fmt.Println(s)
}
