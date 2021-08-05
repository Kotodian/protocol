###微信消息模板

1.微信服务提醒消息
模板ID：OPENTM206919910
模板内容：

```
{{first.DATA}}
时间：{{keyword1.DATA}}
内容：{{keyword2.DATA}}
{{remark.DATA}}
```


2.微信退款提醒消息
模板ID：TM00004
模板内容：

```
{{first.DATA}}
退款原因：{{reason.DATA}}
退款金额：{{refund.DATA}}
{{remark.DATA}}
```


3.微信设备状态变化提醒消息
模板ID：OPENTM206305613
模板内容：

```
{{first.DATA}}
地点：{{keyword1.DATA}}
编号：{{keyword2.DATA}}
类型：{{keyword3.DATA}}
状态：{{keyword4.DATA}}
单号：{{keyword5.DATA}}
{{remark.DATA}}
```

###阿里消息模板

1.阿里服务提醒消息
模板ID ：TM000001374
模板内容：

```
{{first.value}}
提交时间：{{keyword1.value}}
{{remark.value}}
```
2.阿里退款消息
模板ID：TM000003980

```
{{first.value}}
订单编号：{{keyword1.value}}变量长度50个字符 该字段为主要信息
退款金额：{{keyword2.value}}变量长度50个字符 该字段为主要信息
退款编号：{{keyword3.value}}变量长度50个字符
{{remark.value}}
```

3.阿里设备状态变化消息
模板ID：TM000002289

```
{{first.value}}
地点：{{keyword1.value}}
编号：{{keyword2.value}}
类型：{{keyword3.value}}
状态：{{keyword4.value}}
单号：{{keyword5.value}}
{{remark.value}}
```