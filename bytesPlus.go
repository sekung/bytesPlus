package bytesPlus

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"math"
	"strconv"
	"strings"
	"time"
)

// 获取当前时间
func GBNowTime() []byte {
	now := time.Now()
	t := strings.Split(now.Format("05-04-15-02-01-06"), "-")
	var newBytes []byte
	for _, v := range t {
		t, _ := strconv.Atoi(v)
		newBytes = append(newBytes, byte(t))
	}
	return newBytes
}

// 获取当前时间
func NowTimeBCD() []byte {
	b, _ := hex.DecodeString(time.Now().Format("060102150405"))
	return b
}

// 将源字节切片翻转
func Reversed(byt []byte) {
	for i, j := 0, len(byt)-1; i < j; i, j = i+1, j-1 {
		byt[i], byt[j] = byt[j], byt[i]
	}
}

// 字节翻转，生成一个新字节切片
func Reverse(byt []byte) []byte {
	var newBytes []byte
	for i := len(byt) - 1; i > -1; i-- {
		newBytes = append(newBytes, byt[i])
	}
	return newBytes
}

// 得到一个新的切片为原切片的视图
func Slice(byt []byte, startIndex int, endIndex int) []byte {
	max := len(byt)
	var start, end int
	switch {
	case startIndex >= 0 && startIndex <= max:
		start = startIndex
	case startIndex > max:
		panic(fmt.Sprintf("startIndex out of byt range, max is %d", len(byt)))
	case startIndex < 0 && startIndex >= -max:
		start = max + startIndex
	default:
		panic(fmt.Sprintf("startIndex out of byt range, max is %d", len(byt)))
	}
	switch {
	case endIndex == 0:
		end = max
	case endIndex > 0 && endIndex <= max:
		end = endIndex
	case endIndex > max:
		panic(fmt.Sprintf("endIndex out of byt range, max is %d", len(byt)))
	case endIndex < 0 && endIndex >= -max:
		end = max + endIndex
	default:
		panic(fmt.Sprintf("endIndex out of byt range, max is %d", len(byt)))
	}
	if end < start {
		panic("endIndex is too larger then startIndex")
	} else {
		return byt[start:end]
	}
}

// 在index处前插入一个字节切片, 生成一个新字节切片
func Insert(byt []byte, index int, b []byte) []byte {
	max := len(byt)
	var start int
	switch {
	case index >= 0 && index <= max:
		start = index
	case index > max:
		panic(fmt.Sprintf("insert index out of byt range, max is %d", len(byt)))
	case index < 0 && index >= -max:
		start = max + index
	default:
		panic(fmt.Sprintf("insert index out of byt range, max is %d", len(byt)))
	}
	var newBytes []byte
	newBytes = append(newBytes, byt[:start]...)
	newBytes = append(newBytes, b...)
	newBytes = append(newBytes, byt[start:]...)
	return newBytes
}

// 弹出字节切片index的值, 生成一个新字节切片
func Pop(byt []byte, index int) []byte {
	max := len(byt) - 1
	var start int
	switch {
	case index >= 0 && index <= max:
		start = index
	case index > max:
		panic(fmt.Sprintf("Pop index out of byt range, index max is %d", max))
	case index < 0 && index >= -max-1:
		start = max + 1 + index
	default:
		panic(fmt.Sprintf("Pop index out of byt range,index max is %d", max))
	}
	var newBytes []byte
	if start < max {
		newBytes = append(newBytes, byt[:start]...)
		newBytes = append(newBytes, byt[start+1:]...)
	} else {
		newBytes = append(newBytes, byt[:start]...)
	}
	return newBytes
}

// 删除元切片的一段类容，生成一个新字节切片
func Del(byt []byte, startIndex int, endIndex int) []byte {
	max := len(byt)
	var start, end int
	switch {
	case startIndex >= 0 && startIndex <= max:
		start = startIndex
	case startIndex > max:
		panic(fmt.Sprintf("startIndex out of byt range, max is %d", len(byt)))
	case startIndex < 0 && startIndex >= -max:
		start = max + startIndex
	default:
		panic(fmt.Sprintf("startIndex out of byt range, max is %d", len(byt)))
	}
	switch {
	case endIndex == 0:
		end = max
	case endIndex > 0 && endIndex <= max:
		end = endIndex
	case endIndex > max:
		panic(fmt.Sprintf("endIndex out of byt range, max is %d", len(byt)))
	case endIndex < 0 && endIndex >= -max:
		end = max + endIndex
	default:
		panic(fmt.Sprintf("endIndex out of byt range, max is %d", len(byt)))
	}
	if end < start {
		panic("endIndex is too larger then startIndex")
	} else {
		var newBytes []byte
		newBytes = append(newBytes, byt[:start]...)
		newBytes = append(newBytes, byt[end:]...)
		return newBytes
	}
}

// 合并多个切片
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// 合并多个切片
func Combine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// 切片求和
func Sum(byt []byte) int {
	sum := 0
	for i := 0; i < len(byt); i++ {
		sum += int(byt[i])
	}
	return sum
}

// sun8校验
func CheckSum8(byt []byte) int {
	sum := 0
	for i := 0; i < len(byt); i++ {
		sum += int(byt[i])
	}
	return sum & 0xFF
}

// sun8校验
func CheckSum8toByte(byt []byte) byte {
	var sum byte = 0
	for i := 0; i < len(byt); i++ {
		sum += byt[i]
	}
	return sum
}

