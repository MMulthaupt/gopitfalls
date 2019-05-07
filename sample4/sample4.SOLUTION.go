//+build ignore

package retrospectively

import "fmt"

func DoSomething() (err error) {
	a, b := 1, 2
	defer func() {
		if a != b {
			err = fmt.Errorf("%d != %d", a, b)
		}
	}() // Defer does not defer a function, but a function-CALL. The parameters are evaluated immediately.
	a, b = 3, 3
	return err
}
