package utils

import (
	"fmt"
	"runtime"
)

func DeleteEscape(data string) (result string) {
	switch runtime.GOOS {
	case "windows":
		for i := 0; i < (len(data) - 3); i++ {
			result += string([]byte(data)[i])
		}
	case "darwin", "linux":
		for i := 0; i < (len(data) - 1); i++ {
			result += string([]byte(data)[i])
		}
	}
	fmt.Print(result)
	return
}
