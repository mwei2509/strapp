package app

import (
	"errors"
	"fmt"
)

func Do(directory string, appTypeFlags []string) error {
	fmt.Println("create called", appTypeFlags)
	fmt.Println(appTypeFlags)
	return errors.New("BAD")
}
