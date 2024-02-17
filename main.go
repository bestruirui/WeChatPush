package main

import (
	"fmt"

	"wechatpush/openwechat"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	defer reloadStorage.Close()

	// 执行热登录
	if err := bot.HotLogin(reloadStorage); err != nil {
		bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	}

	bot.MessageHandler = func(msg *openwechat.Message) {
		sender, err := msg.Sender()
		if err != nil {
			fmt.Println(err)
			return
		}
		if msg.IsText() {
			fmt.Println(sender.RemarkName, ":", msg.Content)
		} else if msg.IsPicture() {
			fmt.Println(sender.RemarkName, ":", "发送了一张图片")
		} else if msg.IsVoice() {
			fmt.Println(sender.RemarkName, ":", "发送了一段语音")
		} else if msg.IsVideo() {
			fmt.Println(sender.RemarkName, ":", "发送了一段语音")
		} else if msg.IsEmoticon() {
			fmt.Println(sender.RemarkName, ":", "发送了一个表情")
		}

	}
	bot.Block()
}
