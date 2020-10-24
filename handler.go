// DailyGakki - handler
// 2020-10-17 14:03
// Benny <benny.think@gmail.com>

package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"path/filepath"
	"strconv"
	"strings"
)

func startHandler(m *tb.Message) {
	caption := "欢迎来到每日最可爱的Gakki！\n我会每天定是为你发送最可爱的Gakki！"
	filename := "start.gif"

	log.Infof("Start command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

}

func aboutHandler(m *tb.Message) {
	caption := "欢迎来到每日最可爱的Gakki！\n" +
		"开发者：@BennyThink\n" +
		"GitHub: https://github.com/BennyThink/DailyGakki/" +
		"Google Photos 地址：" + album
	filename := "about.gif"

	log.Infof("About command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

}

func newHandler(m *tb.Message) {
	log.Infof("New command: %d", m.Chat.ID)

	// 默认发送3张
	_ = b.Notify(m.Chat, tb.Typing)
	sendAlbum := generatePhotos()
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	_, _ = b.SendAlbum(m.Chat, sendAlbum)

}

func settingsHandler(m *tb.Message) {
	log.Infof("Settings command: %d", m.Chat.ID)

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "在这里可以设置每日推送时间和每日推送次数")
	var btns []tb.Btn
	var selector = &tb.ReplyMarkup{}

	btn := selector.Data("Placeholder", fmt.Sprintf("Placeholder%s%d",
		"Placeholder", m.Chat.ID), "Placeholder")
	//registerButtonNextStep(btn, "addServiceButton")
	btns = append(btns, btn)

	selector.Inline(
		selector.Row(btns...),
	)

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "呀！这部分功能还没做！😅", selector)

}

func registerButtonNextStep(btn tb.Btn, fun func(c *tb.Callback)) {
	log.Infoln("Registering ", btn.Unique)
	b.Handle(&btn, fun)
}

func subHandler(m *tb.Message) {
	// check permission first
	canSubscribe := checkSubscribePermission(m)
	if !canSubscribe {
		log.Infof("Denied subscribe request for: %d", m.Sender.ID)
		_ = b.Notify(m.Chat, tb.Typing)
		_, _ = b.Send(m.Chat, "ええ😉只有管理员才能进行设置哦")
		return
	}

	caption := "已经订阅成功啦！将在每晚18:11准时为你推送最可爱的Gakki！"
	filename := "sub.gif"

	log.Infof("Sub command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

	add(m.Chat.ID)

}

func unsubHandler(m *tb.Message) {
	canSubscribe := checkSubscribePermission(m)
	if !canSubscribe {
		log.Infof("Denied subscribe request for: %d", m.Sender.ID)
		_ = b.Notify(m.Chat, tb.Typing)
		_, _ = b.Send(m.Chat, "ええ😉只有管理员才能进行设置哦")
		return
	}
	caption := "Gakki含泪挥手告别😭"
	filename := "unsub.gif"

	log.Infof("Unsub command: %d", m.Chat.ID)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	data, _ := Asset(filepath.Join("images", filename))
	log.Infof("Find %s from memory...", filename)

	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)
	}

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "😭")

	// 读取文件，增加对象，然后写入
	remove(m.Chat.ID)

}

func messageHandler(m *tb.Message) {
	caption := "私は　今でも空と恋をしています。"
	var filename string

	log.Infof("Message Handler: %d", m.Chat.ID)

	switch m.Text {
	case "😘":
		filename = "kiss.gif"
	case "😚":
		filename = "kiss.gif"
	case "😗":
		filename = "kiss.gif"
	case "❤️":
		filename = "heart1.gif"
	case "❤️❤️":
		filename = "heart2.gif"
	case "❤️❤️❤️":
		filename = "heart3.gif"
	case "🌹":
		filename = "rose.gif"
	case "🦎":
		filename = "lizard.gif"
	default:
		filename = "default.gif"
		_, _ = b.Send(m.Chat, "🤪")
		return

	}

	log.Infof("Choose %s for text %s", filename, m.Text)
	data, _ := Asset(filepath.Join("images", filename))

	log.Infof("Send %s now...", filename)
	_ = b.Notify(m.Chat, tb.UploadingPhoto)
	p := &tb.Animation{File: tb.FromReader(bytes.NewReader(data)), FileName: filename, Caption: caption}
	_, err := b.Send(m.Chat, p)
	if err != nil {
		log.Warnf("%s send failed %v", filename, err)

	}

}

