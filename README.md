# WxMIPush

> 感谢[eatmoreapple/openwechat](https://github.com/eatmoreapple/openwechat)提供的API  

微信消息转发，仅转发，无回复功能
- [x] 借助QQ邮箱使用MIPush推送，再也不用把微信挂后台了
- [x] 无需重复扫码登录
- [ ] 支持多个微信号同时转发

## 缺陷
- 只支持文本消息转发，其余消息只会发送通知，不能查看
- 群聊消息只接受`@所有人`消息，可以自己修改代码
## 部署
安装docker
```
curl -sSL https://get.docker.com/ | sh
```
运行容器，修改好环境变量的值
```
docker run -d \
	--name WxPush \
	-e FROM_NAME=#发送者邮箱 \
	-e FROM_ADDRESS=#发送者名字 \
	-e TO_NAME=#收件人名字 \
	-e TO_ADDRESS=#收件人邮箱 \
	-e SMTP_SERVER=smtp.exmail.qq.com \
	-e SMTP_PORT=465 \
	-e USERNAME=#smtp的邮箱地址 和FROM_NAME的值是一样的 \
	-e PASSWORD=#smtp的邮箱密码 \
	-v ./WxPush:/app/data \
	bestrui/wxpush:latest
# 如果镜像拉取速度慢，可以使用代理`docker.nju.edu.cn/bestrui/wxpush:latest`
```
容器运行后会在控制台打印登录二维码的链接
```
docker logs WxPush -f
```
建议增加一个cron定时任务，不要24小时一直挂着
```
crontab -e

46 23 * * * docker stop WxPush  # 每天23:46 停止容器
33  8 * * * docker start WxPush # 每天 8:33 启动容器
```