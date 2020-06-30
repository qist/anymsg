package http

const testStr = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>test</title>
    
     <style>
            body {
                font-family: "Arial","SimHei","SimSun",sans-serif;
                font-size:0.8em;
                background-image:url('/static/img/nn.jpg');
                background-size:cover
                width:100%;
                height:100%;

                }
            .box {
                border:1px solid transparent;
                border-color:#ddd;
                box-shadow:0 1px 1px rgba(0,0,0,.05);
                height:45%;
                width:45%;
        
            }

            .content {
                color:#aa180b;
                font-size:1.5em;
                background:rgb(168, 191, 193);
                }

            .header{
                height:10%;
                background:rgb(75, 132, 138);
                font-size:2.0em;
                color:#0000ff;
                text-align:center;
                
            }
            span {
                color:#0000ff;
                letter-spacing:normal;
            }
            
         </style>
</head>

<body>
    <div class="box">
        <div class="header">邮件</div>
       
        <div class="content">
            <form action="/sender/mail" method="post">
                <label>发件人:</label>
                <input type="text" name="from" value=""><span>可为空,必须为xxx@staff.qkame.com格式</span>
                <br>
                <label>收件人:</label>
                <input type="text" name="to" value=""><span>*多个用","分隔</span>
                <br>
                <label>标题:</label>
                <input type="text" name="subject" value="">
                <br>
                <label>内容:</label>
                <textarea name="content" cols="30" rows="4"></textarea>
                <br>
                <label>内容类型:</label>
                <input type="text" name="content_type" value=""><span>text or html</span>
                <br>
                <br><br>
                <input type="submit" value="提交">
            </form> 
        </div>
    </div>
    <br>
    <div class="box">
        <div class="header">微信</div>
       
        <div class="content">
            <form action="/sender/wechat" method="post">
                <label>接受者:</label>
                <input type="text" name="to" value=""><span>*多个用"|"分隔</span>
                <br>
                <label>消息:</label>
                <textarea name="content" cols="30" rows="4"></textarea>  
                <br><br>
                <input type="submit" value="提交">
            </form> 
        </div>
    </div>


</body>
</html>
  `
