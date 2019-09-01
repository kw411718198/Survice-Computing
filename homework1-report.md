#作业报告1


院系：数据科学与计算机学院
年纪专业：2017级软件工程
姓名：孔伟
学号：17343055
日期：2019年9月1日

------------

##【实验题目】
- 让你的计算机也能提供云桌面服务
##【实验目的】
1. 初步了解虚拟化技术，理解云计算的相关概念
2. 理解系统工程师面临的困境
3. 理解自动化安装、管理（DevOps）在云应用中的重要性
##【实验环境与要求】
- 用户通过互联网，使用微软远程桌面，远程访问你在PC机上创建的虚拟机
- 虚拟机操作系统 Centos，Ubuntu，或 你喜欢的 Linux 发行版，能使用 NAT 访问外网。
##【实验内容】
1. 安装 VirtualBox
安装 Git 客户端（git bash），下载地址：官网，或 gitfor windows 或 github desktop
安装 Oracle VirtualBox 5.X，官方下载
配置 VirtualBox 虚拟机存储位置，避免找不到虚拟机存储位置，特别是消耗启动盘的宝贵空间
VirtualBox菜单 ：管理 -> 全局设定，常规页面
创建虚拟机内部虚拟网络，使得 Vbox 内部虚拟机可以通过它，实现虚拟机之间、虚拟机与主机的通讯
VirtualBox菜单 ：管理 -> 主机网络管理器，创建一块虚拟网卡，网址分配：192.168.100.1/24
在主机 windows 命令行窗口输入 ipconfig 就可以看到 VirtualBox Host-Only Network #?: 的网卡
创建Linux虚拟机（以 CentoOS 为案例）
下载 Linux 发行版镜像。
如果是 Centos，仅需要 Minimal ISO；如果是 Ubuntu 请下载桌面和服务器
阿里云OPSX 下载
用 VBox 创建虚拟机。 虚拟机名称建议以 centos-xxx 或 ub-xxx 命名，如果向导不能创建 64 bit 虚拟机，请更换电脑!!!
建议虚拟机CPU、内存采用默认。如果是桌面版，CPU建议数1-2，内存不低于2G
显示，显存采用默认。如果是桌面版，显存越大越好
存储，不低于30G。避免以后扩展难。
网络，第一块网卡必须是 NAT；第二块网卡连接方式： Host-Only，接口就是前面创建的虚拟网卡
安装 Base 虚拟机，例如 centos-base。 利用虚拟化软件提供的虚拟机复制功能，避免每次安装 OS 系统的痛苦
按提示安装，直到完成
升级 OS 系统内核
获取 wget, yum install wget
配置源 163源、阿里云源
升级 OS内核， yum update
检查网卡配置
配置网络的UI界面 nmtui，配置第二块网卡地址
ping 主机，例如： ping 192.168.100.1
退出并关闭虚拟机
安装虚拟机
点击 centos-base 选择复制，输入新虚拟机的名，注意必须 选择重新初始化所有网卡的 MAC 地址
然后选 链接复制
配置主机名和第二块网卡
使用 nmtui 修改主机名和第二块网卡IP地址
重启
在主机上，应能 ping 到这个地址，且能通过 ssh 访问该虚拟机（windows 需启动 git bash）
如果你使用 vim 或 emacs
安装 vim 或 emacs
安装 C++ 开发工具
如果你使用 centos 桌面
重新配置虚拟机 CPU，内存，显存
启动虚拟机
安装桌面 yum groupinstall "GNOME Desktop"
设置启动目标为桌面 ln -sf /lib/systemd/system/runlevel5.target /etc/systemd/system/default.target
重启
安装 VirtualBox 增强功能
VirtualBox虚拟机CentOS安装增强功能Guest Additions
How to Install Kernel Headers in CentOS 7
安装 Chrome 浏览器
CentOS7 使用 yum 安装 chrome
配置用远程桌面访问你的虚拟机
参考：如何设置VirtualBox虚拟机远程访问模式
虚拟机无界面启动，用户即可通过网络，使用RDP客户端访问
以上一些操作内容仅适用宿主（hosted）为 window 10 环境，安装 CentOS 7 的操作。
### 【实验总结】
每个实验项目的报告必需写一段，文字不少于300字，可以写心得体会、问题讨论与思考、新的设想、感言总结或提出建议等等。



文件命名格式：
<学号>+<姓名>+<实验项目号>+<版本号>.rar

