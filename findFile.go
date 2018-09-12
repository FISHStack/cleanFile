package main

import(
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

var jsFiles = make([]string, 10)
var htmlFiles = make([]string, 10)
var usedJsFiles = make([]string, 10)
var complete chan int = make(chan int)
var wg sync.WaitGroup

func  main(){
	readDir("E:\\workspace-sts-3.8.1.RELEASE\\demo\\bootdo","")
	fmt.Println(len(jsFiles))
	fmt.Println(len(htmlFiles))

	for k, v := range jsFiles {
		fmt.Println(k,v)
		go readJsInHtml(k,v);
		//for kh, vh := range htmlFiles {
		//	fmt.Println(len(jsFiles),k,len(htmlFiles),kh,vh)
		//	b,err := ioutil.ReadFile(vh)
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//	str := string(b)
		//	if strings.Contains(str,v){
		//		usedJsFiles = append(usedJsFiles,v)
		//		break;
		//	}
		//}
	}
	//<- complete
	wg.Wait()
	fmt.Println(len(jsFiles),len(usedJsFiles))
	//6760 1918

}

func readJsInHtml(jsIndex int,jsFile string){
	wg.Add(1)
	//complete <- 0
	for kh, vh := range htmlFiles {
		fmt.Println(len(jsFiles),jsIndex,len(htmlFiles),kh,vh)
		b,err := ioutil.ReadFile(vh)
		if err != nil {
			fmt.Println(err)
		}
		str := string(b)
		if strings.Contains(str,jsFile){
			usedJsFiles = append(usedJsFiles,jsFile)
			break;
		}
	}
	wg.Done()

}

func readDir(dirPath string, tab string){
	flist,e := ioutil.ReadDir(dirPath)
	if e != nil {
		fmt.Println("Read file error")
		return
	}
	
	for _,f := range flist {
		if f.IsDir(){
			//fmt.Println(tab,"+",dirPath+"/"+f.Name())

			readDir(dirPath+"/"+f.Name(),tab+"\t")
		}else{
			//fmt.Println(tab,",",dirPath+"/"+f.Name())

			//b,err := ioutil.ReadFile(dirPath+"/"+f.Name())
			//if err != nil {
			//	fmt.Println(err)
			//}
			//str := string(b)
			if strings.HasSuffix(f.Name(),".js") {
				//fmt.Println("have js file ",f.Name())
				jsFiles = append(jsFiles,f.Name())
			}
			if strings.HasSuffix(f.Name(),".html"){
				htmlFiles = append(htmlFiles,dirPath+"/"+f.Name())
			}

		}
	}

}
