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

		appName := RandomString(10)
		appCode := RandomString(10)
		hardsn := RandomString(32)
		CompanyName := RandomString(32)
		Copyright := RandomString(32)
		maxNum := rand.Int63()

		licenser := license.MakeLicenser()
		lic_data := license.MakeLicenserData(appCode, hardsn)

		lic_data.AppName = appName
		lic_data.CompanyName = CompanyName
		lic_data.MaxNum = maxNum
		lic_data.Copyright = Copyright

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

			dec_data, err := licenser.DecodeData(lic_data_str111, pub_key)
			if err != nil {
				fmt.Println("解码失败:", err.Error())
			} else {
				fmt.Println("解码成功: ", i)
				fmt.Println(dec_data.InfoToString())
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
