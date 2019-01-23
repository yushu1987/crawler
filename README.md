# cralwer

抓取网页内容和图片，并推送到es，提供搜索

## Overview
* 根据指定输入的url，抓取其网页内容和图片
* 筛选指定的网页格式和图片格式，存储到本地磁盘中
* 把抓取网页的title 和url 推送到es，提供检索功能


## 1.1

* 抓取源改成kafka读取，从上有定时推送抓取任务源
*  修复视频下载的错误
* 修复title编码问题 

## 1.2
* 网页地址和内容保存云存储，方案待调研 
* 把html中meta keywords，description等 推送到es
* 增加定时脚本回扫抓取的url，如果已实效，则删除记录，并把es的数据删除

## License

© wangjian, 2019~time.Now
~
~
