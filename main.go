package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lei006/go-license/ecc_tool"
	"github.com/lei006/go-license/license"
)

func main() {
	test2()
}

func test2() {

	//licensefile := "./license.lic"

	//生成ECC密钥对文件
	pub_key, pri_key, err := ecc_tool.GenerateECCKeyString()
	if err != nil {
		fmt.Println("生成密钥对出错")
		return
	}

	for i := 0; i < 10000; i++ {

		licenser := license.MakeLicenser()
		lic_data := license.MakeLicenserData()

		lic_data.AppCode = RandomString(10)
		lic_data.PubKey = pub_key
		lic_data.HardSn = RandomString(10)
		lic_data.AppName = RandomString(10)
		lic_data.CompanyName = RandomString(32)
		lic_data.MaxNum = rand.Int63()
		lic_data.CopyRight = RandomString(32)
		lic_data.ExData["name"] = RandomString(6)

		lic_data_str := lic_data.ToString()
		//fmt.Println("lic_data_str --> ", lic_data_str)

		// 制做签名
		sign, err := licenser.MakeSign(lic_data_str, pri_key)
		if err != nil {
			fmt.Println("制做签名失败: ", err.Error())
		} else {

			lic_data.Sign = sign
			lic_data_str111, err_json := lic_data.ToJson()
			if err_json != nil {
				fmt.Println("lic_data.ToJson Error: ", err_json.Error())
				continue
			}

			_, err := licenser.VerifySign(lic_data_str111, pub_key)
			if err != nil {
				fmt.Println("验证验名失败:", err.Error())
			} else {
				fmt.Println("验证验名成功: ", i)
				fmt.Println(lic_data_str111)
				//fmt.Println(dec_data.InfoToString())
			}
		}
	}

	/*
		err := licenser.LoadFromFile(licensefile)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("ok")
		}
	*/

	time.Sleep(time.Duration(10) * time.Second)
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
