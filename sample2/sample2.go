

package greetings

import "fmt"

func GetHelloWorld(worldHint int) (string, error) {
	text1, err := getHello()
	if err == nil {
		text2, err := getWorld(worldHint)
		if err == nil {
			text1 = text1 + " " + text2
		}
	}
	return text1, err
}

func getHello() (string, error) {
	return "Hello", nil
}

func getWorld(worldHint int) (string, error) {
	if worldHint == 42 {
		return "World", nil
	}
	return "", fmt.Errorf("Could not find World")
}
