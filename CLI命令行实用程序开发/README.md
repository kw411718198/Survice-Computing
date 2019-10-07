# CLI 命令行实用程序开发

---

[TOC]


## 内容概述

----

CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，通过shell或script使得应用获得最大的灵活性与开发效率。Linux提供了cat、ls、copy等命令与操作系统交互；go语言提供一组实用程序完成从编码、编译、库管理、产品发布全过程支持；容器服务如docker、k8s提供了大量实用程序支撑云服务的开发、部署、监控、访问等管理任务；git、npm等都是大家比较熟悉的工具。尽管操作系统与应用系统服务可视化、图形化，但在开发领域，CLI在编程、调试、运维、管理中提供了图形化程序不可替代的灵活性与效率。

在[开发Linux命令行应用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)中，作者阐述了selpg的基本功能：该实用程序从标准输入或从作为命令行参数给出的文件名读取文本输入。它允许用户指定来自该输入并随后将被输出的页面范围。



## 关于原始selpg

----

### 命令行准则

​	通用Linux实用程序的编写应该遵守以下准则，这些准则经过长期发展，以利于确保用户更加灵活地与程序进行交互。

- 准则1	关于输入
  - 在命令行上指定文件名
  - 标准输入（stdin），缺省情况下为终端（也就是用户的键盘）
  - 使用 shell 操作符“<”（重定向标准输入），也可将标准输入重定向为来自文件
  - 用 shell 操作符“|”（pipe）也可以使标准输入来自另一个程序的标准输出
- 准则2       关于输入
  - 输出应该被写至标准输出，缺省情况下标准输出同样也是终端（也就是用户的屏幕）
  - 使用 shell 操作符“>”（重定向标准输出）可以将标准输出重定向至文件
  - 使用“|”操作符，command 的输出可以成为另一个程序的标准输入
- 准则3       错误输出
  - 错误输出应该被写至标准错误（stderr），缺省情况下标准错误同样也是终端（也就是用户的屏幕）
  - 使用标准错误重定向，也可以将错误重定向至文件，可以将标准输出和标准错误都重定向至不同的文件
  - 如果已将标准输出重定向至某一位置，也可以将标准错误重定向至同一位置。
- 准则4      执行
  - 不管程序的输入源（文件、管道或终端）和输出目的地是什么，程序都应该以同样的方式工作。
- 准则5     命令行参数
  - 如果指定了文件名参数，则程序把它作为输入。否则程序从标准输入进行读取。
  - 所有选项都应以“-”（连字符）开头。选项可以附加参数

### selpg程序逻辑

​	selpg 首先处理所有的命令行参数。在扫描了所有的选项参数（也就是那些以连字符为前缀的参数）后，如果 selpg 发现还有一个参数，则它会接受该参数为输入文件的名称并尝试打开它以进行读取。如果没有其它参数，则 selpg 假定输入来自标准输入。

1. selpg要求用户用两个命令行参数“-sNumber”和“-eNumber”指定要抽取的页面范围的起始页和结束页。selpg将对所给页号进行合理性检查，包括数字是否为有效正整数、结束页不小于起始页等。
2. 可选参数“-f”和"-lNumber"是互斥的，前者表示在输入中寻找换页符作为页定界符处理，后者表示固定每页的长度。
3. “-dDestination”选项将选定的页直接发送至打印机。
4. 对于页长度计数或者换页定界，selpg都有其相应的计数方法。

### selpg.c

​	在selpg.c中声明了一个结构，用来对用户从命令行输入的信息进行存储便于下一步操作,包括文档选择的起始页和结束页，所在的文件名，每页的长度和分页的类型等：

```c
   struct selpg_args

    {
        int start_page;
        int end_page;
        char in_filename[BUFSIZ];
        int page_len; /* default value, can be overriden by "-l number" on command line */
        int page_type; /* 'l' for lines-delimited, 'f' for form-feed-delimited */
                        /* default is 'l' */
        char print_dest[BUFSIZ];
    };
```

​	整个代码方法部分包括四个主要函数：

```c
    void usage(void);
    int main(int ac, char **av);
    void process_args(int ac, char **av, sp_args* psa);
    void process_input(sp_args sa);
```

​	其中main函数是结构体的变量声明以及初始化其他参数并调用其他两个函数；

​	process_args函数是进行进行用户输入的命令检测以及分析，至少需要输入3个以及以上参数，根据每个参数类型做出相应的判断执行；

​	process_input函数是进行输入输出源的检测，包括内容的读取以及输出的定位选择；

​	usage函数是进行提示帮助信息。

----

## selpg.go

​	下面根据以上对于函数的分析进行go语言的转化：

