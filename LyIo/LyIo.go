package LyIo

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

/*
var EOF = errors.New("EOF")
EOF当无法得到更多输入时，Read方法返回EOF。当函数一切正常的到达输入的结束时，就应返回EOF。
如果在一个结构化数据流中EOF在不期望的位置出现了，
则应返回错误ErrUnexpectedEOF或者其它给出更多细节的错误。
*/

/*
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
当从一个已关闭的Pipe读取或者写入时，会返回ErrClosedPipe。
*/

func TestLimitReader() {
	f, err := os.Open("Files\\test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := io.LimitReader(f, 10)
	fmt.Println(reader)
	p := make([]byte, 10)
	var total int
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Println("read value", string(p[:total])) //read value hello
			fmt.Println(total)                           //5
			break
		}
		total = total + n
	}
	fmt.Println(total)
}

/*
MultiReader返回一个将提供的Reader在逻辑上串联起来的Reader接口。
他们依次被读取。当所有的输入流都读取完毕，Read才会返回EOF。
如果readers中任一个返回了非nil非EOF的错误，Read方法会返回该错误。
*/
func TestMultiReader() {
	f1, err := os.Open("Files\\a.txt")
	f2, err := os.Open("Files\\b.txt")
	if err != nil {
		return
	}
	defer f1.Close()
	defer f2.Close()
	reader := io.MultiReader(f1, f2)
	fmt.Println(reflect.TypeOf(reader))
	p := make([]byte, 10)
	var total int
	var data string
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Println("read end", total) //read end 17
			break
		}
		total = total + n
		data = data + string(p[:n])
	}
	fmt.Println("read value", data)  //read value widuu2hello widuu
	fmt.Println("read count", total) // read count 17
}
