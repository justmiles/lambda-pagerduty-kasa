package smartplug

import "fmt"

func logIfErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
