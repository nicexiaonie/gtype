package gtype

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"reflect"
	"strconv"
)

func StringToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//生成32位md5字串
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
		}
	}
	return r
}
