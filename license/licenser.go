package license

import (
	"errors"
	"time"

	"github.com/lei006/go-license/ecc_tool"
)

// licenser 更新回调
type LicenserUpdateCallback func()

type Licenser struct {
}

func MakeLicenser() *Licenser {
	licenser := &Licenser{}
	return licenser
}

func (_lic *Licenser) DecodeData(lic_data string, pub_key string) (*LicenserData, error) {

	license_data := &LicenserData{}

	// 1. 解码licsense数据
	err := license_data.FromJson(lic_data)
	if err != nil {
		return nil, errors.New("格式出错")
	}

	// 2. 检查数据...
	err = _lic.decodeData(license_data, pub_key)
	if err != nil {
		return nil, err
	}

	return license_data, nil
}

// 检查数据...不回调..
func (_lic *Licenser) decodeData(license_data *LicenserData, pub_key string) error {

	// 5. 验证签名
	text := license_data.ToString()
	ret := _lic.EccVerifySign(text, license_data.Sign, pub_key)
	if !ret {
		return errors.New("验签名失败")
	}

	return nil
}

// 制做签名
func (_lic *Licenser) MakeSign(data string, pri_key string) (string, error) {
	return ecc_tool.Sign(data, pri_key)
}

// 效验签名
func (_lic *Licenser) EccVerifySign(data string, sign string, pub_key string) bool {
	return ecc_tool.VerifySign(data, sign, pub_key)
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
