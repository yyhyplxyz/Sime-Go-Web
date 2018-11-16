# README
## 任务描述
+ 一个图书馆管理系统
+ 做标准测试
+ 做性能测试

## 框架选择
使用了httprouter脚手架，因为本人以前曾打过flask web应用，httprouter这个脚手架封装的语法糖
与flask极像，蓝图，路由等都有对应的数据结构，编码web程序十分方便

## 标准测试
+ 使用curl测试根路径，输出welcome
![](https://ws3.sinaimg.cn/large/006tNbRwgy1fx9qq3gngaj30r00a8tgg.jpg)
+ curl只适用于简单的api访问，为了便于测试，我们采用`testing`库，来做标准测试，以下分别是对书籍展示和书籍查询功能的测试。复杂的逻辑可以在代码中描述，避免了人
手工多次操作
![](https://ws4.sinaimg.cn/large/006tNbRwgy1fx9qurqf0qj31ae0aijt9.jpg)
![](https://ws4.sinaimg.cn/large/006tNbRwgy1fx9qv8h46mj31ae0aita6.jpg)

## AB测试
+ 传统的A/B测试，是一种把各组变量随机分配到特定的单变量处理水平，把一个或多个测试组的表现与控制组相比较，进行测试的方式。
+ 在本例子中逻辑操作较简单，没有很明显的优化问题。因此我们考虑很明显的负优化问题
+ 利用go的`benchmark`例子，对以上的标准测试再做性能测试
+ 同时考虑普通版本，在多个语句中加`defer`的版本，和加入了`time.sleep()`函数的版本
+ 发现前两个版本性能几乎没有差异，最后一个版本性能很明显较差。
![](https://ws3.sinaimg.cn/large/006tNbRwly1fx9qxv4vxyj31e80hs78h.jpg)
![](https://ws3.sinaimg.cn/large/006tNbRwly1fx9r2hppqjj31e80qyaed.jpg)
![](https://ws2.sinaimg.cn/large/006tNbRwly1fx9r329j31j31e80ji78b.jpg)