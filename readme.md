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









