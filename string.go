package gtype

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"reflect"
	"strconv"
	"time"
)

func StringToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

// 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

// 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func ToString(v interface{}) string {
	r := ""
	if v != nil {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			r = v.(string)
			break
		case reflect.Int64:
			r = strconv.FormatInt(v.(int64), 10)
			break
		case reflect.Struct:
			rb, _ := json.Marshal(v)
			r = string(rb)
			break
		case reflect.Int:
			r = strconv.Itoa(v.(int))
			break
		case reflect.Map:
			rb, _ := json.Marshal(v)
			r = string(rb)
			break
		default:
			jsons, _ := json.Marshal(v)
			r = string(jsons)
			break
		}
	}
	return r
}

func MapTimeToStringFormat(v map[string]interface{}, timeformat string) map[string]interface{} {
	if len(timeformat) < 1 {
		timeformat = "2006-01-02 15:04:05"
	}
	for yk, yv := range v {
		if v != nil {
			if reflect.TypeOf(yv).String() == "time.Time" {
				v[yk] = yv.(time.Time).Format(timeformat)
			}
		}
	}
	return v
}
