package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main()  {
	// 7: 读取缓冲区中数据字节数(只有执行读才会使用到缓冲区, 否则是没有的)
	inputReadBuf4 := strings.NewReader("中文1234567890")
	reader4 := bufio.NewReader(inputReadBuf4)
	// 下面返回0, 因为还没有开始读取, 缓冲区没有数据
	fmt.Println(reader4.Buffered())
	// 下面返回strings的整体长度16(一个人中文是3长度)
	reader4.Peek(1)
	fmt.Println(reader4.Buffered())
	// 下面返回15, 因为readByte已经读取一个字节数据, 所以缓冲区还有15字节
	reader4.ReadByte()
	fmt.Println(reader4.Buffered())
	// 下面的特别有意思: 上面已经读取了一个字节, 想当于是将"中"读取了1/3, 那么如果现在使用readRune读取, 那么
	// 由于无法解析, 那么仅仅读取一个byte, 所以下面的结果很显然
	// 第一次: 无法解析, 那么返回一个byte, 所以输出的是14
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())
	// 第二次读取, 还剩下"中"最后一个字节, 所以也会err, 所以输出13
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())
	// 现在"中"读完了, 那么开始完整读取"文", 这个OK的, 可以解析的, 所以可以读取三字节, 那么剩下10字节
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())

}