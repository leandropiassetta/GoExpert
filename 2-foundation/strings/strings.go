package main

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

func keyCache(names ...string) string {
	key := names[0]

	for _, name := range names[1:] {
		if name != "" {
			key += "-" + name
		}
	}
	hashF := sha1.New()
	hashF.Write([]byte(key))

	keyCache := fmt.Sprintf("%x\n", hashF.Sum(nil))

	return keyCache
}

func extractBodyText(logMessage string) string {
	splitParts := strings.SplitAfterN(logMessage, "body:", 2)

	bodyText := splitParts[1]

	return bodyText
}

// Hello
func workingWithStrings(word string) string {
	result := "eHlol"
	newWordBytes := []byte(word)

	for i, letter := range word {
		if strings.ToLower(string(letter)) != strings.ToLower(string(result[i])) {
			oldLetter := string(letter)
			newLetter := result[i]

			for _, resultLetter := range result {
				if strings.ToLower(string(resultLetter)) == strings.ToLower(oldLetter) {
					newWordBytes[i] = byte(newLetter)
					break
				}
			}
		}
	}

	return string(newWordBytes)
}

func cleanedMsg(msg string) string {
	cleanedMsg := strings.Replace(msg, "message:_msg_:", "", -1)

	return cleanedMsg
}

func main() {
	keyCache1 := keyCache("name", "age", "", "")
	keyCache2 := keyCache("name", "", "age", "")
	keyCache3 := keyCache("name", "age", "", "")
	keyCache4 := keyCache("name", "age", "", "")

	fmt.Println("keyCache: ", keyCache1)
	fmt.Println("keyCache: ", keyCache2)
	fmt.Println("keyCache: ", keyCache3)
	fmt.Println("keyCache: ", keyCache4)

	logMessage := `{"msg":"Table doesn't exist"};Error Code: internal_server_error;Status: 500;Cause: []time="2023-07-13T10:13:26-04:00" msg="saved in cache the message error MySQL dataCollection: Message: Could not sampling data field, status response: 422 Unprocessable Entity; body: {\"msg\":\"Table doesn't exist\"};Error Code: internal_server_error;Status: 500;Cause: []"`

	msgError := "msg := message:_msg_:database_not_open_error_code:_500_cause:"

	bodyText := extractBodyText(logMessage)

	fmt.Println("cleanedMsg: ", cleanedMsg(msgError))
	fmt.Println("extractBodyText: ", bodyText)
	fmt.Println("workingWithStrings: ", workingWithStrings("Hello"))
}
