package main

import (
	"fmt"
	"test/util"
)

const UserNamePrefix = "user"

func main() {
	// 第一题
	fmt.Println("---------第一题-----------")
	a := [4]int{10, 40, 5, 280}
	b := [5]int{234, 5, 2, 148, 23}
	v := 42

	var c, d []int
	for i := 0; i < len(a); i++ {
		c = append(c, a[i])
	}

	for i := 0; i < len(b); i++ {
		d = append(d, b[i])
	}

	ok := test(c, d, v)
	if ok {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 第二题
	fmt.Println("---------第二题-----------")
	for i := 90000; i < 90010; i++ {
		fmt.Println("i=", i, " username=", CreateUserName(int64(i)))
	}

	// 第三题
	fmt.Println("---------第三题-----------")
	// 使用salt使别人知道即便知道算法，但是无法预测结果
	salt := 12367
	for i := 8000000; i < 8000010; i++ {
		fmt.Println("i=", i, "imporved_userName=", ImprovedCreateUserName(int64(i), salt))
	}

	// 第四题
	fmt.Println("---------第四题-----------")
	testUserName := "user1234"
	userDevice := "09309403-C35C-49FF-F518-3389985E4951"
	passWord := "shoustexu123"
	UserRegister(testUserName, userDevice)
	UserLogin(testUserName, passWord)

	// 第五题
	fmt.Println("---------第五题-----------")
	// 更换手机时，在原手机通过解绑设备接口解绑原设备
	UnBindAccountAndDevice()
	// 在新设备上通过原账户登录，绑定新手机
	BindAccountAndDevice()
}

// 第一题
func test(a []int, b []int, v int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i]+b[j] == v {
				return true
			}
		}
	}
	return false
}

// 第二题：创建用户名
func CreateUserName(a int64) string {
	randomNum := pseudo_encrypt(a)
	return fmt.Sprintf("user%v", randomNum)
}

// pseudo_encrypt 加密算法
func pseudo_encrypt(a int64) int {
	var l1, l2, r1, r2 int64
	l1 = (a >> 16) & 0xffff
	r1 = a & 0xffff

	for i := 0; i <= 3; i++ {
		temp := int64((((1366*r1 + 150889) % 714025) / 714025.0) * 32767)
		l2 = r1
		r2 = l1 ^ temp
		l1 = l2
		r1 = r2
	}
	return int((r1 << 16) + l1)
}

// 第三题：改进后的用户名生成算法
//4位数80进制对应最小10进制数据：512000
//4位数80进制对应最大10进制数据：40959999
func ImprovedCreateUserName(a int64, salt int) string {

	randomNum := pseudo_encrypt(a)
	randomNum += 512000 + salt
	str := util.DecimalToAny(randomNum, 80)
	return UserNamePrefix + str
}

// 第四题：用户注册登录
// userDevice传设备UUID，为安卓或者iOS设备的唯一设备标识
func UserRegister(userName string, userDevice string) {
	// 将userName、userDevice插入数据库
	fmt.Println("用户注册成功")
}

func UserLogin(userName string, password string) {
	// 1.用户登录时生成一个token，及一个auto_token，存入redis，过期时间设置30分钟
	// 2.半小时内，请求header只要带有有效token，即可通过鉴权，正常访问接口；
	// 3.一旦token失效，可通过auto_token通过接口换取有效token
	// 4.如果token及auto_token均失效，需要让用户重新登录
	fmt.Println("用户登录成功")
}

// 第五题：用户更换手机时绑定账户
// 账户与设备绑定
func BindAccountAndDevice() {
	fmt.Println("账户与设备绑定成功")
}

// 账户与设备解绑
func UnBindAccountAndDevice() {
	fmt.Println("账户解绑成功")
}
