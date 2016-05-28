package main

import (
	"fmt"
	"log"
)

var ch chan int

func main() {
	ch = make(chan int)
	url := "http://www.xxxx.com"
	initDownloadInfo()
	go start(url)
	<-ch
	log.Println("下载完成")
}

func start(url string) {
	initData()
	applyNetUrl(url)
}

func stop() {
	//	panic(0)
	ch <- 1
}

func initDownloadInfo() {
	log.Println("***************************************************")
	log.Println("***********输入保存位置")
	log.Println("***********输入下载数量")
	log.Println("***********格式如下: ")
	log.Println("***********/home/wtf/Desktop/output/")
	log.Println("***********1000")
	log.Println("***********请输入:")
	log.Println("***************************************************")
	fmt.Scanln(&imageSavePath)
	fmt.Scanln(&maxCount)
	log.Println("***********下载:", maxCount, "张到", imageSavePath, "目录下")
	if len(imageSavePath) <= 1 {
		log.Println("请输入正确保存位置")
		//		panic(0)
	}
	if maxCount <= 1 {
		log.Println("输入数量太少,默认为1000")
		maxCount = 100
	}
	crateDir()
}
