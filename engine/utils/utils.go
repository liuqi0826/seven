package utils

import (
	"fmt"
)

func Check(title string, err error) {
	if err != nil {
		fmt.Println(title, err)
	}
}
