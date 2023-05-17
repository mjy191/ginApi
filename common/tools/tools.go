package tools

import (
	"crypto/sha1"
	"fmt"
	"ginApi/common/enum"
	"ginApi/common/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func GetEnumValue(i int, maps map[int]string) string {
	if value, ok := maps[i]; ok {
		return value
	} else {
		return "其它"
	}
}

func GetPage(total int64, lastPage *int64, page *int64, pageSize int64) {
	if pageSize == 0 {
		pageSize = 10
	}
	*lastPage = total / pageSize
	if total%pageSize != 0 {
		*lastPage++
	}
	if *page < 1 {
		*page = 1
	}
	if *page > *lastPage {
		*page = *lastPage
	}
}

func GetError(err error, r interface{}) {
	s := reflect.TypeOf(r)
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		panic(response.Response{
			Code: enum.CodeParamError,
			Msg:  err.Error(),
		})
	}
	for _, fieldError := range errs {
		field, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_msg"
		errTagText := field.Tag.Get(errTag)
		errText := field.Tag.Get("msg")
		if errTagText != "" {
			panic(response.Response{
				Code: enum.CodeParamError,
				Msg:  errTagText,
			})
		}
		if errText != "" {
			panic(response.Response{
				Code: enum.CodeParamError,
				Msg:  errText,
			})
		}
		panic(response.Response{
			Code: enum.CodeParamError,
			Msg:  fieldError.Field() + ":" + fieldError.Tag(),
		})
	}
}

func Sha1(str string) string {
	sha := sha1.New()
	sha.Write([]byte(str))
	x := sha.Sum(nil)
	return fmt.Sprintf("%x", x)
}

func GetBody(c *gin.Context) (body string) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	body = string(buf[0:n])
	return body
}

func RandString(lenNum int) string {
	var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	str := strings.Builder{}
	length := len(CHARS)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}
