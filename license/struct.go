package license

import (
	"encoding/json"
	"math/rand"
	"time"
)

// 用来生成格式数据
type LicenseStruct struct {
	Data   string `bson:"data" json:"data"`       //数据
	Sign   string `bson:"sign" json:"sign"`       //签名字符串
	PubKey string `bson:"pub_key" json:"pub_key"` //公钥
}

func MakeLicenseStruct() *LicenseStruct {
	str := &LicenseStruct{}
	return str
}

func (_data *LicenseStruct) ToJson() (string, error) {
	b, err := json.Marshal(_data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (_data *LicenseStruct) FromJson(data_str string) error {
	data := []byte(data_str)
	err := json.Unmarshal(data, _data)
	return err
}

func GetRandomLicenseTestData() (string, error) {

	lic_data := MakeLicenseData()

	lic_data.AppCode = RandomString(10)
	lic_data.HardSn = RandomString(10)
	lic_data.AppName = RandomString(10)
	lic_data.CompanyName = RandomString(32)
	lic_data.MaxNum = rand.Int63()
	lic_data.CopyRight = RandomString(32)
	lic_data.ExData["name"] = RandomString(6)

	lic_data_str, err := lic_data.ToJson()
	return lic_data_str, err
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
