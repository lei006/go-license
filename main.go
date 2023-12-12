package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lei006/go-license/license"
)

func main() {

	for i := 0; i < 100; i++ {
		test2(100)
	}
}

func test2(loop_num int) {

	//licensefile := "./license.lic"
	licenser := license.MakeLicenser()
	//lic_struct := license.MakeLicenseStruct()

	//生成ECC密钥对文件
	pub_key, pri_key, err := licenser.MakeNewKey()
	if err != nil {
		fmt.Println("生成密钥对出错")
		return
	}

	for i := 0; i < loop_num; i++ {

		lic_data_str, err := license.GetRandomLicenseTestData()
		if err != nil {
			fmt.Println("取得随机数据失败: ", err.Error())
			continue
		}

		// 制做签名
		lic_data_sign, err := licenser.MakeSign(lic_data_str, pri_key)
		if err != nil {
			fmt.Println("制做签名失败: ", err.Error())
			continue
		}

		out_string := ""
		{
			// 结构体--帮忙生成交给外部的数据
			lic_struct := license.MakeLicenseStruct()
			lic_struct.Data = lic_data_str
			lic_struct.Sign = lic_data_sign
			lic_struct.PubKey = pub_key

			out_string, err = lic_struct.ToJson()
			if err != nil {
				fmt.Println("生成输出字符串失败: ", err.Error())
				continue
			}

		}
		//fmt.Println("生成的输出字符串为: ")
		//fmt.Println(out_string)

		/////////////////////////////////////////////////////////////////
		in_lic_struct := license.MakeLicenseStruct()
		{
			in_string := out_string
			// 结构体--帮忙生成结构体
			err = in_lic_struct.FromJson(in_string)
			if err != nil {
				fmt.Println("生成输出字符串失败: ", err.Error())
				continue
			}
		}

		bRet, err := licenser.VerifySign(in_lic_struct.Data, in_lic_struct.Sign, in_lic_struct.PubKey)
		if err != nil {
			fmt.Println("验证验名失败:", err.Error())
			continue
		}
		if bRet != true {
			fmt.Println("验证验名失败，但无错误输出")
			continue
		}

		fmt.Println("验证成功: ", i)

	}

	/*
		err := licenser.LoadFromFile(licensefile)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("ok")
		}
	*/

	//time.Sleep(time.Duration(10) * time.Second)
}

func RandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
