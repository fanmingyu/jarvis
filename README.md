# 短信服务3.0服务端

[toc]

## 服务端，发送短信
### 环境搭建
1、安装Go https://golang.org/doc/install

2、golang.org被墙，使用代理 `export GOPROXY=https://goproxy.io`

### 部署服务

1、首先选好数据库，创建表（如果已经建好表就略过），利用数据库文件sms.sql，进入相关库下使用source sms.sql进行创建表(需要该文件在进入数据库时候所在的目录)

2、首先检查配置文件（根据不同的环境选择不同的配置文件），检查是否存在配置文件中所配置的目录路径，如果不存在，那么要创建相应的目录，配置文件中的参数举例

```go
{
    "LogPath": "./logs", //日志目录（需要存在该路径，没有就创建否则无法启动服务）
    "RedisSentinelMasterName": "def_master", //redis master的名字
    "RedisSentinelHosts": ["10.20.69.101:26479", "10.20.69.101:26579", "10.20.69.101:26679"], //redis哨兵的host
    "Mysql": {
        "Addr": "127.0.0.1:3306", //数据库的host和port组合（如果端口号为3306可以省略）
        "DBName": "sms", //数据库名称
        "User": "sms", //数据库连接账号名
        "Passwd": "smssms" //数据库连接密码
    },
    "Port": "10000" //服务监听的端口号(不要和其他服务冲突)
}
```
3、根目录下编译程序 make all #，编译完成后会生成对应的xxxx-console文件

4、检查完配置文件后，根据相应的配置文件启动服务，启动有两种方式：
①使用命令 nohup ./smsgate-server config.test.json & 来启动(其中的配置文件改为自己相应的配置文件)
②修改start.sh文件，将其中的配置文件改为自己相对应的配置文件，再运行脚本

5、一般而言，到此服务是可以启动的，可以进行进程搜索查看是否有进程存在。如果进程不存在，可以查看nohup.out文件来检查错误.

### 使用http接口来调用短信发送

```
POST /sms/send 或者 /send（兼容接口）
```

| 参数 | 说明 | 类型 | 必填 | 示例 |
| - | - | - | - | - |
| app | 站点名 | string | 是 | pcn |
| tpl | 模板名称 | string | 是 | TPL_SMS_VERIFY_CODE |
| mobile | 手机号 | string | 是 | 15900000000 |
| sign | md5的签名 | string | 是(当前不必须) | |
| vars | 模板中需要替换的文本 | string | 否 | ["123", "100元"] |

注：营销短信必须使用配置了营销短信通道的站点(app)来发送，如果是配置的营销短信通道，手机号可以用逗号隔开表示发送多条

```
POST /batch
```
注：此接口为兼容接口，或者为批量发送短信内容不一致的短信
| 参数 | 说明 | 类型 | 必填 | 示例 |
| - | - | - | - | - |
| app | 站点名 | string| 是 | pcn |
| data | json数组 | string | 是 | [{"mobile": "159xxxxxxx0", "content": "hello world 0..."}, {"mobile": "159xxxxxxx1", "content": "hello world 1..."}] |

**使用此接口的时候使用的站点名必须是配置为营销通道的站点**

### 通过http接口来添加和删除黑名单
添加接口
```
POST /blacklist/add
```
| 参数 | 说明 | 类型 | 必填 | 示例 |
| - | - | - | - | - |
| mobile | 手机号 | string | 是 | 159xxxxxxxx |

删除接口
```
POST /blacklist/remove
```
| 参数 | 说明 | 类型 | 必填 | 示例 |
| - | - | - | - | - |
| mobile | 手机号 | string | 是 | 159xxxxxxxx |


## 控制台，管理后台

控制台的主要作用是配置站点应用、模板，同时查询短信记录和添加删除黑名单

### 部署控制台服务

1、数据库建表，部署短信服务的时候如果建好了表就略过，否则参考短信服务的建表

