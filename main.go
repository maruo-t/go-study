package main

// i5 := 500 // エラー
// s5 := "Golang" // エラー
var i5 = 500
var s5 = "Golang"

func outer() {
	// i6 := 600 // エラー
	// s6 := "Golang" // エラー
	var i6 = 600
	var s6 = "Golang"
	println(i6, s6)
}

func main() {
	println("Hello, World!")
	var i int = 100
	println(i)

	var s string = "Hello, World!"
	println(s)

	var t, f bool = true, false
	println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	println(i2, s2)

	var i3 int
	var s3 string
	println(i3, s3)

	i3 = 300
	s3 = "Golang"
	println(i3, s3)

	// 暗黙的な定義
	i4 := 400
	s4 := "Golang"
	println(i4, s4)

	println(i5, s5)

	outer()
}
