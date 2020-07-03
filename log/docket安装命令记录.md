# docket安装命令记录

```
yum update
```

软件包安装
```
yum install -y yum-utils  device-mapper-persistent-data lvm2
```
添加yum源
```
yum-config-manager \
--add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```
查看可安装的版本
```
yum list docker-ce --showduplicates | sort -r
```
安装最新版本
```
yum install docker-ce -y
```
启动并开机自启
```
systemctl start docker
systemctl enable docker
```
查看docker版本
```
docker version 
```