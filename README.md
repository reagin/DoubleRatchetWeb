## DoubleRatchetDemo

>   [!NOTE]
>
>   一个使用 Golang+React 开发的简易聊天软件，使用双棘轮密钥协商协议进行密钥协商
>
>   由于收发信息与确认机制仍存在问题，请尽量在双方在线时通信...

### 如何使用？

1.  克隆本仓库到本地，并搭建好NodeJS和Golang的环境

2.  使用VS Code打开文件`double-ratchet.code-workspace`，以启用工作区

3.  在`server/cconfig`目录下创建包名为`config`的文件，并写入以下内容即可运行后端程序
    ```go
    // 关于 MySQL 数据库的配置
    const (
    	DB_USER     string = "root"
    	DB_SECRET   string = "password"
    	DB_PORT     string = "3306"
    	DB_HOST     string = "localhost"
    	DB_DATABASE string = "database_name"
    )
    
    // 部署时使用 Nginx 反向代理
    const (
    	GIN_PORT   string = "8080"
    	GIN_SERVER string = "localhost"
    )
    
    // 关于 JWT 管理用户认证的密码
    const (
    	JWT_SERRET string = "password"
    )
    ```

4.  使用npm包管理器安装依赖，即可运行前端程序

    