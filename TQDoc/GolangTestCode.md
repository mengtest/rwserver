#### int取值范围 参考 
|类型    | 取值范围      |  字节数 |
| ------ | ------        | ------ |
| int    | 根据运行平台可能是32位或64位  | ------ |
| int8   | -128到127     | 占用一个字节 |
| uint8  | 0到255        | 占用一个字节 |
| int16  | -32768到32767 | 占用2个字节 |
| uint16 | 0到65535      | 占用2个字节 |
| int32  | -2147483648到2147483647  |占用4个字节 |
| uint32 | 0到4294967295 |占用4个字节 |
| int64  | -9223372036854775808到9223372036854775807 |占用8个字节 |
| uint64 | 0到18446744073709551615 |占用8个字节 |



##Golang tcp 测试代码
1.启动TQGate的GateServer  
2.运行测试代码
```
package main 

func main(){
   content:="{\"code\":\"200\",\"msg\":\"sdasdasdasda\"}"
   b, _ :=EncodeHead2Byte(string(content))
   conn.Write(b)
}

func EncodeHead2Byte(message string) ([]byte, error) {
	var length  = int16(len(message))
	var pkg  = new(bytes.Buffer)

	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

```
