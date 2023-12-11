package license

import (
	"encoding/json"
	"fmt"
)

type LicenserData struct {
	AppName     string `bson:"appname" json:"appname"`         //程序名
	AppCode     string `bson:"appcode" json:"appcode"`         //程序名
	CompanyName string `bson:"companyname" json:"companyname"` //授权公司
	HardSn      string `bson:"hardsn" json:"hardsn"`           //授权硬件id
	MaxNum      int64  `bson:"maxnum" json:"maxnum"`           //最大数量
	ExpireAt    int64  `bson:"expire_at" json:"expire_at"`     //过期时间签
	Copyright   string `bson:"copyright" json:"copyright"`     //版权所有人
	Desc        string `bson:"desc" json:"desc"`               //描述
	Sign        string `bson:"sign" json:"sign"`               //签名字符串
	PubKey      string `bson:"pub_key" json:"pub_key"`         //公钥
}

func MakeLicenserData(appcode, hardsn string) *LicenserData {

	licenserData := &LicenserData{
		AppCode: appcode,
		HardSn:  hardsn,
	}
	return licenserData
}

// 生成验签的字符串
func (_data *LicenserData) ToString() string {

	text := fmt.Sprintf("%s%s%s%s-%d%d-%s%s%s",
		_data.AppName, _data.AppCode, _data.CompanyName, _data.HardSn,
		_data.MaxNum, _data.ExpireAt,
		_data.Copyright, _data.Desc, _data.PubKey)

	return text
}

func (_data *LicenserData) ToJson() (string, error) {
	b, err := json.Marshal(_data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (_data *LicenserData) FromJson(data_str string) error {
	data := []byte(data_str)
	err := json.Unmarshal(data, _data)
	return err
}

func (_data *LicenserData) InfoToString() string {
	out_str := ""
	out_str += "CompanyName:" + _data.CompanyName + "\r\n"
	out_str += "AppName    :" + _data.AppName + "\r\n"
	out_str += "AppCode    :" + _data.AppCode + "\r\n"
	out_str += "MaxNum     :" + fmt.Sprintf("%d", _data.MaxNum) + "\r\n"
	out_str += "ExpireAt   :" + UnixToTimeStr(_data.ExpireAt) + "\r\n"
	return out_str
}
