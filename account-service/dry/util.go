package dry

import (
	"fmt"
	"os"
)

func LookupOrPanic(env string) string {
	res, ok := os.LookupEnv(env)
	if ok {
		return res
	}

	panic(fmt.Sprintf("[ENV] %s is not found", env))
}
