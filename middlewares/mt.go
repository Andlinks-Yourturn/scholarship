package middlewares

import (
	"fmt"
	"encoding/hex"
	"net/http"
	"github.com/bitly/go-simplejson"

	"github.com/pkg/errors"
)

//string to HEX
func stringToHex(input string) string{
	result := hex.EncodeToString([]byte(input))
	return result

}
//将数字转化为16进制表示
func getlength(length int) string{

	//防止返回0
	if length == 1 {
		return "0x01"
	}else {

		result := fmt.Sprintf("%02x", length/2)
		return result
	}
}



func getRequest(inputKey string, inputValue string) string {
	//编码
	//string to HEX
	stringValue := stringToHex(inputValue)


	partM := getlength(len(inputKey))
	partL :=getlength(len(partM))

	fmt.Println("lenK: ",partL," ",partM)

	retultK :=partL+partM+inputKey


	partM = getlength(len(stringValue))
	partL =getlength(len(partM))


	fmt.Println("lenM: ",len(stringValue),partM," ",partL)
	retultV :=partL+partM+stringValue

	reault := "0x01"+retultK+retultV

	return reault


}


func InsertIntoMT(url string,inputKey string, inputValue string) error{

	request := url+"/broadcast_tx_commit?tx="+getRequest(inputKey,inputValue)

	fmt.Println("url, ",request)
	res, err := http.Get(request)
	if err !=nil{
		fmt.Println(err)
		return err
	}
	js, err := simplejson.NewFromReader(res.Body) //反序列化
	if err != nil {
		panic(err.Error())
	}

	info := js.Get("error").MustString()


	fmt.Println("data ", info)

	if  info == ""{
		return nil
	}else{
		return errors.New(info)
	}

}

//发送比较文件的请求
func getFileRequest(projectKey string, studentKey string) string {
	//编码
	//string to HEX
	//stringValue := stringToHex(inputValue)  tx=0x020102null0117QmPiGJvbS6Dq8BzV5wP2vQBrzHZDZsyARidEFkbv2noF9u
	partM := getlength(len(projectKey))
	partL :=getlength(len(partM))


	fmt.Println("lenK: ",partL," ",partM)

	retultK :=partL+partM+projectKey


	partM = getlength(len(studentKey))
	partL =getlength(len(partM))


	fmt.Println("lenM: ",partM," ",partL)
	retultV :=partL+partM+studentKey

	reault := "0x02"+retultK+retultV

	return reault


}

func Comparefiles(url string,projectKey string, studentKey string) string{

	request := url+"/broadcast_tx_commit?tx="+getFileRequest(projectKey,studentKey)

	fmt.Println("url, ",request)
	res, err := http.Get(request)
	js, err := simplejson.NewFromReader(res.Body) //反序列化
	if err != nil {
		panic(err.Error())
	}

	info := js.Get("result").Get("deliver_tx").Get("data").MustString()


	fmt.Println("compare result ", info)
	result, err := hex.DecodeString(info)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
	if string(result) == "Matched"{
		return "success"
	}else{
		return "failure"
	}

}



func searchValue(url string, inputKey string) string{

	//request := url+"/abci_query?data=0x"+inputKey+"&path=\"\"&prove=false"
	request :="http://localhost:46657/abci_query?data=0x44CBAE1AC3FC5B5BE5A6C4626147BDFE8BEDB837&prove=false"
	fmt.Println("url, ",request)
	res, err := http.Get(request)
	js, err := simplejson.NewFromReader(res.Body) //反序列化
	if err != nil {
		panic(err.Error())
	}

	info := js.Get("result").Get("response").Get("log").MustString()


	fmt.Println("log ", info)

	if info == "exists"{
		value := js.Get("result").Get("response").Get("value").MustString()
		decoded, err := hex.DecodeString(value)
		if err != nil {
			panic(err.Error())
		}

		return  string(decoded)
	}else{
		return "error"
	}
}
//func main() {
//
//	url := "http://localhost:46657"
//
//	//project
//	key := "44CBAE1AC3FC5B5BE5A6C4626147BDFE8BEDB837"
//
//	value := "Qmeza6kXi6kd3962Yb1biD2juqGNN62JrYwhra3PyeuyEs"
//
//	fmt.Println(insertIntoMT(url,key,value))
//
//	key = "44CBAE1AC3FC5B5BE5A6C4626147BDFE8BEDB836"
//
//	value = "QmWY5RMFpwPkdmAs6tDmUKdyXgE9BYEYzBqGwVk9o1DVcK"
//	fmt.Println(insertIntoMT(url,key,value))
//	fmt.Println(searchValue(url,key))
//}



