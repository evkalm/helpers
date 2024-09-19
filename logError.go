package helpers

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func LogError(params ...interface{}) {
	mess := strings.Builder{}

	timeNow := time.Now().Format("2006-01-02 15:04:05")
	mess.WriteString(colorName(timeNow))

	mess.WriteString(colorErr(" ERROR!"))

	var description string
	var details string
	for _, param := range params {
		switch v := param.(type) {
		case error:
			details += v.Error()

		case *error:
			details += (*v).Error()

		default:
			description += fmt.Sprintf("%v", v) + " "

		}
	}
	if len(description) > 0 {
		mess.WriteString(colorMapKey("\n  Description: "))
		mess.WriteString(description)
	}

	mess.WriteString(colorMapKey("\n  Detail: ") + details + "\n")
	mess.WriteString(colorMapKey("  Place: "))

	_, file, line, ok := runtime.Caller(1)
	if ok {
		mess.WriteString(fmt.Sprintf("%s:%d.\n\n", file, line))
	} else {
		mess.WriteString("undecided \n\n")
	}

	fmt.Print(mess.String())
}
