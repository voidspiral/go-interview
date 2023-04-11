package concurrency

//1. 如下程序打印什么
//2. 如下可以编译吗
//3. 打印的结果是什么
func Q4() {
	defer func() { println("1") }()
	defer func() { println("2") }()
	defer func() { println("3") }()
	panic("错误")
}
