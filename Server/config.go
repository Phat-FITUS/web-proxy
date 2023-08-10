package Server

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Phat-FITUS/web-proxy/HTTP"
)

func readConfig(filename string) string{
	data, err:= os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	temp := string(data)
	return temp
}

func IsInService() bool {
	temp := HTTP.Mapify(readConfig("./Server/proxy.config"), "\n")
	_, exist := temp["time"]
	if exist {
		times := strings.Split(temp["time"], "-")
		start, err := strconv.Atoi(times[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(times[1])
		if err != nil {
			panic(err)
		}
		currentTime := time.Now()
		if (currentTime.Hour() < start || currentTime.Hour() > end){
			return false
		}
	}
	return true
}

func IsAcceptableHost(host string) bool{
	temp := HTTP.Mapify(readConfig("./Server/proxy.config"), "\n") 
	_, exist := temp["whitelisting"]
	if exist {
		whiteList := strings.Split(temp["whitelisting"], ",")
		for i:= 0; i < len(whiteList); i++ {
			if strings.Contains(host, whiteList[i]) {
				return true
			}
		}
		return false
	}
	return true
}