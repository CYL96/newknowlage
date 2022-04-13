package ipToAddr

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type IpSt struct {
	data       []byte
	FirstIndex int64
	LastIndex  int64
	TotalIP    int64
}
type IpInfoExt struct {
	Match bool   `json:"match" comment:"是否匹配到ip地址"`
	Ip    string `json:"ip" comment:"ip地址"`
	City  string `json:"city" comment:"城市"`
	Area  string `json:"area" comment:"区域"`
}

func (this *IpSt) Init(filePath string) (err error) {
	this.data, err = os.ReadFile(filePath)
	if err != nil {
		return err
	}
	this.FirstIndex = ToLong(this.data[0:4])
	this.LastIndex = ToLong(this.data[4:8])
	this.TotalIP = (this.LastIndex-this.FirstIndex)/7 + 1
	return nil
}

// 通过索引区获取到ip的偏移量，然后去纪录区获取纪录
func (this *IpSt) FindIp(ip string) (info IpInfoExt) {
	info.Ip = ip
	// 匹配ip
	ipNum, _ := IpStrToLong(ip)
	ipIndex := int64(-1)
	// TODO 这里可以优化，比如换成二分法，我想偷懒先就顺序查找吧
	for i := int64(0); i < this.TotalIP; i++ {
		start := this.FirstIndex + i*7
		// end := start + 7
		tmpIp := this.data[start : start+4]
		if IpToLong(tmpIp) <= ipNum && (i == this.TotalIP-1 || IpToLong(this.data[start+7:start+11]) > ipNum) {
			ipIndex = start
			break
		}
	}
	//  下面开始去纪录区寻找纪录
	if ipIndex == -1 {
		info.Match = false
		return
	}
	// 开始寻找记录
	recordIndex := ToLong(this.data[ipIndex+4 : ipIndex+7])
	//获取到纪录区的ip结束地址，判断一下ip是否在范围内
	ipEnd := this.data[recordIndex : recordIndex+4]
	if IpToLong(ipEnd) <= ipNum {
		// ip可能并没有在这个范围内
		info.Match = false
		return
	}
	info.Match = true
	dataIndex := recordIndex + 4
	info.City, info.Area = this.GetRecord(dataIndex)
	return
	// return -1
}

// 获取及录取的内容
func (this *IpSt) GetRecord(dataIndex int64) (string, string) {
	var (
		cityIndex int64
		city      string
		endIndex  int64
		areaIndex int64
		area      string
	)
	// 开始寻找记录 这里的寻找方法看官方的文档吧
	switch this.data[dataIndex] {
	case 0x01:
		// 跳过重定向标志
		dataIndex++

		dataIndex = ToLong(this.data[dataIndex : dataIndex+3])
		if this.data[dataIndex] == 0x02 {
			// 跳过重定向标志
			dataIndex++
			cityIndex = ToLong(this.data[dataIndex : dataIndex+3])
			city, _ = this.GetContent(cityIndex)

			if this.data[dataIndex+3] == 0x02 || this.data[dataIndex+3] == 0x01 {
				//记录模式5
				areaIndex = ToLong(this.data[dataIndex+4 : dataIndex+7])
				area, _ = this.GetContent(areaIndex)
			} else {
				// 记录模式4
				areaIndex = dataIndex + 3
				area, _ = this.GetContent(areaIndex)
			}

		} else {
			//记录模式2
			cityIndex = dataIndex
			city, _ = this.GetContent(cityIndex)

			areaIndex = dataIndex + 4
			area, _ = this.GetContent(areaIndex)
		}
	case 0x02:
		//记录模式3
		// 跳过重定向标志
		dataIndex++

		cityIndex = ToLong(this.data[dataIndex : dataIndex+3])
		city, _ = this.GetContent(cityIndex)

		areaIndex = dataIndex + 4
		area, _ = this.GetContent(areaIndex)
	default:
		// 记录模式1
		cityIndex = dataIndex
		city, endIndex = this.GetContent(cityIndex)

		areaIndex = endIndex + 1
		area, _ = this.GetContent(areaIndex)
	}
	return city, area
}

// 通过偏移量获取到内容
func (this *IpSt) GetContent(index int64) (string, int64) {
	i := index
	for ; i < int64(len(this.data)); i++ {
		if this.data[i] == 0x00 {
			//官方规定的是 字符串格式，以二进制0结尾
			break
		}
	}
	b, err := simplifiedchinese.GBK.NewDecoder().Bytes(this.data[index:i])
	if err != nil {
		fmt.Println(err)
	}
	return string(b), i
}

// 将ip 转化 好进行对比
func IpStrToLong(ip string) (ipNum int64, err error) {
	ipS := strings.Split(ip, ".")
	if len(ipS) != 4 {
		err = fmt.Errorf("ipToAddr format error")
		return
	}
	ip_0, _ := strconv.Atoi(ipS[0])
	if ip_0 < 0 || ip_0 > 255 {
		err = fmt.Errorf("ipToAddr format error")
		return
	}
	ip_1, _ := strconv.Atoi(ipS[1])
	if ip_1 < 0 || ip_1 > 255 {
		err = fmt.Errorf("ipToAddr format error")
		return
	}
	ip_2, _ := strconv.Atoi(ipS[2])
	if ip_2 < 0 || ip_2 > 255 {
		err = fmt.Errorf("ipToAddr format error")
		return
	}
	ip_3, _ := strconv.Atoi(ipS[3])
	if ip_3 < 0 || ip_3 > 255 {
		err = fmt.Errorf("ipToAddr format error")
		return
	}
	ipNum = int64(ip_0)<<24 + int64(ip_1)<<16 + int64(ip_2)<<8 + int64(ip_3)
	// ipNum = int64(binary.BigEndian.Uint64([]byte{0, 0, 0, 0, ipS[0], b[2], b[1], b[0]}))
	return

}

// 将[]byte 类型的ip进行转化
func IpToLong(b []byte) int64 {
	switch len(b) {
	case 4:
		return int64(binary.BigEndian.Uint64([]byte{0, 0, 0, 0, b[3], b[2], b[1], b[0]}))
	}
	return 0
}

// 将[]byte 进行转化，获取到文件的index 对应偏移量
func ToLong(b []byte) int64 {
	switch len(b) {
	case 4:
		return int64(binary.BigEndian.Uint64([]byte{0, 0, 0, 0, b[3], b[2], b[1], b[0]}))
	case 3:
		return int64(binary.BigEndian.Uint64([]byte{0, 0, 0, 0, 0, b[2], b[1], b[0]}))
	}
	return 0
}
