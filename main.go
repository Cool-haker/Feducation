package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func secondMax(slice []int) (int, error) {
	if len(slice) < 2 {
		return 0, errors.New("there are less that two values in the slice")
	}

	var firstMax int = slice[0]
	var secondMax int = slice[1]

	if slice[0] < slice[1] {
		firstMax = slice[1]
		secondMax = slice[0]
	}

	for _, num := range slice[2:] {
		if num > firstMax {
			secondMax = firstMax
			firstMax = num
			continue
		}

		if num > secondMax && num != firstMax || firstMax == secondMax {
			secondMax = num
		}
	}

	return secondMax, nil
}

// func main() {
// 	fmt.Println(secondMax([]int{5, 7, 9, 11, 14, 16}))                                //14
// 	fmt.Println(secondMax([]int{16, 14, 11, 9, 7, 5}))                                //14
// 	fmt.Println(secondMax([]int{10, 20}))                                             //10
// 	fmt.Println(secondMax([]int{20, 10}))                                             //10
// 	fmt.Println(secondMax([]int{5, 5, 1}))                                            //1
// 	fmt.Println(secondMax([]int{10, 5, 10}))                                          //5
// 	fmt.Println(secondMax([]int{-10, -5, 0}))                                         //-5
// 	fmt.Println(secondMax([]int{1, 500, 2, 300, 3, 1000}))                            //500
// 	fmt.Println(secondMax([]int{-500, -50, -10, -2}))                                 //-10
// 	fmt.Println(secondMax([]int{-102308042342, 5349785493, 10, -1928439, 12948, 12})) //12948
// 	fmt.Println(secondMax([]int{5, 5, 5}))                                            //5
// }

func mergeAndSort(arr1, arr2 []int) []int {
	for _, arr := range arr2 {
		arr1 = append(arr1, arr)
	}

	sort.Ints(arr1)

	return arr1
}

// func main() {
// 	fmt.Println(mergeAndSort([]int{3, 1, 5}, []int{4, 2, 6})) // [1 2 3 4 5 6]
// 	fmt.Println(mergeAndSort([]int{8, 2, 0}, []int{7, 3, 1})) // [0 1 2 3 7 8]
// }

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println(isLeapYear(2020)) // true
// 	fmt.Println(isLeapYear(1900)) // false
// 	fmt.Println(isLeapYear(2000)) // true
// }

func maxOfThree(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else if c > b && c > a {
		return c
	}
	return a
}

// func main() {
// 	fmt.Println(maxOfThree(3, 7, 5))     // 7
// 	fmt.Println(maxOfThree(10, 2, 8))    // 10
// 	fmt.Println(maxOfThree(-1, -5, -10)) // -1
// 	fmt.Println(maxOfThree(10, 10, 10))  // 10
// 	fmt.Println(maxOfThree(10, 10, 11))  // 11
// 	fmt.Println(maxOfThree(10, 11, 10))  // 11
// 	fmt.Println(maxOfThree(11, 10, 10))  // 11
// 	fmt.Println(maxOfThree(1, 2, 80))    // 80
// }

func reverse(line string) string {
	length := len(line)

	result := make([]rune, length)

	for i, l := range line {
		result[length-i-1] = l
	}

	return string(result)
}

// func main() {
// 	fmt.Println(reverse("hello"))  // "olleh"
// 	fmt.Println(reverse("Привет")) // "тевирП"
// }

func charCase(r rune) string {
	switch {
	case r >= 'a' && r <= 'z':
		return "Латинская строчная"
	case r >= 'A' && r <= 'Z':
		return "Латинская заглавная"
	case r >= 'а' && r <= 'я':
		return "Кириллическая строчная"
	case r >= 'А' && r <= 'Я':
		return "Кириллическая заглавная"
	default:
		return "Вы не понимаете, это другое!"
	}
}

