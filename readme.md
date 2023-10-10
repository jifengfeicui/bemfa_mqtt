## [功能]

使用小爱同学做电脑WOL远程唤醒，通过巴法云MQTT接入

## [使用]

1. 注册巴法云账号https://cloud.bemfa.com/，并获取密钥

2. 新建MQTT主题，例：ethanpc001
   注意：[主题名称的后三位必须是001-009](https://cloud.bemfa.com/docs/#/?id=p-stylefont-weight700margin0px11%e3%80%81%e7%b1%b3%e5%ae%b6%e5%b0%8f%e7%88%b1%e6%94%af%e6%8c%81)，前面随意，不同的编号在米家代表不同的设备。

   > 当主题名字后三位是001时为插座设备。

   > 当主题名字后三位是002时为灯泡设备。

   > 当主题名字后三位是003时为风扇设备。

   > 当主题名字后三位是004时为传感器设备。

   > 当主题名字后三位是005时为空调设备。

   > 当主题名字后三位是006时为开关设备。

   > 当主题名字后三位是009时为窗帘设备。

   

3. 在智能音箱App中添加巴法云设备:

   - 小爱同学: 在米家app-->我的-->其他平台设备-->点击添加-->找到"巴法"，输入巴法云账号即可，设备会自动同步到米家。

4. 在小爱训练-->个人训练-->添加，设置如下：

   <img src=".\picture.jpg" alt="截图" style="zoom:30%;" />

## [配置文件示例]

```
[DEFAULT]
bemfa_broker = bemfa.com
bemfa_port = 9501
# 巴法平台控制台获取的私钥
bemfa_client_id = xxxxxxxx

# 巴法云主题名称 例:ethanpc001
[ethanpc001]
# 广播地址
broadcast = 192.168.1.255 
# 设备mac地址
mac = 00:D8:61:78:4B:13
struct = wol





```



## [远程唤醒 WOL配置]

1. 首先到 BIOS 中打开 WOL 相关开关

   通常 WOL 相关设置会有下面的名称（仅供参考，以主板用户手册为准）：

   - Wake up on LAN
   - Wake-on-LAN from S4/S5
   - Power on by PCIe devices
   - Resume On LAN

2. 进入系统，按 Win+R 打开运行，输入 `devmgmt.msc` 打开设备管理器：

   ![img](https://doc.natfrp.com/assets/wol-1-998d1bd1.png)

3. 找到要用作 WOL 的网卡，双击打开网卡配置界面：

   ![img](https://doc.natfrp.com/assets/wol-2-cbaf0503.png)

   ::: info 不同网卡配置方式可能不一样，如果遇到配置困难请借助百度解决 :::

4. 点击 **高级**，找到 **唤醒模式匹配** 和 **唤醒魔包**，设置为 `启用`：

   ![img](https://doc.natfrp.com/assets/wol-3-690d24a1.png)

5. 点击 **电源管理**，找到 **允许此设备唤醒计算机**，勾上前面的复选框：

   提示

   下面的 **只允许幻数据包唤醒计算机** 建议也勾上，可以避免部分情况下计算机无故开机

   ![img](https://doc.natfrp.com/assets/wol-4-e1910907.png)

6. 如果计算机开启了快速启动，建议关闭快速启动，否则可能造成无法正常唤醒。