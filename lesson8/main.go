// Создайте массив из 7 целых чисел и инициализируйте его значениями от 1 до 7.
// package main

// import "fmt"

// func main() {
//     // Создание и инициализация массива из 7 целых чисел
//     var arr [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}

//     // Вывод массива на консоль
//     fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 5 строк и инициализируйте его значениями "a", "b", "c", "d", "e".
// package main

// import "fmt"

// func main() {
//     // Создание массива из 5 строк
//     arr := [5]string{"a", "b", "c", "d", "e"}

//	    // Вывод массива на экран
//	    fmt.Println(arr)
//	}

//-----------------------------------------------------------------------------------

// Создайте массив из 4 логических значений и инициализируйте его значениями true, false, true, false.
// package main

// import "fmt"

//	func main() {
//		arr := [4]bool{true, false, true, false}
//		fmt.Println(arr)
//	}

//-----------------------------------------------------------------------------------

// Создайте массив из 10 целых чисел без инициализации и выведите его на экран.
// package main

// import "fmt"

// func main() {
//     // Создание массива из 10 целых чисел без явной инициализации
//     var arr [10]int

//     // Вывод массива на экран
//     fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 5 вещественных чисел и инициализируйте его значениями 1.1, 2.2, 3.3, 4.4, 5.5.
// package main

// import "fmt"

//	func main() {
//		arr:=[5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
//		fmt.Println(arr)
//	}

//-----------------------------------------------------------------------------------

// Создайте массив из 3 строковых значений с автоматическим определением длины массива.

// package main

// import "fmt"

// func main() {
// 	arr := [...]string{"a", "b", "c"}
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Выведите на экран первый и последний элемент массива [5]int{10, 20, 30, 40, 50}.

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{10, 20, 30, 40, 50}
// 	fmt.Println(arr[0], arr[cap(arr)-1])
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 логических значений и выведите значения по индексам 1 и 3.
// package main

// import "fmt"

// func main() {
// 	arr := [...]bool{true, true, false, false}
// 	fmt.Println(arr[1], arr[3])
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 6 вещественных чисел и выведите значение элемента по индексу 2.
// package main

// import "fmt"

// func main() {
// 	arr := [...]float64{1.6, 2.5, 3.4, 4.3, 5.2, 6.1}
// 	fmt.Println(arr[2])
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 5 целых чисел и измените значение второго элемента на 100.
// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	arr[1] = 10
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 строк и измените значение последнего элемента на "Go".
// package main

// import "fmt"

// func main() {
// 	arr := [4]string{"Assembler", "C", "C++", "Python"}
// 	arr[cap(arr)-1] = "Go"
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 3 логических значений и измените значение первого элемента на false.
// package main

// import "fmt"

// func main() {
//     arr := [3]bool{true, true, true}

//     arr[0] = false

//     fmt.Println(arr)
//  }

//-----------------------------------------------------------------------------------

// Создайте массив из 7 целых чисел и выведите его длину и вместимость.
// package main

// import "fmt"

// func main() {
// 	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
// 	fmt.Println(len(arr), cap(arr))
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 3 строк и выведите его длину и вместимость.
// package main

// import "fmt"

// func main() {
// 	arr := [3]string{"a", "b", "c"}
// 	fmt.Println(len(arr), cap(arr))
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 5 логических значений и выведите его длину и вместимость.
// package main

// import "fmt"

// func main() {
// 	arr := [5]bool{true, true, true, false, true}
// 	fmt.Println(len(arr), cap(arr))
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 5 целых чисел и выведите все его элементы с помощью цикла for.
// package main

// import "fmt"

// func main() {
// 	arr := [5]int{10, 20, 30, 40, 50}
// 	for i := 0; i < cap(arr); i++ {
// 		fmt.Print(arr[i]," ")
// 	}
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 строк и выведите все его элементы с помощью цикла for range.
// package main

// import "fmt"

// func main() {
// 	arr := [4]string{"a", "b", "c", "d"}
// 	for index, value := range arr {
// 		fmt.Println(index, value)
// 	}
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 3 логических значений и выведите все его элементы с помощью цикла for.
// package main

// import "fmt"

