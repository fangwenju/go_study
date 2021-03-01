package socket

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
)

func Link() {
	var tcpAddr *net.TCPAddr
	// tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.61.10:9600")
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "120.26.63.209:9600")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	sendEasyswooleMsg(conn)
}

func sendEasyswooleMsg(conn *net.TCPConn) {
	var sendData []byte
	data := `{"command":1,"request":{"serviceName":"Device","action":"register","arg":{"imei":"MPD20GT04002734","sim":"869098040410544","sn":"865815043978860"}}}`
	// data := `{"command":1,"request":{"serviceName":"Device","action":"state","arg":{"status":"1","deviceId":"7860"}}}`
	// data := `{"command":1,"request":{"serviceName":"Device","action":"callback","arg":{"imei":"MPD20GT04002734","sim":"8986112020700358717B","sn":"865815043978860"}}}`
	b := []byte(data)
	// 大端字节序(网络字节序)大端就是将高位字节放到内存的低地址端，低位字节放到高地址端。
	// 网络传输中(比如TCP/IP)低地址端(高位字节)放在流的开始，对于2个字节的字符串(AB)，传输顺序为：A(0-7bit)、B(8-15bit)。
	sendData = int32ToBytes8(int32(len(data)))
	// 将数据byte拼装到sendData的后面
	for _, value := range b {
		sendData = append(sendData, value)
	}
	conn.Write(sendData)
	result, _ := ioutil.ReadAll(conn)
	fmt.Println(string(result))
}

func int32ToBytes8(n int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	return buf
}