// func main() {
// 	fmt.Println(charCase('A')) // "Латинская заглавная"
// 	fmt.Println(charCase('я')) // "Кириллическая строчная"
// 	fmt.Println(charCase('Б')) // "Кириллическая заглавная"
// 	fmt.Println(charCase('1')) // "Другое"
// 	fmt.Println(charCase('😊')) // "Другое"
// }

func calculator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0
		}
		return a / b
	default:
		return 0
	}
}

// func main() {
// 	fmt.Println(calculator(10, 5, "+")) // 15
// 	fmt.Println(calculator(10, 5, "-")) // 5
// 	fmt.Println(calculator(10, 5, "*")) // 50
// 	fmt.Println(calculator(10, 0, "/")) // 0
// }

func sumToN(n int) int {
	r := 0
	for i := 1; i <= n; i++ {
		r += i
	}
	return r
}

// func main() {
// 	fmt.Println(sumToN(5))  // 15 (1 + 2 + 3 + 4 + 5)
// 	fmt.Println(sumToN(10)) // 55
// }

func reverseNumber(num int) (int, error) {
	strNum := strconv.Itoa(num)
	length := len(strNum)

	if num < 0 {
		result := make([]rune, length)

		for i, r := range strNum[1:] {
			result[length-i-1] = r
		}
		result[0] = '-'

		res, err := strconv.Atoi(string(result))
		if err != nil {
			return 0, fmt.Errorf("Nu pizdec")
		}
		return res, nil
	} else {
		resultt := make([]rune, length)

		for i, r := range strNum {
			resultt[length-i-1] = r
		}

		res, err := strconv.Atoi(string(resultt))
		if err != nil {
			return 0, err
		}
		return res, nil
	}
}

// func main() {
// 	fmt.Println(reverseNumber(123))  // 321
// 	fmt.Println(reverseNumber(500))  // 5
// 	fmt.Println(reverseNumber(-987)) // -789
// 	fmt.Println(reverseNumber(-100)) // -1
// }

func excrement() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// func main() {
// 	inc := excrement()
// 	fmt.Println(inc()) //1
// 	fmt.Println(inc()) //2
// 	fmt.Println(inc()) //3
// 	fmt.Println(inc()) //4

// 	incr := excrement()
// 	fmt.Println(incr()) //1
// 	fmt.Println(incr()) //2
// 	fmt.Println(incr()) //3

// 	fmt.Println(inc()) //5
// }

func powerFunc(power int) func(int) float64 {
	return func(num int) float64 {
		return math.Pow(float64(num), float64(power))
	}
}

// func main() {
// 	square := powerFunc(2)
// 	fmt.Println(square(4))  // 16
// 	fmt.Println(square(5))  // 25
// 	fmt.Println(square(-5)) // 25
// 	cube := powerFunc(3)
// 	fmt.Println(cube(2))  // 8
// 	fmt.Println(cube(3))  // 27
// 	fmt.Println(cube(-3)) // -27
// }

func swap(a, b *int) {
	*a, *b = *b, *a
}

// func main() {
// 	x, y := 10, 20
// 	swap(&x, &y)
// 	fmt.Println(x, y) // 20 10
// }

func increment(ptr *int) {
	*ptr += 1
}

// func main() {
// 	n := 5
// 	increment(&n)
// 	fmt.Println(n) //6
// 	increment(&n)
// 	increment(&n)
// 	increment(&n)
// 	fmt.Println(n) //9
// }

