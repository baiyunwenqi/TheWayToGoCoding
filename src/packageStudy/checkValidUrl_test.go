package packageStudy

import (
	"testing"
	"log"
)

func TestCheckValidUrl(t *testing.T) {
	urls:=[]string{"http://10.","http://10.103.20.34:23","https://www.cnblogs.com/long_/archive/2010/09/13/1825239.html","http://103.20.34:23"}
	for _,url:=range urls{
		if CheckValidUrl(url){
			log.Println("valid "+url)
		}else{
			log.Println("Invalid "+url)
		}
	}

}
