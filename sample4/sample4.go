

package retrospectively

import "fmt"

func DoSomething() (err error) {
	a, b := 1, 2
	defer func(x int) {
		if x != b {
			err = fmt.Errorf("%d != %d", x, b)
		}
	}(a)
	a, b = 3, 3
	return err
}
