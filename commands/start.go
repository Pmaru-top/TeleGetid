package commands

import (
	"fmt"
	"strconv"

	"github.com/pmaru-top/telegetid/util"
	tele "gopkg.in/telebot.v4"
)

const PROJECT_URL = "https://github.com/pmaru-top/telegetid"

func OnStart(c tele.Context) error {
	sender := c.Message().Sender

	firstName := sender.FirstName
	lastName := sender.LastName
	fullName := firstName + lastName

	userName := sender.Username
	userId := sender.ID
	lang := sender.LanguageCode

	var dcInfoStr string
	dcInfo, err := util.GetTeleDCInfo(userName)
	if err == nil {
		dcInfoStr = fmt.Sprintf("DC%v - %s",
			dcInfo.DCNum,
			dcInfo.DCLocation)
	}

	msg := fmt.Sprintf(
		`Your User Info
User_name: %s
User_id: %s
First_name: %s
Last_name: %s
Full_name: %s
Lang: %s
%s
<a href="%s">Github</a>`,
		util.WithCode(userName),
		util.WithCode(strconv.FormatInt(userId, 10)),
		util.WithCode(firstName),
		util.WithCode(lastName),
		util.WithCode(fullName),
		util.WithCode(lang),
		util.WithCode(dcInfoStr),
		PROJECT_URL,
	)

	return c.Reply(msg, tele.ModeHTML)
}