// func main() {
// 	arr := [3]bool{true, false, true}
// 	for i := 0; i < cap(arr); i++ {
// 		fmt.Println(arr[i])
// 	}
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 5 целых чисел и инициализируйте его значениями от 1 до 5.
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте пустой срез из строк и добавьте в него значения "a", "b", "c".
// package main

// import "fmt"

// func main() {
// 	EmptySlice := []string{}
// 	EmptySlice = append(EmptySlice, "a", "b", "c")
// 	fmt.Println(EmptySlice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 4 логических значений и инициализируйте его значениями true, false, true, false.
// package main

// import "fmt"

// func main() {
// 	slice := []bool{true, false, true, false}
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 6 целых чисел и инициализируйте его значениями от 1 до 6.

// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте пустой срез из строк и добавьте в него значения "hello", "world".
// package main

// import "fmt"

// func main() {
// 	EmptySlice := []string{}
// 	EmptySlice = append(EmptySlice, "hello", "world")
// 	fmt.Println(EmptySlice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 5 вещественных чисел и инициализируйте его значениями 1.1, 2.2, 3.3, 4.4, 5.5.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{1.1, 2.2, 3.3, 4.4, 6.6}
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 7 целых чисел и выведите его длину и вместимость.
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7}
// 	fmt.Println(len(slice), cap(slice))
// }

//-----------------------------------------------------------------------------------

// Создайте пустой срез из строк и добавьте в него 3 элемента, затем выведите его длину и вместимость.
// package main

// import "fmt"

// func main() {
// 	slice := []string{}
// 	slice = append(slice, "I'm", "a", "developer")
// 	fmt.Println(len(slice), cap(slice))
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 4 вещественных чисел и выведите его длину и вместимость.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{1.1, 2.2, 3.3, 4.4}
// 	fmt.Println(slice, len(slice), cap(slice))
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 5 целых чисел с помощью функции make.
// package main

// import "fmt"

// func main() {
// 	slice := make([]int, 5)
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 3 строк с длиной 3 и вместимостью 5.
// package main

// import "fmt"

// func main() {
// 	slice := make([]string, 3, 5)
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// // Создайте пустой срез из логических значений с вместимостью 4.
// package main

// import "fmt"

// func main() {
// 	EmptySlice := make([]bool, 0, 4)
// 	fmt.Println(EmptySlice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 3 целых чисел и добавьте в него значение 4.
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3}
// 	slice = append(slice, 4)
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте пустой срез из строк и добавьте в него значения "a", "b", "c".
// package main

// import "fmt"

// func main() {
// 	EmptySlice := []string{}
// 	EmptySlice = append(EmptySlice, "a", "b", "c")
// 	fmt.Println(EmptySlice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 2 вещественных чисел и добавьте в него 3 новых значения.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{1.1, 2.2}
// 	slice = append(slice, 3.3, 4.4, 5.5)
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 6 целых чисел и создайте срез, содержащий элементы с индексами от 1 до 4.
// package main

// import "fmt"

// func main() {
// 	arr := [6]int{-2, -1, 0, 1, 2, 3}
// 	slice := arr[1:5]
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 5 строк и создайте новый срез, содержащий последние 3 элемента.
// package main

// import "fmt"

// func main() {
// 	slice1 := []string{"a", "b", "c", "d", "e"}
// 	slice2 := slice1[len(slice1)-3:]
// 	fmt.Println(slice2)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 логических значений и создайте срез, содержащий первые 2 элемента.
// package main

// import "fmt"

// func main() {
// 	arr := [4]bool{true, true, false, true}
// 	slice := arr[:2]
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте два среза из 3 целых чисел и скопируйте значения первого среза во второй.
// package main

// import "fmt"

// func main() {
// 	sl1 := []int{100, 200, 300}
// 	sl2 := []int{10, 20, 30}
// 	sl3 := copy(sl2, sl1)
// 	fmt.Println(sl2)
// 	fmt.Println(sl3)
// }

//-----------------------------------------------------------------------------------

// Создайте два среза из 4 строк и скопируйте значения второго и третьего элемента первого среза во второй.
// package main

// import "fmt"

