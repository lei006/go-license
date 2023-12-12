package license

import (
	"encoding/json"
)

type LicenseData interface {
	ToJson() (string, error)
	FromJson(string) error
}

type LicenseDataEx struct {
	AppName     string                 `bson:"appname" json:"appname"`         //程序名
	AppCode     string                 `bson:"appcode" json:"appcode"`         //程序名
	CompanyName string                 `bson:"companyname" json:"companyname"` //接受授权公司
	HardSn      string                 `bson:"hardsn" json:"hardsn"`           //授权硬件id
	MaxNum      int64                  `bson:"maxnum" json:"maxnum"`           //最大数量
	MaxNum0     int64                  `bson:"maxnum0" json:"maxnum0"`         //最大数量
	MaxNum1     int64                  `bson:"maxnum1" json:"maxnum1"`         //最大数量
	ExpireAt    int64                  `bson:"expire_at" json:"expire_at"`     //过期时间--毫秒
	Desc        string                 `bson:"desc" json:"desc"`               //描述
	ExData      map[string]interface{} `bson:"ex_data" json:"ex_data"`         //其它数据
	Authorizer  string                 `bson:"authorizer" json:"authorizer"`   //授权人--授予他人权力或许可的人或组织
	CopyRight   string                 `bson:"copyright" json:"copyright"`     //版权所有人--版权复制
	Maker       string                 `bson:"maker" json:"maker"`             //制做者
}

func MakeLicenseData() *LicenseDataEx {
	licenserData := &LicenseDataEx{
		ExData: map[string]interface{}{},
	}
	return licenserData
}

func (_data *LicenseDataEx) ToJson() (string, error) {
	b, err := json.Marshal(_data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (_data *LicenseDataEx) FromJson(data_str string) error {
	data := []byte(data_str)
	err := json.Unmarshal(data, _data)
	return err
}
