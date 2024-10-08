//КАК ПЕЧАТАТЬ? !!!---------------------------------------------------------------
//
// package main

// import "fmt"

// func main() {
//fmt.Println("Hello Go!");fmt.Println("Hello Golang!");fmt.Println("Hello Go!")

// 	fmt.Print("hello\n")

// 	fmt.Print("hello")

//fmt.Print("FizzBuzz", " ")

// 	fmt.Printf("Наибольший общий делитель чисел %d и %d равен: %d\n", 1, 2, 3)

//fmt.Printf("Имя: %s, Возраст: %d лет\n", "Shodmon", 23)

//Если мы хотим сразу определить несколько переменных и присвоить им начальные
//значения, то можно обернуть их в скобки:
// package main
// import "fmt"

// func main() {
//     var (
//         name string = "Tom"
//         age int = 27
//     )

//     fmt.Println(name)   // Tom
//     fmt.Println(age)    // 27
// }
//------------------------------------------------------------------------------------

// КОММЕНТАРИИ

// Однострочный комментарий располагается в одну строку после двойного слеша . Все, что идет после этих символов, воспринимается компилятором как комментарий. Многострочный комментарий заключается между символами /* и */ и может занимать несколько строк:

/*
   Первая программа
   на языке Go
*/
// package main    // определение пакета для текущего файла
// import "fmt"    // подключение пакета fmt

// // определение функции main
// func main() {
//     fmt.Println("Hello Go!")    // вывод строки на консоль
// }

//------------------------------------------------------------------------------------

// Бесконечный цыкл или процедура
//
// package main

// import "fmt"

//	func main() {
//		newFunc()
//	}
//
//	func newFunc()  {
//		fmt.Print("hello\n")
//		newFunc()
//	}
//
// -----------------------------------------------------------------------------------

// НОД двух чисел
//
// package main

// import "fmt"

// var a, b, c int

// func main() {
// 	fmt.Scan(&a, &b)

//		for b != 0 {
//			c = a
//			a = b
//			b = c % b
//		}
//		fmt.Println(c)
//	}
//
// -----------------------------------------------------------------------------------

//БЕСКОНЕЧНЫЙ FOR
//
// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Print("Hello", "      ")
// 	}
// }
//------------------------------------------------------------------------------------

// НАХОЖДЕНИЕ ПРОСТЫХ ЧИСЕЛ
//
// package main

// import (
// 	"fmt"
// )

//	func main() {
//			var number int
//			fmt.Scan(&number)
//			fmt.Println(isPrime(number))
//		}
//
//	func isPrime(n int) bool {
//		if n == 1 {
//			return false
//		}
//		for i := 2; i*i <= n; i++ {
//			if n%i == 0 {
//				return false
//			}
//		}
//		return true
//	}

//------------------------------------------------------------------------------------

////НЕ ЗАБУДЬ: дефолтное значение ДЕФОЛТНОЕ ЗНАЧЕНИЕ ДЛЯ "BOOLEAN"- ЭТО "FALSE"

// -----------------------------------------------------------------------------------

//ТАБЛИЦА УМНОЖЕНИЯ
//
// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	fmt.Println("ТАБЛИЦА УМНОЖЕНИЯ\n")
// 	for i := 0; i <= 10; i++ {
// 		for j := 0; j <= n; j++ {
// 			fmt.Print(j, "x", i, "=", i*j)
// 			if i != 10 {
// 				if i*j > 9 {
// 					fmt.Print("     ")
// 				}
// 				if i*j <= 9 {
// 					fmt.Print("      ")
// 				}
// 			} else {
// 				if i*j > 9 {
// 					fmt.Print("    ")
// 				}
// 				if i*j <= 9 {
// 					fmt.Print("     ")
// 				}
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

//------------------------------------------------------------------------------------

// СЛОЖЕНИЕ В  ==>> sring
//
// Напишите функцию changeString, которая принимает указатель на строку и
// добавляет к ней слово "Go".
// package main

// import "fmt"

// func main() {
// 	a := "lets's "
// 	changeString(&a)
// 	fmt.Println(a)

