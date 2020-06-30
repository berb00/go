package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	_ "io"
	"log"
	"math"
	"encoding/binary"
	"time"
)

type numbers struct {
	Num1 string `json:"num1"`
	Num2 string `json:"num2"`
}

type numsResponseData struct {
	Add string `json:"add"`
	Mul string `json:"mul"`
	Sub string `json:"sub"`
	Div string `json:"div"`
}

func process(numsdata numbers) (numsResponseData) {
	
	var numsres numsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2
	// numsres.Mul = numsdata.Num1 * numsdata.Num2
	// numsres.Sub = numsdata.Num1 - numsdata.Num2
	// numsres.Div = numsdata.Num1 / numsdata.Num2

	return numsres
}

func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numsData numbers
	var numsResData numsResponseData

	decoder.Decode(&numsData)
	numsResData = process(numsData)

	// 写文件
	fileHandle(numsData)
	
	// 返回值
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(numsResData); err != nil {
        panic(err)
    }
}

/**
 检查错误
 */
 func checkErr(err error){
	if err!=nil{
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}

func fileHandle (numsdata numbers) {
	fileInfo,err:=os.Stat("position.txt")
	if err!=nil{
		if os.IsNotExist(err){
			os.Create("position.txt")
		}
	}
	fmt.Println(fileInfo)

	//可写方式打开文件
	file, err := os.OpenFile("position.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	checkErr(err)
	defer file.Close()

	//写字节到文件中
    t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf(numsdata.Num1)
	fmt.Printf(numsdata.Num2)
	// bytesWritten, err := file.Write(Float64ToByte(numsdata.Num1))
	byteSlice := []byte(t + "    " + numsdata.Num1 + "," + numsdata.Num2 + "\n")
	bytesWritten, err := file.Write(byteSlice)

	checkErr(err)
	log.Printf("wrote %d bytes.\n", bytesWritten)

}

func Float64ToByte(float float64) []byte {
    bits := math.Float64bits(float)
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, bits)
 
    return bytes
}