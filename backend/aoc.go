package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type crackerInput struct {
	Secretkey            string
	RequiredLeadingZeros int
}

type crackerOutput struct {
	AnswerKey    string
	AnswerMD5sum string
}

// Copied from https://stackoverflow.com/questions/2377881/how-to-get-a-md5-hash-from-a-string-in-golang
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Solution for Advent of code 2015/day4
func BruteForceMD5(secretkey string, nrOfLeadingZeros int) (string, string) {
	var cnt int = 0
	var curMD5 string = ""
	var expectedPrefix string = ""

	for i := 0; i < nrOfLeadingZeros; i++ {
		expectedPrefix += "0"
	}

	for {
		curMD5 = GetMD5Hash(secretkey + strconv.Itoa(cnt))
		if strings.HasPrefix(curMD5, expectedPrefix) {
			fmt.Println(curMD5)
			return strconv.Itoa(cnt), curMD5
		}
		cnt++
	}
}

func LogRequestToFile(filename string, logmessage string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	file.WriteString(logmessage)
	file.Sync()
	file.Close()
}

// Thanks to https://go.dev/doc/tutorial/web-service-gin
func BruteForceMD5FromGin(c *gin.Context) {
	var input crackerInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println("bindjson failed")
		return
	}

	key, mdsum := BruteForceMD5(input.Secretkey, input.RequiredLeadingZeros)

	var result = crackerOutput{AnswerKey: key, AnswerMD5sum: mdsum}

	elemIn, _ := json.Marshal(input)
	elemOut, _ := json.Marshal(result)

	var logtext = string(elemIn) + ",\n" + string(elemOut) + ",\n"

	LogRequestToFile("logfile.txt", logtext)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.POST("/mine", BruteForceMD5FromGin)
	router.Run("localhost:2512")
	// var secretkey = "abcdef"
	// fmt.Println(BruteForceMD5(secretkey, 5))
	// fmt.Println(BruteForceMD5(secretkey, 6))
}