// }
//
//	func changeString(ptr *string) {
//		*ptr = *ptr + "Go"
//	}
//
// -----------------------------------------------------------------------------------

//РАБОТА С УКАЗАТЕЛЕМ
//
// package main

// import "fmt"

// func main() {
// 	num := 10
// 	fmt.Println(num)

//		ptrNum := &num
//		sum(ptrNum)
//		fmt.Println(num)
//	}
//
//	func sum(ptrNum *int) {
//		*ptrNum = *ptrNum - 1
//	}
//
// -----------------------------------------------------------------------------------

////Реализация и Указатель
//
// package main

// import "fmt"

// func main() {
// 	var name string
// 	var age int
// 	namePtr := &name
// 	agePtr := &age
// 	storeName(namePtr, "Шодмон")
// 	storeAge(agePtr, 23)
// 	printPerson(namePtr, agePtr)
// }
// func storeName(namePtr *string, name string) {
// 	*namePtr = name
// }

// func storeAge(agePtr *int, age int) {
// 	*agePtr = age
// }

//	func printPerson(namePtr *string, agePtr *int) {
//		fmt.Printf("Имя: %sец, Возраст: %d лет\n", *namePtr, *agePtr)
//	}
//
// -----------------------------------------------------------------------------------

//ЗАДАЧА НА ТЕМУ ТИПЫ
//
// Применение функции к числу
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функции для удвоения
// и утроения числа. Используйте тип UnaryOperation для вызова этих функций.
// package main

// import "fmt"

// type UnaryOperation func(int) int

// // Функция для удвоения числа
// func double(a int) int {
// 	return a * 2
// }

// // Функция для утроения числа
// func triple(a int) int {
// 	return a * 3
// }
// func main() {
// 	var op2 UnaryOperation
// 	op2 = double
// 	fmt.Println("Double:", op2(4))
// 	op2 = triple
// 	fmt.Println("Triple:", op2(3))
// }

//------------------------------------------------------------------------------------

// Тип+int ==>> не работает.     Но псевдоним+int==>> работает
//------------------------------------------------------------------------------------

//ЗАДАЧА НА ТЕМУ ПСЕВДОНИМ
//
// Применение скидки
// Создайте псевдоним для типа float64 под названием Price.
// Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.
// package main

// import "fmt"

// type Price = float64

// func newPrice(a Price) Price {
// 	return 9 * a / 10
// }
// func main() {
// 	var a Price = 1000
// 	fmt.Println(newPrice(a))
// }

//-----------------------------------------------------------------

//Задача на тему "структура"
//
// package main

// import "fmt"

// // Person
// // Определяем структуру
// type Person struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		// Создаем и инициализируем структуру
//		p := Person{Name: "Alice", Age: 30}
//		fmt.Println(p)
//	}
//
// -----------------------------------------------------------------------------------

//     Структуры можно создавать и инициализировать несколькими способами.
// - Использование литерала структуры:

// p := Person{Name: "Alice", Age: 30}

// - Без указания имен полей (порядок важен):

// p := Person{"Alice", 30}

// - Создание структуры с использованием ключевого слова new:

// p := new(Person)
// p.Name = "Alice"
// p.Age = 30

// - Создание структуры через указатель:

// p := &Person{Name: "Alice", Age: 30}
//Обращение к полям структуры
//    Обращение к полям структуры осуществляется с помощью оператора точки (.).
// package main

// import "fmt"

// type Person struct {
//     Name string
//     Age  int
// }

// func main() {
//     p := Person{Name: "Alice", Age: 30}
//     fmt.Println(p.Name)
//     fmt.Println(p.Age)
// }
// Указатели на структуры
// Указатели на структуры позволяют изменять значения полей структуры без необходимости создания копии структуры.
// package main

// import "fmt"

// type Person struct {
//     Name string
//     Age  int
// }

// func main() {
//     p := &Person{Name: "Alice", Age: 30}
//     p.Age = 31
//     fmt.Println(p)
// }
//Вложенные структуры
//Структуры могут содержать другие структуры в качестве полей, что позволяет создавать более сложные данные.
//  package main

