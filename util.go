package main

import (
	"log"
	"os"
)

type reqBean struct {
	fileType string //下载文件类型
}

func (bean *reqBean) setFileType(fileType string) {
	bean.fileType = fileType
}

func writeToFile(data []byte, filePath string) (err error) {
	//	log.Println("--------------------------------- write file : " + filePath)
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	file.WriteString(string(data))
	downloadCount = downloadCount + 1
	log.Println("---------------------------------  ", downloadCount, "  :  ", maxCount)
	if downloadCount >= maxCount {
		stop()
	}
	return
}

func getFileNameFromUrl(url string) (filePath string) {
	//	log.Println("--------------------------------- find last name form image url = ", url)
	//	startIndex := strings.LastIndex("//", url)
	//	log.Println("--------------------------------- get startIndex = ", startIndex)
	name := string([]byte(url)[len(url)-10:])
	filePath = imageSavePath + name
	//	log.Println("--------------------------------- get filename = " + filePath)
	return
}

func crateDir() {
	_, err := os.Stat(imageSavePath)
	if err != nil {
		log.Println("--------------------------------- 输入的目录不存在,开始创建")
		err := os.MkdirAll(imageSavePath, os.ModePerm)
		if err != nil {
			log.Println("--------------------------------- 创建失败,退出程序")
			stop()
		} else {
			log.Println("--------------------------------- 创建成功,开始下载")
		}
	}

}
