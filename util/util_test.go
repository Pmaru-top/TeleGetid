package util

import (
	"fmt"
	// "os"
	"testing"
)

func TestParseDC(t *testing.T) {
	// os.Setenv("HTTP_PROXY", "socks5://127.0.0.1:10010")
	// os.Setenv("HTTPS_PROXY", "socks5://127.0.0.1:10010")

	dc, err := GetTeleDCInfo("durov")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("dc:%v location:%s", dc.DCNum, dc.DCLocation)
}