// sun16校验
func CheckSum16(byt []byte) int {
	sum := 0
	for i := 0; i < len(byt); i++ {
		sum += int(byt[i])
	}
	return sum & 0xFFFF
}

// sun16校验
func CheckSum16Byte(byt []byte) []byte {
	cn := make([]byte, 2)
	binary.BigEndian.PutUint16(cn, uint16(CheckSum16(byt)))
	return cn
}

// CRC modbus校验 返回 int
func CheckCRCModbus(byt []byte) int {
	num1 := 0xFFFF
	num2 := 0xA001
	for _, v := range byt {
		num1 ^= int(v)
		for i := 0; i < 8; i++ {
			last := num1 % 2
			num1 >>= 1
			if last == 1 {
				num1 ^= num2
			}
		}
	}
	return num1
}

// CRC modbus校验 返回 int
func CheckCRCModbusRes(byt []byte) int {
	return Dec(CheckCRCModbusByte(byt))
}

// CRC modbus校验 返回 []byte
func CheckCRCModbusByte(byt []byte) []byte {
	var num1 uint16 = 0xFFFF
	var num2 uint16 = 0xA001
	for _, v := range byt {
		num1 ^= uint16(v)
		for i := 0; i < 8; i++ {
			last := num1 % 2
			num1 >>= 1
			if last == 1 {
				num1 ^= num2
			}
		}
	}
	cn := make([]byte, 2)
	binary.LittleEndian.PutUint16(cn, num1)
	return cn
}

// CRC modbus校验 将结果赋值拼接到原切片莫尾
func CheckCRCModbusMerge(byt []byte) []byte {
	return BytesCombine(byt, CheckCRCModbusByte(byt))
}

func CheckCRCXmodem(buf []byte) int {
	var crc16tab = []uint16{
		0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50a5, 0x60c6, 0x70e7,
		0x8108, 0x9129, 0xa14a, 0xb16b, 0xc18c, 0xd1ad, 0xe1ce, 0xf1ef,
	}
	var crc uint16 = 0
	var ch byte = 0
	size := len(buf)
	for i := 0; i < size; i++ {
		ch = byte(crc >> 12)
		crc <<= 4
		crc ^= crc16tab[ch^(buf[i]/16)]
		ch = byte(crc >> 12)
		crc <<= 4
		crc ^= crc16tab[ch^(buf[i]&0x0f)]
	}
	return int(crc)
}

// BCC 校验 异或和校验
func CheckBCC(byt []byte) int {
	sum := byt[0]
	for i := 1; i < len(byt); i++ {
		sum ^= byt[i]
	}
	return int(sum)
}

// BCC 校验 异或和校验
func CheckBCCToByte(byt []byte) byte {
	sum := byt[0]
	for i := 1; i < len(byt); i++ {
		sum ^= byt[i]
	}
	return sum
}

// 字节拼接取整
func Dec(byt []byte) int {
	l := len(byt)
	sum := 0
	for i := 0; i < len(byt); i++ {
		sum += int(byt[i]) << ((l - i - 1) * 8)
	}
	return sum
}

//读字节切换进行拆分，返回切片数组
func DeBuff(buff, b []byte) [][]byte {
	var rt [][]byte
	for {
		if bytes.Index(buff, b) == -1 {
			rt = append(rt, buff)
			break
		} else {
			index := bytes.Index(buff, b)
			data := buff[:index+1]
			rt = append(rt, data)
			buff = buff[index+1:]
		}
	}
	return rt
}

// 字节切片转16进制字符串
func Hex(byt []byte) string {
	return hex.EncodeToString(byt)
}

// 4字节切片转浮点数,大端
func Bytes32ToFloatBe(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(b))
}

// 4字节切片转浮点数,小端
func Bytes32ToFloatLe(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

// 将相邻的两个字节进行合并
func ByteMerge3X3XToXX(byt []byte) []byte {
	l := len(byt)
	if l%2 != 0 {
		panic(fmt.Sprintf("The byte lens is %d。must be double ", l))
	}
	var newBytes []byte
	for i := 0; i < l; i += 2 {
		by := byt[i : i+2]
		newBytes = append(newBytes, (by[0]&0x0f)<<4|(by[1]&0x0f))
	}
	return newBytes

}

// 字节解码
//你可以使用GBK、gb18030、gb2312、utf8、utf16le、utf16be 进行解码
func Decode(byt []byte, code string) (string, error) {
	switch code {
	case "GBK", "gbk":
		gbkData, err := simplifiedchinese.GBK.NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(gbkData), nil
		}
	case "GB18030", "gb18030":
		gbkData, err := simplifiedchinese.GB18030.NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(gbkData), nil
		}
	case "gb2312", "GB2312":
		gbkData, err := simplifiedchinese.HZGB2312.NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(gbkData), nil
		}
	case "utf8", "utf-8", "UTF8", "UTF-8":
		return string(byt), nil
	case "utf-16-le", "utf16le", "UTF-16-LE", "UTF16LE":
		data, err := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(data), nil
		}
	case "utf-16-be", "utf16be", "UTF-16-BE", "UTF16BE", "UTF16":
		data, err := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder().Bytes(byt)
		if err != nil {
			return "", err
		} else {
			return string(data), nil
		}
	default:
		panic(fmt.Sprintf("unknown decoding: %s", code))
	}
}
