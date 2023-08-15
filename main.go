package main

func Plus(x, y int) int {
	return x + y
}

func Divide(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	println("No return")
	return
}

func returnFunc() func() {
	return func() {
		println("I'm a function")
	}
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	i := Plus(1, 2)
	println(i)

	i2, _ := Divide(9, 2)
	println(i2)

	i4 := Double(1000)
	println(i4)

	Noreturn()

	// 無名関数
	f := func(x, y int) int { return x + y }
	i5 := f(1, 2)
	println(i5)

	i6 := func(x, y int) int { return x + y }(1, 2)
	println(i6)

	// 関数を返す関数
	f2 := returnFunc()
	f2()

	// クロージャ
	counter1 := incrementGenerator()
	counter2 := incrementGenerator()
	println(counter1())
	println(counter1())
	println(counter2())
	println(counter2())
	println(counter2())
	println(counter1())

	// ジェネレーター
	gen := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
	}
	counter3 := gen()
	println(counter3())
	println(counter3())
	println(counter3())
	println(counter3())
	println(counter3())
}