2、创建目录smsgate-console作为目录，将压缩包smsgate-console解压缩，进入bin目录

3、首先检查配置文件，选取相应的配置文件，其中的配置文件中相应的参数举例如下

```go
{
    "LogPath": "./logs", //日志目录，相对路径是相对bin文件的路径
    "Mysql": {
        "Addr": "127.0.0.1:3306",
        "DBName": "sms",
        "User": "sms",
        "Passwd": "sms",
    },                   //数据库配置，和短信服务配置的数据库保持一致即可
    "Port": "10001"   //端口号，不能与其他服务重合
}
```

4、配置文件配置好以后就可以参考短信服务进行启动控制台服务了

### 后台地址
后台地址为： ip+port 例10.20.69.101:10001

### 使用docker-compose部署

linux内核需要大约3.1.0 ,下面是centos 7 下面部署的过程：


1、安装docker

```
sudo curl -sSL https://get.daocloud.io/docker | sh
```

2、安装 docker compose，资料：[install-compose](https://docs.docker.com/compose/install/#install-compose)

```
sudo curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
```
3、根据需求配置mysql、redis、golang、etcd的docker镜像，如镜像拉取缓慢可使用阿里云镜像加速

4、
构建镜像，直至正常运行
```
docker-compose up --build
```
5、启动镜像，并后台运行
```
docker-compose up -d
```
6、拉取代码至golang镜像的对应目录下，并进入golang容器
```
docker-compose exec golang bash
```
8、在golang容器中参照上述部署文档编译代码，并修改相应的连接配置

9、docker-compose.yml参考
```
version: "2"

services:
        # 这里的名字可以自定义，服务启动后以这里的名字来命名容器的名字
        db:
                # 这里是我们要引用的容器镜像，本地如果没有该镜像会自动从线上库拉取
                image: mariadb:10.1
                environment:
                        MYSQL_ROOT_PASSWORD: "root"
                        MYSQL_DATABASE: "app"
                        MYSQL_USER: "app"
                        MYSQL_PASSWORD: "123123"
                # 这里用来配置我们需要挂载的目录或文件，主要有配置文件、日志文件和代码
                volumes: 
                # ：前面是宿主机文件或目录的地址，：后面是容器中对应文件或目录所在地址，只需修改前面宿主机文件或目录地址即可
                - ./services/db/mysql/data:/var/lib/mysql
                - ./services/db/mysql/sql:/var/backups
                - ./services/db/mysql/config:/etc/mysql/conf.d

        golang:
                build: ./golang
                # 这里是我们要和宿主机进行映射的端口
                ports:
                        - "10001:10001"
                        - "10000:10000"
                links:
                        - "db"
                        - "redis"
                        - "etcd"
                volumes:
                        - ./app:/go
                tty: true  

        etcd:
                image: quay.io/coreos/etcd:v3.3.12
                volumes:
                        - ./services/etcd/data:/etcd-data
                command:
                        - "/usr/local/bin/etcd"
                        - "--name"
                        - "s1"
                        - "--data-dir"
                        - "/etcd-data"
                        - "--advertise-client-urls"
                        - "http://0.0.0.0:2379"
                        - --listen-client-urls
                        - "http://0.0.0.0:2379"
                        - "--initial-advertise-peer-urls"
                        - "http://0.0.0.0:2380"
                        - "--listen-peer-urls"
                        - "http://0.0.0.0:2380"
                        - "--initial-cluster-token"
                        - "tkn"
                        - "--initial-cluster"
                        - "s1=http://0.0.0.0:2380"
                        - "--initial-cluster-state"
                        - "new"   

        redis:
                image: redis:3.2.1
                ports:
                        - "16379:6379"      
# networks:

volumes:
        db:
                driver: local

```
10、关于代码更新发布
代码更新需进入相应容器重新编译go源码，kill掉相应的进程后重启即可，容器部署与实体服务器部署相同