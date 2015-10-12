// You can edit this code!
// Click here and start typing.

package main

// http://stackoverflow.com/questions/4729736/go-array-initialization
// http://blog.golang.org/go-slices-usage-and-internals - difference between arrays and slices
func identityMat4() [16]float64 {
	return [...]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

func main() {
	// 终止这个循环,只打印 0 到 5
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
			println("--------------")
		}
		println(i)
	}
	// Reverse a
	var a [16]float64
	a = identityMat4()
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
		println(a[i], a[j])
	}
}
