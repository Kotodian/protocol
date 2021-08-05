# 桩到后台的通讯协议规范


## MQTT
以硬件的角色为主，pub为硬件的发送通道（即软件的接收通道），sub为硬件的接收通道（即软件的发送通道）

### Topic的path规范
> /[hardware]/[id]/[action]/[priority]

| 字段 | 说明 |
| --- | --- |
| hardware | 目前支持box、charger  |
| id  | 硬件设备的id |
| action  | 操作行为，如pub,sub |
| priority  | [0,1] 0为优先级较低的通道，qos=0不支持retained；1为优先级较高的通道，支持retained |

### charger的pub通道
| topic | qos | 支持retained | 说明 |
| --- | --- | --- |--- |
| /charger/[id]/pub/0 | 0 | 不支持 | 硬件上报优先级较低的通道，且qos=0，可以运行丢信息的通道 |
| /charger/[id]/pub/1 | 1,2 | 支持 | 硬件上报优先级较高的通道，且qos>=1|

### charger的sub通道
| topic | qos | 支持retained | 说明 |
| --- | --- | --- |--- |
| /charger/[id]/sub/1 | 0,1,2 | 支持 | 硬件接受平台下发指令的通道，硬件在subscribe的时候需要设置qos=2|


