# WxMIPush

> 感谢[eatmoreapple/openwechat](https://github.com/eatmoreapple/openwechat)提供的API  

微信消息转发

- 借助QQ邮箱使用MIPush推送，再也不用把微信挂后台了
- 无需重复扫码登录

- [ ] 支持多个微信号同时转发~
## 部署
运行容器
```
docker run -d \
	--name WxPush \
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
```
容器运行后会在控制台打印登录二维码的链接
```
docker logs WxPush
```