// func main() {
// 	slice1 := []string{"a", "b", "c", "d"}
// 	slice2 := []string{"e", "f", "g", "h"}
// 	copy(slice2,slice1[1:3])
// 	fmt.Println(slice2)
// }

//-----------------------------------------------------------------------------------

// Создайте два среза из 5 вещественных чисел и скопируйте все значения первого среза во второй.
// package main

// import "fmt"

// 	func main() {
// 		slice1 := []float64{0, 1.1, 2.2, 3.3, 4.4}
// 		slice2 := []float64{5.5, 6.6, 7.7, 8.8, 9.9}
// 		copy(slice2, slice1)
// 		fmt.Println(slice2)
// 	}

//-----------------------------------------------------------------------------------

// Создайте массив из 5 целых чисел, где значения второго и четвертого элемента равны 10 и 20 соответственно.
// package main

// import "fmt"

// func main() {
// 	arr := [5]int{-2, -1, 0, 1, 2}
// 	arr[1] = 10
// 	arr[3] = 20
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 6 строк, где первая и последняя строка равны "start" и "end" соответственно.
// package main

// import "fmt"

// func main() {
// 	arr := [6]string{"1", "-->you", "can", "run", "fast<--", "6"}
// 	arr[0] = "start"
// 	arr[cap(arr)-1] = "finish"
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 логических значений, где третий элемент равен true.
// package main

// import "fmt"

// func main() {
// 	arr := [4]bool{false, false, false, false}
// 	arr[2] = true
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 целых чисел и скопируйте его в новый массив, затем измените второй элемент нового массива.
// package main

// import "fmt"

// func main() {
// 	arr1 := [4]int{1, 2, 3, 4}
// 	arr2 := arr1
// 	arr2[1] = 0
// 	fmt.Println(arr2)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 5 строк и скопируйте его в новый массив, затем измените первый элемент нового массива.
// package main

// import "fmt"

// func main() {
// 	arr1 := [5]string{"C", "is", "a", "cute", "language"}
// 	arr2 := arr1
// 	arr2[0] = "Go"
// 	fmt.Println(arr2)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 3 вещественных чисел и скопируйте его в новый массив, затем измените последний элемент нового массива
// package main

// import "fmt"

// func main() {
// 	arr1 := [3]float64{0, 1.1, 2.2}
// 	arr2 := arr1
// 	arr2[2] = 3.3
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 6 целых чисел и создайте срез, содержащий элементы с индексами от 1 до 3. Выведите длину и вместимость среза.
// package main

// import "fmt"

// func main() {
// 	arr := [6]int{1, 2, 3, 4, 5, 6}
// 	slice := arr[1:4]
// 	fmt.Println(len(slice), cap(slice))
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 5 строк и создайте новый срез, содержащий элементы с индексами от 2 до 4. Выведите длину и вместимость нового среза.
// package main

// import "fmt"

// func main() {
// 	slice1 := []string{"a", "b", "c", "d", "e"}
// 	slice2 := slice1[2:5]
// 	fmt.Println(slice2)
// 	fmt.Println(len(slice2), cap(slice2))
// }

//-----------------------------------------------------------------------------------

// Создайте массив из 4 логических значений и создайте срез, содержащий элементы с индексами от 0 до 2.
// Измените значение первого элемента среза и выведите исходный массив.
// package main

// import "fmt"

// func main() {
// 	arr := [4]bool{true, true, false, true}
// 	slice := arr[0:3]
// 	slice[0] = false
// 	fmt.Println(arr)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 3 целых чисел и добавьте в него значение 4, затем добавьте еще одно значение 5.
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3}
// 	slice = append(slice, 4)
// 	slice= append(slice, 5)
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Создайте пустой срез из строк и добавьте в него значения "x", "y", "z".
// package main

// import "fmt"

// func main() {
// 	EmptySlice := []string{}
// 	EmptySlice = append(EmptySlice, "x", "y", "z")
// 	fmt.Println(EmptySlice)
// }

//-----------------------------------------------------------------------------------

// Создайте срез из 2 вещественных чисел и добавьте в него значения 3.3, 4.4, 5.5.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{1.1, 2.2}
// 	slice = append(slice, 3.3, 4.4, 5.5)
// 	fmt.Println(slice)
// }