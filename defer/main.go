package main

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main()  {
	println(f1())//1 f1返回命名变量r，return 的时候r=0，然后执行defer对r++， 在调用方获取的时候就变成了 1
	println(f2())//5 f2返回命名变量r，return的时候已经把t的值5传递给了r，这时候r的值是5，然后执行defer语句改变的只是t的值，和r没关系，这里如果采用匿名返回，就会返回修改后的值
	println(f3())//1 f3返回命名变量r，return的时候返回1，然后执行defer语句，这里defer是采用传参的形式，defer后面的函数中修改和外面的r没关系(当然切片、map这些要除外)
}
