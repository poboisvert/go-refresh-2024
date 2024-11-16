package main

func main() {
	var a [3]int

	a[0] = 11
	a[1] = 12
	a[2] = 13

	for index, el := range a {
		println(index, el)
	}

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	modifyMe(&a)

	for index, el := range a {
		println(index, el)
	}

	q := [...]int{1, 2, 3}
	print(q[0])

	r := [...]byte{1, 2, 3}
	print(r[0], r[1], r[2])

	n := [...]rune{9: 'a'}
	println(n[9])

	// Slice
	sl := n[3:10]
	println(sl[6], len(sl))

}

func modifyMe(arr *[3]int) {
	arr[1] = 100
}