func appendValue(slice *[]int, value int) {
	*slice = append(*slice, value)
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	appendValue(&nums, 4)
// 	fmt.Println(nums) // [1 2 3 4]
// 	mk := []int{}
// 	appendValue(&mk, 5)
// 	fmt.Println(mk) // [5]

// }

func removeAtIndex(arr []int, index int) []int {
	if index < 0 || index > len(arr)-1 || len(arr) == 0 {
		return arr
	}

	return append(arr[:index], arr[index+1:]...)
}

// func main() {
// 	fmt.Println(removeAtIndex([]int{1, 2, 3, 4, 5}, 2)) // [1 2 4 5]
// 	fmt.Println(removeAtIndex([]int{10, 20, 30}, 5))    // [10 20 30]
// }

func safeDivide(a, b int) int {
	defer func() int {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		return 0
	}()

	if b == 0 {
		panic("поломалась :(")
	}

	return a / b
}

// func main() {
// 	fmt.Println(safeDivide(10, 2)) // 5
// 	fmt.Println(safeDivide(10, 0)) // 0
// }

func safeFunction(f func()) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
	}()

	f()
}

// func main() {
// 	safeFunction(func() {
// 		fmt.Println("Дела делаются, мутки мутятся...")
// 		panic("ОМОН пришёл, мутки накрыл")
// 	})
// }

func wordCount(text string) map[string]int {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	final := make(map[string]int, len(text))

	for _, word := range words {
		cleanWord := strings.Trim(word, ",.!?;:")
		if unicode.IsLetter([]rune(cleanWord)[0]) {
			final[cleanWord]++
		}
	}
	return final
}

// func main() {
// 	fmt.Println(wordCount("Hello crazy world. I want to say you Z. Yes, hello, yes, I'am crazy too"))
// }

func invertMap(m map[string]int) map[int]string {
	newMap := make(map[int]string, len(m))
	for i, val := range m {
		if _, ok := newMap[val]; ok {
			continue
		}
		newMap[val] = i
	}
	return newMap
}

// func main() {
// 	m := map[string]int{
// 		"a": 1,
// 		"b": 2,
// 		"c": 1,
// 	}
// 	fmt.Println(invertMap(m))
// }

func mergeMap(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, val := range m1 {
		if _, ok := m2[key]; ok {
			result[key] = val + m2[key]
		} else {
			result[key] = val
		}
	}
	for key, val := range m2 {
		if _, ok := result[key]; ok {
			continue
		} else {
			result[key] = val
		}
	}

	return result
}

// func main() {
// 	m1 := map[string]int{
// 		"banana": 5,
// 		"apple":  3,
// 		"tomato": 1,
// 	}
// 	m2 := map[string]int{
// 		"orange": 4,
// 		"tomato": 4,
// 		"banana": 3,
// 	}
// 	fmt.Println(mergeMap(m1, m2))
// }

type Rectangle1 struct {
	Widht  float64
	Height float64
}

func (r Rectangle1) Area() float64 {
	return r.Height * r.Widht
}

// func main() {
// 	r := Rectangle1{Widht: 5, Height: 10}
// 	fmt.Println(r.Area())
// }

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products map[int]Product
}

func NewInventory() *Inventory {
	return &Inventory{
		make(map[int]Product),
	}
}

func (i Inventory) AddProduct(p Product) {
	i.Products[p.ID] = p
}

func (i *Inventory) SellProduct(id, qty int) error {
	product, ok := i.Products[id]
	if !ok {
		return fmt.Errorf("А такого: %d нету(", id)
	}
	if product.Quantity < qty {
		return fmt.Errorf("А столько нету. Запрошено: %d, а есть только: %d", qty, product.Quantity)
	}

	i.Products[id] = Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity - qty,
	}
	return nil
}

// func main() {
// 	inv := NewInventory()
// 	inv.AddProduct(Product{ID: 1, Name: "Телефон", Price: 30000, Quantity: 10})
// 	fmt.Println(inv)
// 	fmt.Println(inv.SellProduct(1, 3)) //nil
// 	fmt.Println(inv.SellProduct(1, 8)) // Error: недостаточно
// }

type Ticket struct {
	ID            int
	PassengerName string
	Destination   string
}

type BookingSystem struct {
	Tickets map[int]Ticket
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		make(map[int]Ticket),
	}
}

