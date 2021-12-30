# anymsg

- [邮件API](https://github.com/qist/anymsg/blob/master/email/README.md)

- [企业微信API](https://github.com/qist/anymsg/blob/master/wechat/README.md)

## cfg.json 是纯json格式

> ```json
>{
>    "debug": true,
>    "http": {
>        "listen": "0.0.0.0:4000", //监听ip端口
>        "allow":["*"],// 填写ip，"*" 代表允许全部
>        "deny":[]
>    },
>    "smtp": {//邮件配置
>        "address": "smtp.exmail.qq.com:25",//邮件发送服务器地址
>        "username": "qist@example.com",
>        "password": "123456",
>        "authtype":"LOGIN"//认证类型/CRAM-MD5/LOGIN/PLAIN,默认PLAIN
>    },
>    "wechat":{//企业微信配置
>        "CorpID":"ww2085a342", //企业ID
>        "AgentId":1000002,//应用id，通过新建企业微信应用>获取
>        "Secret":"5WsjwD2DqyR4PMTWnJJp_qvyOothRjDAZs>aKc"//密串，企业微信应用中可以得到
>    }
>    "dingding":{//钉钉配置
>        "Url":"https://oapi.dingtalk.com/robot/send?access_token=%s", //钉钉机器人连接地址
>        "AccessToken":"xxxxxxxx" //创建钉钉机器人后连接地址access_token后面的字符串

>    }
>    "lark":{//飞书配置
>        "Url":"https://open.feishu.cn/open-apis/bot/v2/hook/%s", //钉钉机器人连接地址
>        "AccessToken":"xxxxxxxx" //创建飞书机器人后连接地址https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxxxxxxxxxxx ， xxx的字符串就是AccessToken

>    }
>}
>```

## 测试方法

>```shell
>普通文本（text）发送格式
>curl -d "to=test@qq.com,test@sina.com&subject=test&content=测试报文体" "http://10.1.1.202:4000/sender/mail"
>curl -d "to=qist&&content=测试报文体" "http://10.1.1.202:4000/sender/wechat"
>curl -d "to=qist&&content=测试报文体" "http://10.1.1.202:4000/sender/dingding"
>markdown格式的消息，请求体内关键字段："msg_type": "markdown"
>企业微信
>curl -H 'Content-Type: json'  -d '{"to":"wangGang","content": "您的会议室已经预定，稍后会同步到`邮箱 \n  **事项详情** \n 事　项：<font color=\"info\">开会</font> \n 组织者：@miglioguan>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang> \n 会议室：<font color=\"info\">广州TIT 1楼 301</font> \n日　期：<font color=\"warning\">2018年5月18日</font> \n 时　间：<font color=\"comment\">上午9:00-11:00</font> > \n 请准时参加会议。> \n 如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)" ,"msg_type": "markdown"}' "http://10.1.1.202:4000/sender/wechat"
>钉钉机器人（钉钉机器人需要添加关键字，关键字在content里面包含就行，例如这里的"监控报警"）
>curl 127.0.0.1:4000/sender/dingding -H 'Content-Type: json' -d '{"to":"周浩","title":"监控报警","content": "#### 监控报警 \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n" ,"msg_type": "markdown"}'
>飞书
>curl -H 'Content-Type: json' -d '{"content":"监控报警"}' "http://127.0.0.1:4000/sender/lark"
>```

```git
echo "# msgsender" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add origin https://github.com/qist/anymsg.git
git push -u origin master
```

