package common

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"regexp"
	"time"
)

func Md5Encode(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

// 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
func IsValidateUsername(username string) error {
	if len(username) == 0 {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

// 验证是否是合法的邮箱
func IsValidateEmail(email string) (err error) {
	if len(email) == 0 {
		err = errors.New("邮箱格式不符合规范")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

// 验证是否是合法的邮箱
func IsValidateMobile(mobile string) (err error) {
	if len(mobile) == 0 {
		err = errors.New("手机号格式不符合规范")
		return
	}
	pattern := `^1+\d{10}$`
	matched, _ := regexp.MatchString(pattern, mobile)

	if !matched {
		err = errors.New("手机号格式不符合规范")
	}
	return
}

// 是否是合法的密码
func IsValidatePassword(password, rePassword string) error {
	if len(password) == 0 {
		return errors.New("请输入密码")
	}
	if len(password) < 6 {
		return errors.New("密码过于简单")
	}
	return nil
}

//生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
