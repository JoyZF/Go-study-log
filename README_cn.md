# Less Is More
## 不支持三元
其他语言中：
```
    cond ? true_value : false_value
```
而在Golang 中并没有三元表达式。
只能写成
if cond {
return true_value
} else {
return false_value
}

官方的解释：
![image](https://mmbiz.qpic.cn/mmbiz_png/KVl0giak5ib4iaLoqX7Cp17TQxQoFTpy8O0gNQ6eO8iawGYcficfzEIfMTZeYexzAxicr0lSxTmx2dapOKjS0iashcmiaw/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

目的为了防止三元的滥用：
比如说
- 嵌套三元

## 廉价的goroutine
在Java中开一个线程需要实现Thread ，重写run方法，调用start方法来启动 或者 通过Runnable接口创建线程类。
为了解决频繁创建线程而产生过多的开销，还需要创建一个线程池。

而在Golang中实现协程只需要
```
    go func
```

两者对比起来 Golang显示并发的代码量太"廉价"了。这会带来很多问题：
- goroutine泄露
- OOM
- ...

## 别扭的OOP
![image](https://mmbiz.qpic.cn/mmbiz_png/KVl0giak5ib4grz5IlIjYicic3SYiaLsMnN6mXMc53XnzyJqVzibGX7uYWQ2fMzHqXb4hwBsgRB0k1oWdvr5oLicTBB8Q/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

### 封装
在Golang中属性访问的权限是通过首字母大小写来控制的
- 首字母大写，代表公共的、可被外部访问
- 首字母小写，代表私有的，不可以被外部访问
example:
```
type Animal struct {
 name string
}

func NewAnimal() *Animal {
 return &Animal{}
}

func (p *Animal) SetName(name string) {
 p.name = name
}

func (p *Animal) GetName() string {
 return p.name
}

```
### 继承
Go中没有extent 关键字，在继承上我们可以用组合的方式来实现
```
type Animal struct {
 Name string
}

type Cat struct {
 Animal
 FeatureA string
}

type Dog struct {
 Animal
 FeatureB string
}

```

### 多态
面向对象中的 “多态” 指的同一个行为具有多种不同表现形式或形态的能力，具体是指一个类实例（对象）的相同方法在不同情形有不同表现形式。
Go中多态是通过接口实现的
```
type AnimalSounder interface {
    MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder){
    animalSounder.MakeDNA()
}

```

在CAT或DOG中也实现了MakeSomeDNA 那么我们可以认为他是AnimalSounder接口类型的
```
type AnimalSounder interface {
 MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder) {
 animalSounder.MakeDNA()
}

func (c *Cat) MakeDNA() {
 fmt.Println("实现interface的方法,则Cat也是AnimalSounder类型")
}

func (c *Dog) MakeDNA() {
 fmt.Println("实现interface的方法,则Dog也是AnimalSounder类型")
}

func main() {
 MakeSomeDNA(&Cat{})
 MakeSomeDNA(&Dog{})
}
```


## 不支持范型😢
不过已经提上ISSUE



