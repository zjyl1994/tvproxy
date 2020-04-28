# TVProxy

适用于 Kodi 的 TVB 和 RTHK 直播源，可以部署在本机或者服务器上。

**注意：** 如果部署在服务器上请自行做好防护措施，TVB是走官方CDN，RTHK是通过程序代理的，所以会消耗大量流量。

## 安装方法

你可以通过Docker的方式安装，我已经写好了一个Dockerfile。使用时请注意正确设置`TVPROXY_BASE_URL`和`TVPROXY_LISTEN`这两个环境变量并且处理好自己的网络环境。

此处介绍直接安装并通过Systemd管理的方法。

```bash
git clone https://github.com/zjyl1994/tvproxy.git
cd tvproxy
go build
sudo cp tvproxy /usr/bin/tvproxy
sudo chmod +x /usr/bin/tvproxy
sudo cp example.env /etc/tvproxy.env
sudo cp tvproxy.service /etc/systemd/system/
sudo systemctl daemon-reload
```

然后，编辑`/etc/tvproxy.env`，常用的参数如下：

|环境变量|作用|
|---|---|
|TVPROXY_LISTEN|决定程序监听哪一个地址上的哪个端口，默认值为 "127.0.0.1:10086"|
|TVPROXY_BASE_URL|决定程序对外服务的地址前缀，默认值为 "http://127.0.0.1:10086/"|
|HTTP_PROXY|程序使用的代理服务器，代理RTHK时使用，无默认值，跟随系统设置|

现在有几个场景可供参考:（端口号任意，默认10086）

1. 本地环境使用：设置好`HTTP_PROXY`，然后直接观看，此时M3U地址为`http://127.0.0.1:10086/iptv.m3u`
2. 路由器使用：设置好`HTTP_PROXY`，`TVPROXY_LISTEN`填写“0.0.0.0:10086”，`TVPROXY_BASE_URL`填写“http://路由器ip:10086/”

## 使用方法

以本地使用(`http://127.0.0.1:10086/`)为例，下面是本软件支持的所有的直播源。

|频道|地址|
|---|---|
|Kodi PVR 使用|http://127.0.0.1:10086/iptv.m3u|
|無線新聞台|http://127.0.0.1:10086/tvb/inews.m3u8|
|無線財經資訊台|http://127.0.0.1:10086/tvb/finance.m3u8|
|RTHK 31|http://127.0.0.1:10086/rthk/31.m3u8|
|RTHK 32|http://127.0.0.1:10086/rthk/32.m3u8|

Kodi 的 PVR 客户端可以使用 http://127.0.0.1:10086/iptv.m3u 。

当你不使用 Kodi 只是想看某一个台的时候，你可以用VLC直接打开网络串流，填入相应的链接即可。


本项目为个人使用，不接受任何增加功能的需求issue。
当然如果你发现直播源挂了可以发issue，我会尽力修，但不提供修好保障。
直播源能不能用这个东西还是要看缘分。

