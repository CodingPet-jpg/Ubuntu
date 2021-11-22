package main

func main() {
	s := "Hello 世界"
	println(s[:9]) //字符串截断也是字节截断所以底层的utf8编码会被破坏
}
