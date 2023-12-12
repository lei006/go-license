package license

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/lei006/go-license/ecc_tool"
)

// licenser 更新回调
type LicenserUpdateCallback func()

type Licenser struct {
	key string //公钥
}

func MakeLicenser() *Licenser {
	licenser := &Licenser{}
	return licenser
}

func (_lic *Licenser) MakeNewKey() (string, string, error) {

	pub_key, pri_key, err := ecc_tool.GenerateECCKeyString()
	if err != nil {
		return "", "", err
	}

	encode_pub_key := _lic.encodeToString(pub_key)

	encode_pri_key := _lic.encodeToString(pri_key)

	return encode_pub_key, encode_pri_key, nil
}

// 制做签名--base64
func (_lic *Licenser) MakeSign(data string, encode_pri_key string) (string, error) {

	pri_key, err := _lic.decodeFromString(encode_pri_key)
	if err != nil {
		//解码私钥出错
		return "", errors.New("VerifySign _lic.decodeFromString error:" + err.Error())
	}

	//用私钥签名
	sign_data, err := ecc_tool.Sign(data, pri_key)
	if err != nil {
		return "", errors.New("VerifySign ecc_tool.Sign error:" + err.Error())
	}

	//把签名，编码
	encode_sign_data := _lic.encodeToString(sign_data)
	return encode_sign_data, nil
}

// 效验签名
func (_lic *Licenser) VerifySign(data string, sign_base64 string, pub_key_base64 string) (bool, error) {

	// 解码公钥
	pub_key, err := _lic.decodeFromString(pub_key_base64)
	if err != nil {
		return false, errors.New("VerifySign decodeFromString error:" + err.Error())
	}

	// 解码签名
	sign, err := _lic.decodeFromString(sign_base64)
	if err != nil {
		return false, errors.New("VerifySign decodeFromString error:" + err.Error())
	}
	//验证签名
	bret, err := ecc_tool.VerifySign(data, sign, pub_key)
	if err != nil {
		return false, errors.New("VerifySign ecc_tool.VerifySign error:" + err.Error())
	}

	return bret, nil
}

func (_lic *Licenser) encodeToString(data string) string {
	encodedStr := base64.URLEncoding.EncodeToString([]byte(data))
	return encodedStr
}

func (_lic *Licenser) decodeFromString(data string) (string, error) {
	out_data, err := base64.URLEncoding.DecodeString(data)
	return string(out_data), err
}

type DataItem struct {
	Key string      `bson:"key" json:"key"`
	Val interface{} `bson:"val" json:"val"`
}

func UnixToTimeStr(t_data int64) string {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//时间戳转日期
	dataTimeStr := time.Unix(t_data, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr

}
