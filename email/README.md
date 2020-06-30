# email

## 请求URL

- `http://ip:port/sender/mail`

## 请求方式：

- POST

## 请求头

>|参数名|是否必须|类型|说明|
>|:----|:---|:----- |:-----|
>|Content-Type |是  |string |请求类型：application/x-www-form-urlencoded&#124;&#124;application/json|

## 请求参数

>|参数名|是否必须|类型|说明|
>|:----|:---|:----- |--------  |
>|from|否  |string | 发件人，必须xxx@staff.qkagame.com格式|
>|to|是  |string | 收件人地址，多个收件人用(,)分隔|
>|subject|是  |string | 邮件标题|
>|content|是  |string | 邮件内容|
>|content_type|否  |string |填 `html`or`txt` 默认`txt`|
