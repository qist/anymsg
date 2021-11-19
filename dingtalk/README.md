# wechat-sender

## 请求URL

>- `http://ip:port/sender/wechat`

### 请求方式：

- POST

#### 请求头

>|参数名|是否必须|类型|说明|
>|:----|:---|:----- |:-----|
>|Content-Type |是  |string |请求类型：application/x-www-form-urlencoded &#124;&#124; application/json|

### 请求参数
```
 参数名 是否必须     类型   说明
to  是   string  收件人地址，多个接收者用|分割 全体填@all
content 是  string  消息 内容
contentType 否   string  填 text(保留字段)
```