func (b *BookingSystem) BookTicket(id int, name, destination string) error {
	_, ok := b.Tickets[id]
	if ok {
		return fmt.Errorf("А такой билет уже есть")
	}

	b.Tickets[id] = Ticket{
		ID:            id,
		PassengerName: name,
		Destination:   destination,
	}

	return nil
}

func (b *BookingSystem) CancelTicket(id int) error {
	_, ok := b.Tickets[id]
	if !ok {
		return fmt.Errorf("А такого билета и нету")
	}

	delete(b.Tickets, id)
	return nil
}

func (b *BookingSystem) GetTicket(id int) (Ticket, error) {
	_, ok := b.Tickets[id]
	if !ok {
		return Ticket{}, fmt.Errorf("Нет такого билета(")
	}

	return b.Tickets[id], nil
}

// func main() {
// 	bs := NewBookingSystem()
// 	bs.BookTicket(1, "Иван", "Москва")
// 	fmt.Println(bs.GetTicket(1))
// 	fmt.Println(bs.CancelTicket(1))
// 	fmt.Println(bs.GetTicket(1))
// }

type TextProcessor struct {
	Text string
}

func (t *TextProcessor) WordCount() map[string]int {
	text := strings.ToLower(t.Text)
	words := strings.Fields(text)
	final := make(map[string]int, len(words))
	for _, val := range words {
		cleanWord := strings.Trim(val, ",.!?:;")
		final[cleanWord] += 1
	}

	return final
}

func (t *TextProcessor) ReplaceWord(old, new string) {
	oldUp := strings.Title(old)
	newUp := strings.Title(new)
	t.Text = strings.ReplaceAll(t.Text, old, new)
	t.Text = strings.ReplaceAll(t.Text, oldUp, newUp)
}

// func main() {
// 	tp := TextProcessor{"Hello world, hello again!"}
// 	fmt.Println(tp.WordCount()) //map[hello:2, world:1, again:1]
// 	tp.ReplaceWord("hello", "hi")
// 	fmt.Println(tp.Text) // Hi world, hi again
// }

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Widht  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Widht
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Widht) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printShapeInfo(s Shape) {
	if _, ok := s.(Rectangle); ok {
		fmt.Printf("Прямоугольник -  Площадь: %.2f Периметр: %.2f\n", s.Area(), s.Perimeter())
	} else if _, ok := s.(Circle); ok {
		fmt.Printf("Круг - Площадь: %.2f Периметр: %.2f\n", s.Area(), s.Perimeter())
	}
}

// func main() {
// 	r := Rectangle{Widht: 5, Height: 10}
// 	c := Circle{Radius: 3}
// 	printShapeInfo(r)
// 	printShapeInfo(c)
// }

type Sorter interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (b BubbleSort) Sort(m []int) []int {
	if len(m) <= 1 {
		return m
	}
	n := len(m)
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if m[j] > m[j+1] {
				m[j], m[j+1] = m[j+1], m[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return m
}

type QuickSort struct{}

func (q QuickSort) Sort(m []int) []int {
	if len(m) <= 1 {
		return m
	}
	quickSort(m, 0, len(m)-1)
	return m
}

func quickSort(m []int, low, high int) {
	if low < high {
		pivotIndex := partition(m, low, high)
		quickSort(m, low, pivotIndex-1)
		quickSort(m, pivotIndex+1, high)
	}
}

func partition(m []int, low, high int) int {
	pivot := m[high]
	i := low
	for j := low; j < high; j++ {
		if m[j] <= pivot {
			m[i], m[j] = m[j], m[i]
			i++
		}
	}
	m[i], m[high] = m[high], m[i]
	return i
}

func sortNumbers(s Sorter, nums []int) []int {
	final := make([]int, len(nums))
	copy(final, nums)
	s.Sort(final)
	return final
}

func main() {
	b := BubbleSort{}
	q := QuickSort{}
	nums := []int{5, 2, 9, 1, 5, 6}
	fmt.Println(sortNumbers(b, nums))
	fmt.Println(sortNumbers(q, nums))
}
