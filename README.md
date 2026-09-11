# goweb

`goweb` 是一个使用 Go 编写的社区 Web 应用示例。它提供用户注册与登录、社区查询、帖子发布与查询、帖子投票等功能，并包含静态前端页面、Swagger API 文档、结构化日志和优雅停机能力。

项目采用分层单体结构：

```text
router / middleware → controller → logic → MySQL / Redis DAO
```

- MySQL 保存用户、社区和帖子。
- Redis 保存帖子发布时间、排序分数和用户投票状态。
- JWT 用于接口认证，Snowflake 用于生成用户和帖子 ID。

## 技术栈

- Go 1.25
- Gin：HTTP 路由与中间件
- GORM + MySQL：业务数据持久化
- go-redis + Redis：帖子分数与投票状态
- JWT：用户认证
- Viper：配置文件与环境变量
- Zap + Lumberjack：结构化日志与日志轮转
- Swaggo：Swagger 文档
- Docker / Docker Compose：本地容器化运行

## 核心功能

- 用户注册、登录和 JWT 身份认证
- 社区列表与社区详情查询
- 帖子发布、分页列表和详情查询
- 帖子赞成、反对和取消投票
- 7 天投票窗口及 Redis 原子投票更新
- 健康检查、Swagger 文档和开发环境 pprof

主要接口定义可在启动后的 Swagger 页面查看，也可以直接阅读 [`router/router.go`](router/router.go) 和 [`docs/swagger.yaml`](docs/swagger.yaml)。

## 如何运行

### 使用 Docker Compose（推荐）

需要安装 Docker Desktop 或其他支持 Docker Compose 的环境。

```bash
docker compose up --build
```

服务启动后：

- Web 页面：<http://127.0.0.1:8887>
- 健康检查：<http://127.0.0.1:8887/ping>
- Swagger：<http://127.0.0.1:8887/swagger/index.html>

停止并删除本项目容器：

```bash
docker compose down
```

Compose 会启动 MySQL、Redis 和应用容器，并在两个数据服务通过健康检查后启动应用。数据库表由 GORM 自动创建；仓库暂未提供社区种子数据。

### 本地直接运行

准备以下环境：

- Go 1.25 或兼容版本
- MySQL，默认地址 `127.0.0.1:3306`，数据库名 `goweb1`
- Redis，默认地址 `127.0.0.1:6379`

配置位于 [`conf/config.yaml`](conf/config.yaml)，也可以使用环境变量覆盖，例如 `MYSQL_HOST`、`MYSQL_PASSWORD`、`REDIS_HOST` 和 `APP_PORT`。

JWT 密钥不会使用仓库内默认值。启动前必须设置至少 32 字节的 `AUTH_JWT_SECRET`：

PowerShell：

```powershell
$env:AUTH_JWT_SECRET = "replace-with-a-random-secret-of-32-bytes-or-more"
go run .
```

Bash：

```bash
export AUTH_JWT_SECRET='replace-with-a-random-secret-of-32-bytes-or-more'
go run .
```

本地默认监听 `http://127.0.0.1:8888`。可用以下请求确认服务正常：

```bash
curl http://127.0.0.1:8888/ping
```

预期结果：

```json
{"message":"pong"}
```

## 测试与检查

```bash
go test ./...
go vet ./...
docker compose config --quiet
```



