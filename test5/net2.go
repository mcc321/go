package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start :=time.Now()
	ch :=make(chan string)
	for key,url := range os.Args[1:]{
		go fetch(url,ch,key)
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed \n",time.Since(start).Seconds())
}

func fetch(url string,ch chan<-string,key int){
	start :=time.Now()
	resp,err :=http.Get(url)
	if err != nil{
		ch<-fmt.Sprint(err)
		return
	}

	nbytes,err :=io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if err != nil{
		ch<-fmt.Sprintf("while reading %s : %v",url,err)
		return
	}
	secs := time.Since(start).Seconds()
	ch<-fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
	r ,err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err!=nil{
		fmt.Fprintf(os.Stderr,"fetch: reading %s : %v\n",url,err)
		os.Exit(2)
	}
	ioutil.WriteFile("mcc"+strconv.Itoa(key)+".txt",[]byte(r),0777)

}
