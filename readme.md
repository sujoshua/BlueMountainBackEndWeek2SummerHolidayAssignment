## 作业 

1.用chan写一个简单生产者消费者模型。代码提交至git仓库，注意做好每次作业的文件分层。

2.算法题：

	https://www.luogu.com.cn/problem/P2404
	
	https://www.luogu.com.cn/problem/P1443

3.一些基础题：

（1）一个结构体所占空间大小，与下面哪些相关：

			A 成员本身大小；B 成员对齐系数；C 系统字长

> 答：A、B、C

（2）下面这个结构体的成员怎么排布，占用内存最小？（假设系统64位）

```go
type A struct{
  byte1 byte
  a struct{}
  num1 int32
  str string
}
```

> byte: 1字节   struct{}: 0字节(不能放在struct最后)  int32：4字节 string：固定16字节
>
> ```go
> type A struct{
>  a struct{}
>  byte1 byte
>  num1 int32
>  str string
> } //反正a不能放在最后
> ```
>
> 

（3）Go字符串中，每个字符占用多少字节：

			A 1；B 3；C 1-4

> 答案：C

   (4)  Go的map使用的数据结构是：

			A B+树；B 开放寻址法的Hash表 ；C 拉链法Hash

> 答案：C

（5）空结构体的地址在任何时候都是zerobase？

		   A 是；B 不是

> 答案： B struct里面的空结构体就不是zerobase，而是结构体里面的某个位置

（6）空接口就是nil接口？

> 答案：是