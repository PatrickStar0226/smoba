# smoba
* 王者营地自动签到每周浏览点赞任务
* 支持多账号
* 微信推送
* 华为云函数，免费自动每天刷

## 使用手册
[使用手册](https://gosmoba.vercel.app/)

## 编译使用
*如果不使用华为云函数请在函数入口main下直接使用run()函数*
```go
func main() {
	run()
}
```
*编译(Linux) 华为云函数*
```Terminal
go mod tidy
set GOOS=linux
set GOARCH=amd64
go build -o handler main.go
zip main.zip handler config.go static
```

*编译(Windows)*
```cmd
go mod tidy
set GOOS=windows
set GOARCH=386
go build -o handler main.go
```

## 免责声明
* 本厂库内容仅用于学习研究，禁止用于任何商业或非法用途，违反者作者概不负责。
* 如果任何单位或个人认为该厂库内容可能涉嫌侵犯其权利。则应及时通知并提供身份证明，所有权证明，我将在收到认证文件后删除厂库内容。

---

> 以上内容 fork 自 20220623 https://github.com/superggfun/smoba 未作改动

---

## 个人修改
* 新增企业微信-机器人推送方式

### 企业微信推送方式说明
> 企业微信应用消息配置说明参照 https://github.com/easychen/wecomchan#企业微信应用消息配置说明
> 
> 群机器人配置说明参照 https://developer.work.weixin.qq.com/document/path/91770

#### 第一步，注册企业
用电脑打开企业微信官网，注册一个企业
#### 第二步，创建机器人并获取到Webhook地址
1. 注册成功后需 手机企业微信登录,在一个群聊中点击右上角-菜单-群机器人-添加-新建
2. 输入机器人名称,点击添加.
3. 在添加完成界面获取到`Webhook`地址,或者在一个群聊中点击右上角-菜单-群机器人 点击需要推送的机器人,也可以获取到`Webhook`地址
#### 第三步,配置Webhook地址
将第二步中获取到的`Webhook`地址,填写进`config.json`中
>例如:
> 
>```
>"qywxpush": {
>   "webhookurl": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
>}
>```

> 该`Webhook`地址 是官方配置文档中的示例,不是实际地址
> 
> 特别特别要注意：一定要<font color=red><b>保护好机器人的webhook地址</b></font>，避免泄漏！不要分享到github、博客等可被公开查阅的地方
