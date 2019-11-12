## 概述

开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

## 任务目标

  1. 熟悉 go 服务器工作原理 
  2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。 
  3. 使用 curl 工具访问 web程序
  4. 对 web 执行压力测试

## 相关知识

[课件](http://blog.csdn.net/pmlpml/article/details/78404838)

## 基本要求

  1. 编程 web 服务程序 类似 cloudgo 应用。 
   -  要求有详细的注释 
   -  是否使用框架、选哪个框架自己决定，请在 README.md，说明你决策的依据。
  2. 使用 curl 测试，将测试结果写入 README.md 。
  3. 使用 ab 测试，将测试结果写入README.md。并解释重要参数。

 ## 实现过程
  1. 框架选择
  
  对于go的框架选择，我还是处于比较迷惑的状态，因此了解了一下go的六种主要框架类型的应用和差异比较，综合起来觉得beego似乎更适合自己，因此选择了使用beego框架进行开发。
        		beego是提供给go语言进行网站后台的web后台框架，采用了经典的传统MVC架构设计模式。
        		beego框架环境搭建：
        		执行指令下载beego的框架源码：
```bash
	go get -u -v github.com/astaxie/beego
```
执行指令下载beego的框架工具包：

```bash
	go get -u -v github.com/beego/bee
```
## 内部函数
 使用beego的框架函数来实现：

```go
	 package main
	
	import(
		"flag"
		"github.com/astaxie/beego"//use beego framework
	)
	
	type MainController struct{
		beego.Controller//beego controller
	}
	
	func(this *MainController) Get(){
		name := this.Ctx.Input.Param(":name")//get router information
		this.Ctx.WriteString("Welcome to this page," + name +" !")//write
	}
	
	func main(){
		port := flag.String("port","","port:default is 8080")//input port number
		flag.Parse()//parse
		beego.Router("cloudgo/:name",&MainController{})//set router
		beego.Run(":"+*port)//run
	}
```

 		 

 ## 程序运行
 - 在终端运行：

```bash
	 go run server.go -port 9090
```
- 网页输出
  ![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112101434337.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2t3NDExNzE4MTk4,size_16,color_FFFFFF,t_70)
- curl测试
  首先启动程序server.go，保持对端口处于监听状态。
  之后使用指令：
```bash
	curl -v http://local:9090/cloudgo/kw
```
网页的html结果如下，css与html相互嵌套：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112102923501.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2t3NDExNzE4MTk4,size_16,color_FFFFFF,t_70)
- ab测试
  首先使用以下指令进行ab测试的工具安装：
```bash
	sudo apt install apache2-utils
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112103655793.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2t3NDExNzE4MTk4,size_16,color_FFFFFF,t_70)
	接着启动程序server.go，按照ab测试的指令：
	ab -n 全部请求数 -c 并发数 测试url
	

```bash
	ab -n 9090 -c 808 https://localhost:9090/cloudgo/kw
```
808并发的情况下共测试访问localhost:9090 9090次，测试结果如下:
	![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112103737203.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2t3NDExNzE4MTk4,size_16,color_FFFFFF,t_70)
	重要参数解释如下：  
	Server Hostname:服务器的主页名为localhost  
	Server Port:服务器的端口号为9090  
	Document Path:服务器所在的文件路径  
	Document Length: 请求文档大小为25bytes  
	Concurrency Level:并发数为808  
	Time taken for tests:全部请求完成耗时2.454s  
	Complete requests: 全部请求数为9090  
	Failed requests: 失败的请求为0  
	Total transferred:总传输大小为1290780bytes  
	HTML transferred:整个场景中的HTML内容传输量  
	Requests per second:每秒请求数(平均）  
	Time per request: 每次并发请求时间(所有并发)   
	Time per request:每一请求时间(并发平均)  
	Transfer rate: 传输速率
