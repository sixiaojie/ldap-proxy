# ldap_proxy
主要用于验证的中间状态。 一般需要结合nginx中auth_request 模块一起使用
 
## 实现方式：

<h5> client -> nginx ->（auth）-> backend

nginx auth module模块： 一般需要安装nginx auth module（默认安装自带）模块即可.

目前ldap_proxy 只支持ldap的验证。

## 配置

配置文件config.yaml:

expire: 过期时间

store: 存储登陆状态。 目前支持redis、内存

ldap： ldap 连接服务的账户信息。



## 用例

1、nginx.conf 复制到目的nginx下，修改对应到upstream服务

2、编译 go build -o ldap-proxy . (不同系统自行选择不同编译方式)

3、将编译后的文件放置到执行目录下。并在执行文件当前目前下创建config、statis目录，并将代码中的配置和代码复制过对应的目录。

4、修改config.yml中对应的参数，使用方式

5 启动即可。（正常访问nginx的服务，会自动跳转到认证页面）。


不足： 目前仅支持ldap，但主流的oauth 认证 还在研究中 哈哈哈。 

如果后端服务需要从代码中获取到用户信息，可以直接在后端代码中，加入对header authorization的抓取。就可以获取到用户的信息呢
