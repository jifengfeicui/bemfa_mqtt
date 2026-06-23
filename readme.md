# bemfa_mqtt

通过巴法云 MQTT 接入米家/小爱，用 Ubuntu 上的程序控制 Windows 电脑：

- `pc001`：WOL 远程唤醒
- `screen006`：006 开关，`on` 旋转屏幕到 90 度，`off` 恢复到 0 度

## 部署文件

Ubuntu 上部署同一目录放这几个文件：

```text
wol
config.ini
tmp/rotatescreen.ps1
```

启动：

```bash
chmod +x wol
./wol
```

程序启动时从当前目录读取 `config.ini`，日志写到 `stdout.log`。

## 巴法云配置

1. 注册巴法云账号并获取私钥：`bemfa_client_id`
2. 创建 MQTT 主题：
   - WOL 唤醒可用 `pc001`
   - 屏幕旋转用 `screen006`
3. 在米家 App 中添加巴法云设备。

主题后三位代表设备类型，`006` 是开关，所以屏幕旋转用 `screen006`。

## config.ini 示例

```ini
[DEFAULT]
bemfa_broker = bemfa.com
bemfa_port = 9501
bemfa_client_id = xxxxxxxxxx

[pc001]
struct = wol
broadcast = 192.168.x.255
mac = xx:xx:xx:xx:xx:xx
ip = 192.168.xx.xx
user = xx
password =

[screen006]
struct = screen_rotate
ssh_host = 192.168.x.x
ssh_port = 22
ssh_user = user
ssh_password = password
display = 2
script = tmp/rotatescreen.ps1
```

## 屏幕旋转

`screen006` 收到巴法云消息后通过 SSH 远程执行 Windows PowerShell：

- `on`：旋转到 90 度
- `off`：旋转到 0 度

Windows 电脑需要开启 OpenSSH Server，`ssh_user` 对应账号需要能执行 PowerShell。

配置项：

- `ssh_host`：Windows 电脑 IP
- `ssh_port`：SSH 端口，默认 `22`
- `ssh_user`：Windows SSH 用户名
- `ssh_password`：Windows SSH 密码
- `display`：屏幕编号，默认 `2`，对应 `\\.\DISPLAY2`
- `script`：本机脚本路径，默认 `tmp/rotatescreen.ps1`

## WOL 唤醒

WOL 需要在 BIOS 和网卡驱动中打开相关选项：

- Wake on LAN
- Wake-on-LAN from S4/S5
- Power on by PCIe devices
- 允许此设备唤醒计算机
- 唤醒魔包

如果 Windows 开启了快速启动，建议关闭，否则可能影响关机后的 WOL 唤醒。

## 远程关机

Linux 下远程关机依赖 `net rpc shutdown`：

```bash
sudo apt install samba-common-bin
```

Windows 目标机还需要给对应用户开启“从远程系统强制关机”权限。
