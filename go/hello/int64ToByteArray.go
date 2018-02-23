package main

import (
	"fmt"
	"encoding/binary"
)

/*
go语言中byte是uint8的别名，即8位二进制;
go语言数据转二进制时，按照无符号数进行；
如果按照二进制转无符号，则将当前二进制是无符号数；
如果按照二进制转有符号，则将当前二进制是有符号数；
*/
func main() {
	var i int64 = -123
	buf := Int64ToBytes(i)
	fmt.Println(int64(i))
	fmt.Println(uint64(i))
	fmt.Println(buf)
	fmt.Println(BytesToInt64(buf))
	fmt.Println(BytesToUint64(buf))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func BytesToUint64(buf []byte) uint64 {
	return uint64(binary.BigEndian.Uint64(buf))
}

// 这里是大端模式
func f2() {
	var v2 uint32
	var b2 [4]byte
	v2 = 257
	// 将 256转成二进制就是
	// | 00000000 | 00000000 | 00000001 | 00000001 |
	// | b2[0]    | b2[1]   | b2[2]    | [3]      | // 这里表示b2数组每个下标里面存放的值

	// 这里直接使用将uint32l强转成uint8
	// | 00000000 0000000 00000001 | 00000001  直接转成uint8后等于 1
	// |---这部分go在强转的时候扔掉---|
	b2[3] = uint8(v2)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移8位 转成uint8后等于 1
	// 下面是右移后的数据
	// |          | 00000000 | 00000000 | 00000001 |
	b2[2] = uint8(v2 >> 8)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移16位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          | 00000000 | 00000000 |
	b2[1] = uint8(v2 >> 16)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移24位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          |          | 00000000 |
	b2[0] = uint8(v2 >> 24)

	fmt.Printf("%+v\n", b2)
	// 所以最终将uint32转成[]byte数组输出为
	// [0 0 1 1]
}

// 这里是小端模式
// 在上面我们讲过，小端刚好和大端相反的，所以在转成小端模式的时候，只要将[]byte数组的下标首尾对换一下位置就可以了
func f3() {
	var v3 uint32
	var b3 [4]byte
	v3 = 257
	// 将 256转成二进制就是
	// | 00000000 | 00000000 | 00000001 | 00000001 |
	// | b3[0]    | b3[1]   | b3[2]    | [3]      | // 这里表示b3数组每个下标里面存放的值

	// 这里直接使用将uint32l强转成uint8
	// | 00000000 0000000 00000001 | 00000001  直接转成uint8后等于 1
	// |---这部分go在强转的时候扔掉---|
	b3[0] = uint8(v3)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移8位 转成uint8后等于 1
	// 下面是右移后的数据
	// |          | 00000000 | 00000000 | 00000001 |
	b3[1] = uint8(v3 >> 8)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移16位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          | 00000000 | 00000000 |
	b3[2] = uint8(v3 >> 16)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移24位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          |          | 00000000 |
	b3[3] = uint8(v3 >> 24)

	fmt.Printf("%+v\n", b3)
	// 所以最终将uint32转成[]byte数组输出为
	// [1 1 0 0 ]
}