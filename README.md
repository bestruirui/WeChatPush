# WxPush
使用了[eatmoreapple/openwechat](https://github.com/eatmoreapple/openwechat)提供的API  
使用Desktop登陆协议
## 部署
```
docker run -it \
	-e FROM_NAME=#发送者邮箱 \
	-e FROM_ADDRESS=#发送者名字 \
	-e TO_NAME=#收件人名字 \
	-e TO_ADDRESS=#收件人邮箱 \
	-e SMTP_SERVER=smtp.exmail.qq.com \
	-e SMTP_PORT=465 \
	-e USERNAME=#smtp的邮箱地址 应该和FROM_NAME的值是一样的 \
	-e PASSWORD=#smtp的邮箱密码 \
	-v ./data:/app/data \
	bestrui/wxpush:1.0    