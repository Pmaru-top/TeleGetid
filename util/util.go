package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v4"
)

var TeleDataCenter = map[int]string{
	1: "MIA, Miami FL, USA",
	2: "AMS, Amsterdam, NL",
	3: "MIA, Miami FL, USA",
	4: "AMS, Amsterdam, NL",
	5: "SIN, Singapore, SG",
}

type TeleDCInfo struct {
	DCNum      int
	DCLocation string
}

func SetupProxy(proxyURL string) {
	if len(proxyURL) != 0 {
		os.Setenv("HTTP_PROXY", proxyURL)
		os.Setenv("HTTPS_PROXY", proxyURL)
	}
}

type AutoReconnectPoller struct {
	BasePoller tele.Poller
}

// reconnection
func (p *AutoReconnectPoller) Poll(b *tele.Bot, updates chan tele.Update, stop chan struct{}) {
	for {
		p.BasePoller.Poll(b, updates, stop)

		log.Println("trying to reconnect...")
		time.Sleep(5 * time.Second)
	}
}

func WithCode(msg string) string {
	return fmt.Sprintf("<code>%s</code>", msg)
}

func GetTeleDCInfo(userName string) (dcInfo *TeleDCInfo, err error) {
	if len(userName) == 0 {
		err = fmt.Errorf("userName is empty")
		return
	}

	url := fmt.Sprintf("https://t.me/%s", userName)
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	text := string(data)

	re := regexp.MustCompile(`cdn(\d).cdn-telegram.org`)
	match := re.FindStringSubmatch(text)

	if len(match) <= 1 {
		err = fmt.Errorf("get dcNum Failed")
	}

	dcInfo = &TeleDCInfo{}
	dcNum, _ := strconv.Atoi(match[1])
	dcInfo.DCNum = dcNum
	dcInfo.DCLocation = TeleDataCenter[dcNum]

	return
}
