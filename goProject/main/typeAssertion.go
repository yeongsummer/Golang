package main

func main() {

	// 인터페이스
	var a interface{} = 1

	i := a       // 포인터 주소
	j := a.(int) // 포인터 값

	println(i)
	println(j)

	var k int = 10

	p := &k
	q := k

	println(p) // 포인터 주소
	println(q) // 포인터 값
}
