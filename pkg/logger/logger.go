package logger

import "fmt"

func LogInfo(message interface{}) {

}

func LogWarning(message interface{}) {

}

func LogDebug(message interface{}) {

}

func LogError(message interface{}) {
	fmt.Println(message)
}
