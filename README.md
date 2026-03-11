<div align="center">

# ⚡ ZhiX Backend

**极志社区 · 后端引擎**

*Go 写的不只是代码，是态度。*

[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-Framework-008ECF?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

---

🏠 作者主页：[点击访问](https://www.macfans.app/) &nbsp;·&nbsp; Author's Homepage: [Visit Here](https://www.macfans.app/)

</div>

---

## 🧬 这是什么

> 一个跑在云上的 Go 后端服务。  
> 负责极志社区的全部数据流转——用户、文章、收藏、支付，一个不落。  
> 快、稳、轻。没有废话。

---

## 🛠️ 技术栈

| | 技术 | 版本 |
|---|---|---|
| 🐹 语言 | Go | 1.23 |
| 🌐 框架 | Gin | latest |
| 🗄️ 数据库 | PostgreSQL | 15 |
| ⚡ 缓存 | Redis | 7 |
| 🔐 认证 | JWT | — |
| 📦 容器 | Docker | — |

---

## 📂 项目结构

```
backend/
├── config/          # 数据库 & Redis 初始化
├── controllers/     # 核心业务逻辑
├── middleware/      # JWT 认证 & 权限拦截
├── models/          # 数据模型定义
├── routes/          # 路由注册
├── Dockerfile       # 容器化配置
└── main.go          # 程序入口
```

---

## 🚀 快速启动

**第一步：配置环境**

```bash
cp .env.example .env
# 填写你的数据库 & Redis 配置
```

**第二步：跑起来**

```bash
go mod download
go run main.go
# ✅ 服务运行在 http://localhost:8080
```

**或者用 Docker**

```bash
docker build -t zhix-backend .
docker run -p 8080:8080 --env-file .env zhix-backend
```

---

## 🔑 环境变量

```env
# 数据库
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=zhix

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# 认证
JWT_SECRET=your-secret-key

# 服务
PORT=8080
```

---

## 📡 API 一览

<details>
<summary><b>🔐 认证</b></summary>

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/auth/register` | 注册 |
| `POST` | `/api/auth/login` | 登录 |

</details>

<details>
<summary><b>📝 文章</b></summary>

| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| `GET` | `/api/articles` | 可选登录 | 文章列表 |
| `GET` | `/api/articles/homepage` | 公开 | 首页文章 |
| `GET` | `/api/articles/:id` | 公开 | 文章详情 |
| `GET` | `/api/articles/GetArticle/:l1/:l2/:l3` | 公开 | 按路径获取 |
| `POST` | `/api/articles` | 🔒 管理员 | 创建文章 |
| `PUT` | `/api/articles/:id` | 🔒 管理员 | 更新文章 |
| `DELETE` | `/api/articles/:id` | 🔒 管理员 | 删除文章 |
| `POST` | `/api/articles/:id/like` | 登录 | 点赞 |
| `POST` | `/api/articles/:id/view` | 公开 | 浏览 |
| `POST` | `/api/articles/:id/favorite` | 登录 | 收藏 |
| `DELETE` | `/api/articles/:id/favorite` | 登录 | 取消收藏 |
| `GET` | `/api/articles/:id/favorite/check` | 登录 | 收藏状态 |

</details>

<details>
<summary><b>👤 用户 & 统计</b></summary>

| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| `PUT` | `/api/user/avatar` | 登录 | 更新头像 |
| `GET` | `/api/stats/me` | 登录 | 个人统计 |

</details>

<details>
<summary><b>🎨 素材 & 其他</b></summary>

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/health` | 健康检查 |
| `GET` | `/api/avatar/random` | 随机头像 |
| `GET` | `/api/cover/random` | 随机封面图 |
| `GET` | `/api/cover/cartoon` | 随机插画封面 |
| `GET` | `/api/cover/video` | 随机视频封面 |
| `POST` | `/api/payment/unionpay-h5` | 云闪付 H5 支付 |
| `POST` | `/api/transcode/mp4-to-fmp4` | 视频转码（管理员）|

</details>

---

## 👥 用户体系

```
普通用户成长路径：
新手用户  →  普通用户  →  活跃用户  →  资深用户  →  传奇用户
  Lv.1        Lv.2        Lv.3        Lv.4        Lv.5

管理员成长路径：
实习编辑  →  编辑  →  资深编辑  →  高级编辑  →  首席编辑
  Lv.1      Lv.2     Lv.3        Lv.4        Lv.5
```

> 等级由行为积分自动计算，浏览、点赞、收藏、发布均可获得积分。

---

## 🧪 测试

```bash
go test -v ./...
```

---

<div align="center">

**Built with ❤️ · Powered by Go · Made for ZhiX**

</div>