func pingHandler(m *tb.Message) {
	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "pong")
}

func statusHandler(m *tb.Message) {
	_ = b.Notify(m.Chat, tb.Typing)
	currentJSON := readJSON()
	var isSub = false
	for _, user := range currentJSON {
		if user.ChatId == m.Chat.ID {
			isSub = true
		}
	}
	if isSub {
		_, _ = b.Send(m.Chat, "Gakki与你同在😄")
	} else {
		_, _ = b.Send(m.Chat, "还木有每日Gakki💔")

	}
}

func checkSubscribePermission(m *tb.Message) (allow bool) {
	allow = false
	if !m.Private() {
		admins, _ := b.AdminsOf(m.Chat)
		for _, admin := range admins {
			if admin.User.ID == m.Sender.ID {
				allow = true
			}
		}
	} else {
		allow = true
	}
	return
}

func photoHandler(m *tb.Message) {
	userID, _ := strconv.Atoi(reviewer)
	mm := tb.Message{
		Sender: &tb.User{
			ID: userID,
		},
	}

	_ = b.Notify(m.Chat, tb.Typing)
	botSent, _ := b.Reply(m, "你的Review已经发出去惹……请耐心等待😄")

	var btns []tb.Btn
	var selector = &tb.ReplyMarkup{}
	// text, unique string, data
	p1, p2 := botSent.MessageSig()
	data := fmt.Sprintf("%v|%v", p1, p2)

	approve := selector.Data("Yes", fmt.Sprintf("Yes%v", m.Photo.UniqueID), data)
	deny := selector.Data("No", fmt.Sprintf("No%v", m.Photo.UniqueID), data)

	registerButtonNextStep(approve, approveButton)
	registerButtonNextStep(deny, denyButton)
	btns = append(btns, approve, deny)

	selector.Inline(
		selector.Row(btns...),
	)

	fwd, err := b.Forward(mm.Sender, m, selector)
	if err != nil {
		log.Errorln(err)
		_, _ = b.Edit(botSent, "呃……由于某种神秘的原因，Review请求发送失败了，你再发一下试试\n"+err.Error())
	} else {
		_, _ = b.Reply(fwd, "请Review", selector)

	}

}

func approveButton(c *tb.Callback) {
	_ = b.Respond(c, &tb.CallbackResponse{Text: "Approved"})
	_ = b.Delete(c.Message)
	_ = b.Delete(c.Message.ReplyTo)

	s := strings.Split(c.Data, "|")
	cid, _ := strconv.ParseInt(s[1], 10, 64)
	m := tb.StoredMessage{
		MessageID: s[0],
		ChatID:    cid,
	}

	_, _ = b.Edit(m, "你的图片被接受了😊")
	photo := c.Message.ReplyTo.Photo
	picPath := filepath.Join(photos, photo.UniqueID+".jpg")
	log.Infof("Downloading photos to %s", picPath)
	err = b.Download(&photo.File, picPath)
	if err != nil {
		log.Errorln("Download failed", err)
	}

}

func denyButton(c *tb.Callback) {
	_ = b.Respond(c, &tb.CallbackResponse{Text: "Denied"})
	_ = b.Delete(c.Message)
	_ = b.Delete(c.Message.ReplyTo)

	s := strings.Split(c.Data, "|")
	cid, _ := strconv.ParseInt(s[1], 10, 64)
	m := tb.StoredMessage{
		MessageID: s[0],
		ChatID:    cid,
	}

	_, _ = b.Edit(m, "你的图片被拒绝了😫")

}

func submitHandler(m *tb.Message) {

	_ = b.Notify(m.Chat, tb.Typing)
	_, _ = b.Send(m.Chat, "想要向我提交新的图片吗？直接把图片发送给我就可以！单张，多张为一组，转发都可以的！\n"+
		"目前暂时还不支持以文件的形式发送。如有问题可以联系 @BennyThink")

}
