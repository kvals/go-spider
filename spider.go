package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	//网址前缀
	urlHead = "http://www.xxxx.com"
	//图片保存位置
	imageSavePath string
	//网路图片类型
	imageType = "image/jpeg"
	//url匹配正则表达式
	regexpUrlStr = `'[a-zA-Z0-9_/\-\.]*\.htm'`
	//jpg格式image匹配
	regexpImageStr = `"http://[a-zA-Z0-9_/\-\.]*\.jpg"`
	//已访问网络地址map
	visitUrlMap map[string]string
	//图片下载数量
	downloadCount int
	//最大下载量
	maxCount int
)

func initData() {
	visitUrlMap = make(map[string]string)
}

//访问url地址,获取网页内容,解析网页头信息,为图片则保存,不是图片则继续获取url
func applyNetUrl(url string) {
	visitUrlMap[url] = "1"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//	fileType := res.Header.Get("Content-Type")
	//	log.Println("------------------------------------------ content-type " + fileType)
	//	ok := strings.EqualFold(fileType, imageType)
	//	log.Println("------------------------------------------ content-type is image : ", ok)
	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//	if ok {
	//		saveFileName := getFileNameFromUrl(url)
	//		writeToFile(bodyData, imageSavePath+saveFileName)
	//	} else {
	imgList, urlList := getUrlMatch(bodyData)

	if size := len(imgList); size > 1 {
		go downloadImage(imgList)
	}

	for _, v := range urlList {
		v = urlHead + v
		v = strings.Replace(v, "'", "", -1)
		//		log.Println("----------------------------------------- get html url : ", v)
		_, ok := visitUrlMap[v]
		if !ok {
			applyNetUrl(v)
		}
	}
	//	}
}

//匹配网页所有href地址,获取图片或url地址
func getUrlMatch(res []byte) (imageLists, urlLists []string) {
	//	log.Println("-----------------------------------------start check url & imageurl ")
	regUrl := regexp.MustCompile(regexpUrlStr)
	urlLists = regUrl.FindAllString(string(res), -1)
	//	for _, v := range urlList {
	//		urlLists = append(urlLists, string(v))
	//	}

	regImage := regexp.MustCompile(regexpImageStr)
	imageLists = regImage.FindAllString(string(res), -1)
	//	for _, v := range imageList {
	//		imageLists = append(imageLists, string(v))
	//	}
	//	log.Println("----------------------------------------- get url like : ", urlLists)
	//	log.Println("----------------------------------------- get imageurl like : ", imageLists)
	return
}

func downloadImage(imageUrl []string) {
	for _, v := range imageUrl {
		v = strings.Replace(v, "\"", "", -1)
		//		log.Println("----------------------------------------- get Image url : ", v)
		res, err := http.Get(v)
		if err != nil {
			break
		}
		fileType := res.Header.Get("Content-Type")
		ok := strings.EqualFold(fileType, imageType)
		bodyData, err := ioutil.ReadAll(res.Body)
		if ok {
			filePath := getFileNameFromUrl(v)
			//			log.Println("----------------------------------------- save file path : ", filePath)
			writeToFile(bodyData, filePath)
		}
	}

}
