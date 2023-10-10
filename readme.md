## [功能]

通过巴法云，使用小爱同学远程唤醒电脑WOL。

## [使用]

1. 注册巴法云账号，并获取密钥
2. 在HACS中搜索 bemfa 安装，或者 clone 此项目, 将 custom_components/bemfa 目录拷贝至 Home Assistant 配置目录的 custom_components 目录下。
3. 重启 Home Assistant 服务。
4. 在 Home Assistant 的集成页面，搜索 "bemfa" 并添加。
5. 根据提示输入巴法云密钥后提交
6. 安装成功后，点击集成左下角“选项”，同步需要的实体至巴法云。
7. 在智能音箱App中添加巴法云设备:
   - 小爱同学: 在米家app-->我的-->其他平台设备-->点击添加-->找到"巴法"，输入巴法云账号即可，设备会自动同步到米家。
   - 天猫精灵: 打开天猫精灵app，在app中搜索：巴法云。找到巴法云技能，点击绑定账号，登陆你的巴法云账号.
   - 小度音箱: 打开小度音箱app或者小度app，在app首页点+号-->添加设备-->搜索巴法，找到"巴法"，输入巴法云账号即可。