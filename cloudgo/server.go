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