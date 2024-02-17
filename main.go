package main

import (
	"bestrui/wechatpush/mail"
	"bestrui/wechatpush/openwechat"
	"fmt"
	"strings"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("/app/data/storage.json")

	defer reloadStorage.Close()

	// 登录
	if err := bot.HotLogin(reloadStorage); err != nil {
		fmt.Println("热登陆失败，尝试免扫码登录")
		bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	}

	fmt.Println("登陆成功")

	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsSendBySelf() { //自己发送的消息
			//跳过
			return
		} else if msg.IsSendByFriend() { //好友发送的消息
			friendSender, err := msg.Sender()
			if err != nil {
				fmt.Println(err)
				return
			}

			friendSenderName := friendSender.RemarkName
			if len(friendSender.RemarkName) == 0 {
				friendSenderName = friendSender.NickName
			}

			if msg.IsText() {
				fmt.Println(friendSenderName, ":", msg.Content)
				mail.SendEmail(friendSenderName, msg.Content)
			} else if msg.IsPicture() {
				fmt.Println(friendSenderName, ":", "[图片]")
				mail.SendEmail(friendSenderName, "[图片]")
			} else if msg.IsVoice() {
				fmt.Println(friendSenderName, ":", "[语音]")
				mail.SendEmail(friendSenderName, "[语音]")
			} else if msg.IsVideo() {
				fmt.Println(friendSenderName, ":", "[视频]")
				mail.SendEmail(friendSenderName, "[视频]")
			} else if msg.IsEmoticon() {
				fmt.Println(friendSenderName, ":", "[动画表情]")
				mail.SendEmail(friendSenderName, "[动画表情]")
			}
		} else { //群聊发送的消息
			groupSender, err := msg.SenderInGroup()
			if err != nil {
				fmt.Println(err)
				return
			}
			if msg.IsText() {
				//群聊中只接受 @所有人 消息
				if strings.Contains(msg.Content, "@所有人") {
					fmt.Println(groupSender.NickName, ":", msg.Content)
					mail.SendEmail(groupSender.NickName, msg.Content)
				}
			}
		}
	}

	bot.Block()
}