//  import "fmt"

//  type Address struct {
// 	 City   string
// 	 Street string
//  }
//  type Student struct {
// 	 Name    string
// 	 Age     int
// 	 Address Address
//  }

//	 func Inf(a Student) {
//	 fmt.Printf("Name:%s, Age:%d\n",a.Name, a.Age )
//	 fmt.Printf("Address:%s, %s",a.Address.City,a.Address.Street)
//	 }
//	 func main() {
//		 a := Student{
//			 Name: "Shodmon",
//			 Age:  23,
//			 Address: Address{
//				 City:   "Dushanbe",
//				 Street: "Nusratullo Makhsum",
//			 },
//		 }
//		 Inf(a)
//	 }
//
//Хранение ссылки на структуру того же типа

//Структуры могут содержать указатели на другие структуры того же типа, что позволяет создавать
//связные структуры данных, такие как деревья или графы.

//Задача
// Дерево сотрудников
// Создайте структуру Employee с полем Manager,
// которое указывает на другого сотрудника (вложенная структура).
// Напишите функцию, которая рекурсивно выводит информацию о сотруднике и его менеджере.
// package main

// import "fmt"

// type Employee2 struct {
// 	Name     string
// 	Position string
// 	Manager  *Employee2
// }

// // Функция для рекурсивного вывода информации о сотруднике и его менеджере
// func printEmployeeHierarchy(e *Employee2) {
// 	if e == nil {
// 		return
// 	}
// 	fmt.Printf("Name: %s, Position: %s\n", e.Name, e.Position)
// 	if e.Manager != nil {
// 		fmt.Printf("Manager: %s\n", e.Manager.Name)
// 		printEmployeeHierarchy(e.Manager)
// 	}
// }
// func main() {
// 	manager := &Employee2{Name: "Alice", Position: "Senior Manager"}
// 	employee2 := &Employee2{Name: "Bob", Position: "Junior Developer", Manager: manager}
// 	printEmployeeHierarchy(employee2)
// }
// Анонимные структуры

// Анонимные структуры позволяют создавать структуры без явного определения типа. Они полезны для
// создания временных структур или структур, используемых в ограниченном контексте.

// package main

// import "fmt"

// func main() {
//     person := struct {
//         Name string
//         Age  int
//     }{
//         Name: "Alice",
//         Age:  30,
//     }
//     fmt.Println(person)
// }

// -----------------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{10, 20, 30, 40, 50}
// 	length := len(arr)
// 	capasity := cap(arr)
// 	fmt.Println(length, capasity)
// }

// package main

// import "fmt"

//	func main() {
//		arr := [5]int{1, 2, 3, 4, 5}
//		slice := arr[1:4]
//		fmt.Println(slice)
//	}

// package main

// import "fmt"

// func main() {
// 	var EmptySlice []int
// 	fmt.Println(EmptySlice)
// }

// package main

// import "fmt"

//	func main() {
//		EmptySlice := []int{}
//		fmt.Println(EmptySlice)
//	}

// package main

// import "fmt"

//	func main() {
//		slice:=[]int{1,2,3,4,5}
//		fmt.Println(len(slice))
//		fmt.Println(cap(slice))
//	}
//
// ЗАДАЧА
// Создайте срез из 5 строк и создайте новый срез, содержащий последние 3 элемента.
// package main

// import "fmt"

//	func main() {
//		arr := []string{"a", "b", "c", "d", "e"}
//		arr2 := arr[len(arr)-3:]
//		fmt.Println(arr2) // [d e]
//	}
//
// ЗАДАЧА
// Создайте массив из 4 логических значений и создайте срез, содержащий первые 2 элемента.
// package main

// import "fmt"

// func main() {
// 	arr := [4]bool{true, false, true, false}
// 	arr2 := arr[:2]
// 	fmt.Println(arr2) // [true false]
// }

// ЗАДАЧА
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	copy(slice[1:], slice)
// 	fmt.Println(slice) // Вывод: [1 1 2 3 4]
// }

