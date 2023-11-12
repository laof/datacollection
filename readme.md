##### PE U Install
```
physdiskwrite.exe -u openwrt-22.03.3-x86-64-generic-ext4-combined-efi.img
```

##### Update login background

```
/www/luci-static/argon/img/bg1.jpg
```

##### Install git

```
opkg update
opkg install git git-http ca-bundle
```

##### Build image 
```
docker build -t laof/sharefile:0.1 .
```

##### Check image
```
docker ps // 运行状态
docker images // 镜像
docker image rm 4afgeqo4456 // 删除
```

##### Docker run 
```
docker run -d -p 6200:6200 --name sharefile
```

##### qinglong

```
docker run -dit \
-v $PWD/ql/config:/ql/config \
-v $PWD/ql/log:/ql/log \
-v $PWD/ql/db:/ql/db \
-p 5700:5700 \
--name qinglong \
--hostname qinglong \
--restart always \
whyour/qinglong:latest
```

##### 备份openwrt的所有用户数据
```
tar zcvf /tmp/backup.tar.gz /overlay/upper
```

##### 恢复openwrt的所有用户数据
```
tar x -zvC / -f /tmp/backup.tar.gz
```

##### 修改界面
```
/usr/lib/lua/luci/controller/admin/system.lua  // 系统菜单

/usr/lib/lua/luci/view/admin_system/  // xxx.htm文件
```