1. ​	结构体声明

   ```go
   type sp_args struct {
   	start_page  int
   	end_page    int
   	in_filename string
   	page_len    int
   	page_type   bool
   	print_dest  string
   }
   ```

   直接按照原c文件进行结构体声明转化即可。

2. 主函数

   ```go
   func main() {
   	var sa sp_args
   	init_args(&sa)
   	process_args(&sa)
   	process_input(&sa)
   }
   ```

   ​	对于原selpg.c中的main函数中初始化的部分这里合并成成了一个init_args函数，声明sp_args变量sa并且进行参数检验以及输出运行。

3. init_args函数

   ```go
   func init_args(sa *sp_args) {
   	flag.IntVarP(&(sa.start_page), "start_page", "s", -1, "Define startpage")
   	flag.IntVarP(&(sa.end_page), "end_page", "e", -1, "Define endpage")
   	flag.IntVarP(&(sa.page_len), "page_len", "l", 72, "Define page length")
   	flag.BoolVarP(&(sa.page_type), "page_type", "f", false, "Define page type")
   	flag.StringVarP(&(sa.print_dest), "print_dest", "d", "", "Define print_dest")
   
   	flag.Usage = func() {
   		fmt.Printf("Usage:\n\n")
   		fmt.Printf("\tselpg -s start_page -e end_page [-f (speciy how the page is sperated)| -l lines_per_page_default_72] [-d dest] [filename]\n\n")
   	}
   	flag.Parse()
   }
   ```

   ​	pflag包中的函数IntVarP/BoolVarP/StringVarP可以取出命令行参数名称的参数的值,该函数无返回值.获得flag参数后，要用pflag.Parse()函数才能把参数解析出来。因此在这里进行参数的提取，并定义Usage函数当进行输入信息提示。

4. process_args函数

   ```go
   if sa.page_len <= 0 {
   		fmt.Fprintf(os.Stderr, "\n[Error]The pageLen can't be less than 1 ! Please check your command!\n")
   		flag.Usage()
   		os.Exit(7)
   	} else {
   		fmt.Printf("\n[ArgsStart]\n")
   		fmt.Printf("startPage: %d\nendPage: %d\ninputFile: %s\npageLength: %d\npageType: %s\nprintDestation: %s\n[ArgsEnd]", sa.start_page, sa.end_page, sa.in_filename, sa.page_len, sa.page_type, sa.print_dest)
   	}
   ```

   ​	参数检查函数类似c文件，对于不符合规定的参数进行提示信息的输出以及程序进程的退出。包括参数不完整、起始页不正确、页行数发生冲突等问题的提示处理。

5. process_input函数

   ```go
   	for true {
   		var line string
   		var err error
   		if sa.page_type {
   			line, err = buf.ReadString('\f')
   			page_count++
   		} else {
   			line, err = buf.ReadString('\n')
   			line_count++
   			if line_count > sa.page_len {
   				page_count++
   				line_count = 1
   			}
   		}
   		if err == io.EOF {
   			break
   		}
   		if err != nil {
   			fmt.Fprintf(os.Stderr, "\n[Error]%s:", "Input pipe open\n", "file read in\n")
   			os.Exit(10)
   		}
   		if (page_count >= sa.start_page) && (page_count <= sa.end_page) {
   			var outputErr error
   			_, outputErr = fout.Write([]byte(line))
   			if outputErr != nil {
   				fmt.Fprintf(os.Stderr, "\n[Error]%s:", "pipe input")
   				os.Exit(11)
   			}
   
   			if outputErr != nil {
   				fmt.Fprintf(os.Stderr, "\n[Error]%s:", "Error happend when output the pages.")
   				os.Exit(12)
   			}
   		}
   	}
   ```

   ​	该函数作为数据输出函数，负责参数的辨别以及对应信息的打印。对于输入文件以及输出文件，如果未进行初始化则等待标准输入，对于管道出入同样需要判断并给出相应的错误信息，根据不同的推出提示位置进行判断，否则进行输出。

## 测试结果

​	根据“使用selpg”中的提示做了部分测试：

​	测试1：

​	![测试1](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%951.PNG)

​	测试2：

​	![测试2](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%952.PNG)

​	测试3：

​	![测试3](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%953.PNG)

​	测试4：

​	![测试4](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%954.PNG)

​	测试5：

​	![测试5](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%955.PNG)

​	测试6：

​	![测试6](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%956.PNG)

​	测试7：

​	![测试7](https://github.com/kw411718198/Survice-Computing/blob/master/CLI%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%AE%9E%E7%94%A8%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91/image3/%E6%B5%8B%E8%AF%957.PNG)