// ЗАДАЧА
// Создайте массив из 6 целых чисел и создайте срез, содержащий элементы с индексами от 1 до 3. Выведите длину и вместимость среза.
// package main

// import "fmt"

// func main() {
// 	arr := [6]int{1, 2, 3, 4, 5, 6}
// 	slice := arr[1:4]
// 	fmt.Println(slice)
// 	fmt.Println(len(slice), cap(slice))
// }

// ЗАДАЧА
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	slice2 := slice[1:3]
// 	fmt.Println(slice2)
// 	fmt.Println(len(slice2), cap(slice2))
// }

// -----------------------------------------------------------------------------------

// ЗАДАЧА
// Создать копию массива.

// package main

// import "fmt"

// func copyArray(arr []int) []int {
// 	copyArr := make([]int, len(arr))
// 	copy(copyArr, arr)
// 	return copyArr
// }
// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	fmt.Println(copyArray(arr)) // [1 2 3 4 5]
// }

//-----------------------------------------------------------------------------------

//ЗАДАЧА
//ОДАМКАШОНИ ВА ЧУНИН МОШИН
// package main

// import "fmt"

// type car struct {
// 	color string
// 	brend string
// 	model string
// 	count int
// }

//	func (c car) odamKashoni(d int) int {
//		result := 1 + d + c.count
//		return result
//	}
//
//	func main() {
//		var c car = car{
//			count: 1,
//		}
//		fmt.Println(c.odamKashoni(3))
//	}

// -------------------------------------------------------------------------------------------------

// Вариант 17
// Напишите функцию, которая принимает срез структур Animal и возвращает структуру
//Animal, у которой средний возраст наибольший.

// package main

// import (
// 	"fmt"
// )

// type Animal struct {
// 	Name string
// 	Age  int
// }

// func AnimalWithHighestAverageAge(animals []Animal) Animal {
// 	if len(animals) == 0 {
// 		return Animal{}
// 	}

// 	highestAverageAgeAnimal := animals[0]
// 	highestAverageAge := float64(highestAverageAgeAnimal.Age)

// 	for _, animal := range animals {
// 		averageAge := float64(animal.Age)
// 		if averageAge > highestAverageAge {
// 			highestAverageAge = averageAge
// 			highestAverageAgeAnimal = animal
// 		}
// 	}

// 	return highestAverageAgeAnimal
// }

// func main() {
// 	animals := []Animal{
// 		{"Lion", 10},
// 		{"Elephant", 25},
// 		{"Giraffe", 15},
// 		{"Tiger", 12},
// 	}

// 	result := AnimalWithHighestAverageAge(animals)
// 	fmt.Printf("Animal with the highest average age: %s, Age: %d\n", result.Name, result.Age)
// }

//-----------------------------------------------------------------------------------

// Обрабо1.Изменение среза внутри функции
// package main

// import (
// 	"fmt"
// )

// func modifySlice(s []int) {
// 	for i := range s {
// 		s[i] = s[i] * 2
// 	}
// }

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Println("Before:", slice)
// 	modifySlice(slice)
// 	fmt.Println("After:", slice)
// }

// 2. Добавление элементов в срез
// package main

// import (
// 	"fmt"
// )

// func appendToSlice(s []int) []int {
// 	s = append(s, 10)
// 	fmt.Println("Inside function:", s)
// 	return s
// }

// func main() {
// 	slice := []int{1, 2, 3}
// 	fmt.Println("Before:", slice)
// 	slice = appendToSlice(slice)
// 	fmt.Println("After:", slice)
// }

// 3. Передача среза без изменения исходного среза
// package main

// import (
// 	"fmt"
// )

// func modifySliceCopy(s []int) {
// 	for i := range s {
// 		s[i] = s[i] * 2
// 	}
// }

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Println("Before:", slice)
// 	sliceCopy := append([]int(nil), slice...)
// 	modifySliceCopy(sliceCopy)
// 	fmt.Println("After:", slice)
// 	fmt.Println("Modified Copy:", sliceCopy)
// }

// package main

// import "fmt"

// type Rectangle struct {
// 	width, height float64
// }

