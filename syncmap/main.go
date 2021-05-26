package main

import "fmt"

func main()  {
	s := `
		结论：
			sync.map 写入速度慢、查找、删除速度快.在读多写少的使用场景建议使用sync.map
	sync.Map 类型的底层数据结构如下：
	type Map struct {
		 mu Mutex //互斥锁 用于保护read 和 dirty
		 read atomic.Value // readOnly 只读数据 支持并发读取 atomic.Value 
		 dirty map[interface{}]*entry //原生map 线程不安全
		 misses int //统计有多少次读取read没有命中，每次失败之后misses累加1
		}
		
		// Map.read 属性实际存储的是 readOnly。
		type readOnly struct {
		 m       map[interface{}]*entry
		 amended bool //用于标记read 和dirty的数据是否一致
	}
	
	type entry struct {
 		p unsafe.Pointer // *interface{}
	}
	
	当我们从sync.map 类型中读取数据时,其先会在read中是否包含所需要的元素
	若有，则通过 atomic 原子操作读取数据并返回。
若无，则会判断 read.readOnly 中的 amended 属性，他会告诉程序 dirty 是否包含 read.readOnly.m 中没有的数据；因此若存在，也就是 amended 为 true，将会进一步到 dirty 中查找数据。
	
	写入过程：
	如果发现read中存在该元素，但已经被标记为删除，则说明dirty不等于nil（dirty肯定不包含该元素），则执行如下操作
		将元素从已删除改为nil
		将元素插入dirty
	如果发现read中不存在该元素，但dirty中存在该元素，则直接希尔更新entry的指向
	如果read和dirty都不存在该元素，则从read中付总未被标记删除的数据并想dirty中插入该元素，赋值元素值entry的指向。
	总结：
	查read、read上没有，或者已标记删除状态
	上互斥锁
	操作dirty，根据数据情况和状态处理
	
	删除流程：
	delete是将entry.p标记为删除状态，而不是真正删除。
	`
	fmt.Println(s)
}