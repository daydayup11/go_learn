package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type MyError struct {
	err  error
	msg  string
	code int
	i18n string
}

// Error 实现 error 接口的方法
func (e *MyError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}
func (c *MyError) Wrap(err error, message string) *MyError {
	c.err = errors.Wrap(err, message)
	return c
}
func (c *MyError) Code() int {
	return c.code
}
func (c *MyError) WithMessage(message string) *MyError {
	c.msg = message + c.msg
	return c
}

func (c *MyError) Message() string {
	if c.msg == "" {
		// 使用 Bundle 获取本地化的消息
		return GetErrorMessage(c.i18n, c.code)
	}
	return c.msg
}

// NewMyError 构造 MyError 实例
func NewMyError(err error, code int, messages, locale string) *MyError {
	if messages == "" {
		messages = GetErrorMessage(locale, code)
	}
	return &MyError{
		code: code,
		msg:  messages,
		err:  err,
		i18n: locale,
	}
}