// // Area Public метод
// func (r Rectangle) Area() float64 {
// 	return r.width * r.height
// }

// // private метод
// func (r *Rectangle) scale(factor float64) {
// 	r.width *= factor
// 	r.height *= factor
// }

// func main() {
// 	rect := Rectangle{width: 10, height: 5}
// 	fmt.Println("Area:", rect.Area()) // Output: Area: 50

// 	// Вызов приватного метода изнутри пакета
// 	rect.scale(2)
// 	fmt.Println("Scaled Area:", rect.Area()) // Output: Scaled Area: 200
// }

//-------------------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// s := "Hello World"
// c := s[1]
// fmt.Println(c)
// sub := s[6:11]
// fmt.Println(sub)
// s1 := "abc"
// s2 := "def"
// fmt.Println(s1 == s2)
// fmt.Println(s1 < s2)
// fmt.Println(s1 > s2)
// fmt.Println(strings.Contains(s, "World"))
// fmt.Println(strings.Index(s, "Word"))    // 6 (начальный индекс)
// fmt.Println(strings.Contains(s, "W"))
// s:="Hello , 世界"
// s1 := "Hello"
// s2 := "Go"
// s := s1 + " " + s2
// fmt.Println(s)
// s := "Hello Go Hello Go Hello Go Hello Stop Hello Go Stop"
// count := strings.Count(s, "Stop")
// fmt.Println(count)
// hasprefix := strings.HasPrefix(s, "Hello")
// fmt.Println(hasprefix)
// hassuffix := strings.HasSuffix(s, "Stop")
// fmt.Println(hassuffix)
// index := strings.Index(s, "S")
// fmt.Println(index)
// lastindex := strings.LastIndex(s, "")
// fmt.Println(lastindex)
// parts := []string{"Hello", "Go", "Hello", "Go", "Hello", "Go", "Hello", "Stop"}
// joined := strings.Join(parts, ", ")
// fmt.Println(joined)
// s := "Bashir"
// repeat := strings.Repeat(s, 2)
// fmt.Println(repeat)
// s := "Hello Go Hello Go Hello Go Hello Stop Hello Go Stop"
// replaced := strings.Replace(s, "Hello", "Asalom",2)
// fmt.Println(replaced)
// replacedall := strings.ReplaceAll(s, "Go", "Walk")
// fmt.Println(replacedall)
// s := "abcd"
// split := strings.Split(s, "")
// fmt.Println(split)
// split = strings.SplitN(s, "", 3)
// fmt.Println(split)
// s := "HeLlO wOrLd"
//
//	lower := strings.ToLower(s)
//	fmt.Println(lower)
//	upper := strings.ToUpper(s)
//	fmt.Println(upper)
//
// trimmed := strings.Trim(s, " ")
// fmt.Println(trimmed)
// trimmed := strings.TrimSpace(s)
// fmt.Println(trimmed)
// prefix := "HeLlO"
// suffix := "wOrLd"
// trimprefix := strings.TrimPrefix(s, prefix)
// trimsuffix := strings.TrimSuffix(s, suffix)
// fmt.Println(trimprefix)
// fmt.Println(trimsuffix)
//
//		s = "Hello, World!"
//		newS := strings.Replace(s, "World", "Go", -1)
//		fmt.Println(newS) // "Hello, Go!"
//	}
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) Introduce() {
// 	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
// }

// type Employee struct {
// 	Person
// 	position string
// }

// // Introduce
// // Переопределение метода Introduce для Employee
//
//	func (e Employee) Introduce() {
//		fmt.Printf("Hi, I'm %s, a %s, and I'm %d years old.\n", e.name, e.position, e.age)
//	}
//
//	func main() {
//		p := Person{name: "Alice", age: 30}
//		e := Employee{Person: Person{name: "Bob", age: 40}, position: "Manager"}
//		p.Introduce() // Output: Hi, I'm Alice and I'm 30 years old.
//		e.Introduce() // Output: Hi, I'm Bob, a Manager, and I'm 40 years old.
//	}
// package main

// import (
// 	"fmt"
// 	"unicode"
// )

//	func main() {
//		r := 'a'
//		fmt.Println(string(unicode.ToUpper(r))) // "A"
//		r = 'A'
//		fmt.Println(string(unicode.ToLower(r))) // "a"
//		fmt.Println(unicode.ToUpper(r))
//		fmt.Println(unicode.ToLower(r))
//		s:="Abdussabur"
//		fmt.Println(s[0])
//	}

// package main

// import "fmt"

// func main() {
// 	s := "Hello "
// 	c := s[5]
// 	bytes := []byte(s)
// 	fmt.Println(bytes)
// 	fmt.Println(c)
// }

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	s := "Hello, 世界"
// 	// Длина строки в байтах
// 	fmt.Println(len(s)) // 13
// 	// Длина строки в рунах (символах)
// 	fmt.Println(utf8.RuneCountInString(s)) // 9
// 	// Итерация по рунам в строке
// 	for i, r := range s {
// 		fmt.Printf("%d: %c\n", i, r)
// 	}
// }

// Этот пример показывает, как строка "Hello, 世界" занимает 13 байт, но содержит 9
// символов (рун). Итерация по строке с использованием range позволяет корректно
// обрабатывать каждый символ Unicode независимо от его длины в байтах.

// package main

// import "fmt"

// func main() {
// 	str:="ABCD"
// 	// fmt.Println(str[0])
// 	// fmt.Printf("%b\n",str[1])
// 		fmt.Printf("type:%T,value:%v\n",str,str)
// }

// -----------------------------------------------
// package main

// import "fmt"

// // Интерфейс PointerOperations с методами Increment и Decrement
// type PointerOperations interface {
// 	Increment()
// 	Decrement()
// }

// // Структура IntPointer с указателем на целое число
// type IntPointer struct {
// 	Int *int
// }

// // Метод Increment увеличивает значение на 1
// func (i *IntPointer) Increment() {
// 	if i.Int != nil {
// 		(*i.Int)++ // Увеличение значения по указателю
// 	}
// }

// // Метод Decrement уменьшает значение на 1
// func (i *IntPointer) Decrement() {
// 	if i.Int != nil {
// 		(*i.Int)-- // Уменьшение значения по указателю
// 	}
// }

// // Функция A, принимающая интерфейс PointerOperations
// func A(po PointerOperations) {
// 	po.Decrement()
// 	po.Increment()
// }

// // Функция B, создающая объект IntPointer и вызывающая функцию A
// func B() {
// 	val := 63
// 	ip := IntPointer{Int: &val}
// 	A(&ip) // Передача указателя на IntPointer
// 	fmt.Println(*ip.Int) // Вывод изменённого значения
// }

// // Основная функция
// func main() {
// 	B()
// }
//----------------------------------------------
/////////////////УКАЗАТЕЛЬ/////////////////

// package main

// import "fmt"

// func main() {
// 	var pointer *int
// 	fmt.Println(pointer)
// 	intt:=65
// 	pointer=&intt
// 	fmt.Println(pointer)
// 	fmt.Println(*pointer)
// }
//--------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	parts := []string{"Hello", "Go", "Hello", "Go", "Hello", "Go", "Hello", "Stop"}
// 	joined := strings.Join(parts, ", ")
// 	fmt.Println(joined)
// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

// countCharacters принимает имя файла, считывает его содержимое и возвращает количество символов
func countCharacters(fileName string) (int, error) {
	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Создаем буфер для чтения файла
	reader := bufio.NewReader(file)

	var charCount int

	// Читаем файл по строкам
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				// Достигнут конец файла
				break
			}
			return 0, err
		}
		charCount += len(line)
		// Добавляем количество символов новой строки (если требуется)
		// charCount += len("\n")
	}

	return charCount, nil
}

func main() {
	// Пример использования функции
	fileName := "example.txt" // Замените на имя вашего файла
	count, err := countCharacters(fileName)
	if err != nil {
		fmt.Println("Ошибка при подсчете символов:", err)
		return
	}
	fmt.Printf("Количество символов в файле '%s': %d\n", fileName, count)
}
