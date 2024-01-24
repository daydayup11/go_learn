package main

import (
	"github.com/pkg/errors"
)

//func main() {
//	LoadDefault()
//	locale := "en"
//	userID := 123
//	user, err := GetUserByID(userID)
//	if err != nil {
//		myErr := NewMyError(err, GetUsersNotFound, "", locale)
//		errMsg := myErr.Message()
//		fmt.Printf("Error: %s\n", errMsg)
//		// 在这里可以根据需要处理错误，例如记录日志、返回特定错误码等
//	} else {
//		// 处理用户信息
//		fmt.Printf("User: %+v\n", user)
//	}
//}

// GetUserByID 模拟获取用户信息的函数
func GetUserByID(userID int) (*User, error) {
	// 这里模拟获取用户信息的逻辑，实际中需要根据您的业务实现
	// 假设用户不存在
	return nil, errors.New("用户不存在")
}

// User 用户结构体（根据您的实际情况进行定义）
type User struct {
	ID   int
	Name string
	// 其他字段...
}
