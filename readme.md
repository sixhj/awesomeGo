
# 参考

go web编程 

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md

go 相关电子书 

https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books-zh.md#golang

---


go 是一门编译型语言，静态类型语言

将源代码和依赖一起打包，进行静态编译生成本地指令

go的代码是通过package 来组织的

go不支持继承

go的并发有两种 协程、channel,使用消息传递来共享内存。

```
& 仅用于生成其操作数对应的地址，也就是用于生成指针会出现在两个内容上：
一个是类型，* Type 这样的格式代表了一个指针类型 
一个是指针，* Pointer 这样的格式用于获取指针所对应的基本值
```

---

bee执行流程

```
1、main文件监听8080 
2、用户新连接 
3、初始化context 判断类型，WebSocket？
4、执行过滤器 判断是否设置过滤器 ， 有 =》 52 、 没有 =》 51 
51、静态文件处理 
52、执行过滤器AfterStatic 
53、查找固定路由、正则、自动等匹配 
6、执行controller逻辑 
7、过滤器before exec 
8、执行controller init 、 enable XSRF 是否启动跨域 一般都是 beego.Controller 的初始化，不建议用户继承的时候修改该函数 
9、执行controller prepare 一般用户参数初始化 
10、执行controller get/post 
11、执行controller finish 预留给用户用来重写的，用于释放一些资源。释放在 Init 中初始化的信息数据。 
12、执行过滤器 after exec 13、执行controller destructor

是否开启admin 监控统计url

```