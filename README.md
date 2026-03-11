# ZhiX Backend

极志社区后端服务，基于 Go + Gin 构建的 RESTful API 服务。

> 作者主页：[macfans.app](https://www.macfans.app/)

---

> Author's Homepage: [macfans.app](https://www.macfans.app/)

## 技术栈

- **语言**: Go 1.23
- **框架**: Gin
- **ORM**: GORM
- **数据库**: PostgreSQL 15
- **缓存**: Redis 7
- **认证**: JWT
- **容器**: Docker

## 项目结构

```
backend/
├── config/          # 数据库、Redis 初始化
├── controllers/     # 业务逻辑处理
├── middleware/      # JWT 认证、权限中间件
├── models/          # 数据模型
├── routes/          # 路由注册
├── Dockerfile
└── main.go
```

## 快速开始

### 1. 配置环境变量

```bash
cp .env.example .env
# 编辑 .env 填写数据库等配置
```

### 2. 本地运行

```bash
go mod download
go run main.go
# 服务启动在 http://localhost:8080
```

### 3. Docker 运行

```bash
docker build -t zhix-backend .
docker run -p 8080:8080 --env-file .env zhix-backend
```

## 环境变量

| 变量 | 说明 | 示例 |
|------|------|------|
| `DB_HOST` | 数据库地址 | `localhost` |
| `DB_PORT` | 数据库端口 | `5432` |
| `DB_USER` | 数据库用户 | `postgres` |
| `DB_PASSWORD` | 数据库密码 | `password` |
| `DB_NAME` | 数据库名 | `zhix` |
| `REDIS_HOST` | Redis 地址 | `localhost` |
| `REDIS_PORT` | Redis 端口 | `6379` |
| `JWT_SECRET` | JWT 密钥 | `your-secret-key` |
| `PORT` | 服务端口 | `8080` |

## API 接口

### 认证

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/auth/register` | 注册 |
| POST | `/api/auth/login` | 登录 |

### 文章

| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| GET | `/api/articles` | 可选登录 | 文章列表 |
| GET | `/api/articles/homepage` | 公开 | 首页文章 |
| GET | `/api/articles/:id` | 公开 | 文章详情 |
| GET | `/api/articles/GetArticle/:l1/:l2/:l3` | 公开 | 按路径获取文章 |
| POST | `/api/articles` | 管理员 | 创建文章 |
| PUT | `/api/articles/:id` | 管理员 | 更新文章 |
| DELETE | `/api/articles/:id` | 管理员 | 删除文章 |
| POST | `/api/articles/:id/like` | 登录 | 点赞 |
| POST | `/api/articles/:id/view` | 公开 | 浏览 |
| POST | `/api/articles/:id/favorite` | 登录 | 收藏 |
| DELETE | `/api/articles/:id/favorite` | 登录 | 取消收藏 |
| GET | `/api/articles/:id/favorite/check` | 登录 | 检查收藏状态 |

### 用户

| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| PUT | `/api/user/avatar` | 登录 | 更新头像 |
| GET | `/api/stats/me` | 登录 | 个人统计 |

### 其他

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查 |
| GET | `/api/avatar/random` | 随机头像 |
| GET | `/api/cover/random` | 随机封面图 |
| GET | `/api/cover/cartoon` | 随机插画封面 |
| GET | `/api/cover/video` | 随机视频封面 |
| POST | `/api/payment/unionpay-h5` | 云闪付 H5 支付 |
| POST | `/api/transcode/mp4-to-fmp4` | 视频转码（管理员）|

## 用户角色

- `user`：普通用户，可点赞、收藏、浏览
- `admin`：管理员，可创建/编辑/删除文章

用户等级根据行为积分自动计算，分为 5 级（新手用户 → 传奇用户 / 实习编辑 → 首席编辑）。

## 测试

```bash
go test -v ./...
```
