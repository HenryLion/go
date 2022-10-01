package main

import (
	"encoding/base64"
	"fmt"
)

// base64编码
// 因为不是所有字节都可以用可打印字符打印出来
// 为了将这些字符打印出来，将bit流按照6bit分割，这样只需要2的6次幂
// 共64个可打印字符就能将数据编码
// 64个可打印的字符是 A-Z, a-z, 0-9,外加另外两个如'+'，'/'

func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

}
