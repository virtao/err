package err

import (
	"errors"
	"bytes"
	"strings"
	"strconv"
)

var (
	separator string = "\n"
)

func JoinStringError(args ...string) error {
	var bs bytes.Buffer
	for _, s := range args {
		bs.WriteString(s)
		bs.WriteString(separator)
	}
	return errors.New(bs.String())
}

func JoinError(args ...error) error {
	var bs bytes.Buffer
	for _, e := range args {
		bs.WriteString(e.Error())
		bs.WriteString(separator)
	}
	return errors.New(bs.String())
}

func JoinVarError(args ...interface{}) error {
	var bs bytes.Buffer
	for _, v := range args {
		switch v.(type) {
		case string:
			bs.WriteString(v.(string))
		case error:
			bs.WriteString(v.(error).Error())
		case int:
			bs.WriteString(strconv.Itoa(v.(int)))
		case float32:
			bs.WriteString(strconv.FormatFloat((float64(v.(float32))), 'E', -1, 64))
		case float64:
			bs.WriteString(strconv.FormatFloat((v.(float64)), 'E', -1, 64))
		default:
			bs.WriteString("Join Errors error, unknown type.")
		}
		bs.WriteString(separator)
	}
	return errors.New(bs.String())
}

func SetErrorSeparator(s string) {
	separator = s
}

func GetLastError(err error) string {
	return strings.SplitAfterN(err.Error(), separator, 2)[0]
}
