package crc16

import (
	"fmt"
	"strconv"
)

func Crc16Verify(msg string) {

	//16位全为1的寄存器
	var crc uint16 = 0xFFFF
	//高8位字节
	var HighByte uint16
	//移出的数位
	var flag uint16

	var msgBytes []byte = []byte(msg) //被校验的信息
	var length int = len(msgBytes)    //被校验的信息的长度
	for i := 0; i < length; i++ {
		HighByte = crc >> 8                  // 16位寄存器的高位字节
		crc = HighByte ^ uint16(msgBytes[i]) //取被校验串的一个字节与16位寄存器的高位字节进行异或运算

		for j := 0; j < 8; j++ {
			flag = crc & 0x0001 //移出的数位
			crc = crc >> 1      // 把这个 16 寄存器向右移一位

			if flag == 0x0001 { //移出位是1与多项式0xA001异或
				crc ^= 0xA001
			}
		}
	}
	//crc = (crc << 8) ^ (crc >> 8) //高低字节对调
	fmt.Printf("%v\n", crc%16384)
}

//通用modbus CRC校验算法
func ModbusCRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		//通用modbus取寄存器的低8位参与异或运算
		crc = ((crc << 8) >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}

	//得到的十六进制校验码是按照高字节在前低字节在后的字符串
	//要翻转，按照低字节在前高字节在后
	//校验码必须是4个字符，不足4位的需要在开头补0
	hex := strconv.FormatInt(int64(crc), 16) //格式化为16进制字符串
	tmp := hex[2:] + hex[:2]
	if len(tmp) == 3 {
		tmp = "0" + tmp
	}
	return tmp
}

//HJ212 CRC校验算法
func Hjt212CRC(dataString string) string {
	crc := 0xFFFF
	length := len(dataString)
	for i := 0; i < length; i++ {
		//hj212取寄存器的高8位参与异或运算
		crc = (crc >> 8) ^ int(dataString[i])
		for j := 0; j < 8; j++ {
			flag := crc & 0x0001
			crc >>= 1
			if flag == 1 {
				crc ^= 0xA001
			}
		}
	}
	//因为是基于右移位运算的结果，得到的本身就是低字节在前高字节在后的结果
	//不足4位的需要在开头补0
	hex := strconv.FormatInt(int64(crc), 16)
	if len(hex) == 3 {
		hex = "0" + hex
	}
	return hex
}
