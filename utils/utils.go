package utils

import (
	"fmt"
)

func ErrorHandle(title string, err error) {
	if err != nil {
		fmt.Println("Error <", title, ">:", err.Error())
	}
}
