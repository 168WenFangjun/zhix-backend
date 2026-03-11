<div align="center">

```
███████╗██╗  ██╗██╗██╗  ██╗
╚══███╔╝██║  ██║██║╚██╗██╔╝
  ███╔╝ ███████║██║ ╚███╔╝
 ███╔╝  ██╔══██║██║ ██╔██╗
███████╗██║  ██║██║██╔╝ ██╗
╚══════╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝
```

### 极志社区 · 后端

**数据在这里流转，逻辑在这里生长。**

---

[![作者主页](https://img.shields.io/badge/🔥_作者是谁？点进来就知道了-→-FF3B30?style=for-the-badge)](https://www.macfans.app/)
&nbsp;
[![Author](https://img.shields.io/badge/🌐_Who_built_this%3F_Find_out_→-6C63FF?style=for-the-badge)](https://www.macfans.app/)

---

</div>

---

## 🇨🇳 中文版

<br>

<div align="center">

**不是所有后端都叫极志。**

Go 语言驱动，AWS 云端部署，为高并发而生。  
快 · 稳 · 轻 · 没有废话。

</div>

<br>

### 🧱 技术栈

<div align="center">

![Go](https://img.shields.io/badge/Go_1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00B4D8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL_15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis_7-DC382D?style=for-the-badge&logo=redis&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-Auth-black?style=for-the-badge&logo=jsonwebtokens)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![AWS ECS](https://img.shields.io/badge/AWS_ECS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)

</div>

<br>

### 📁 项目结构

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

<br>

### ⚡ 三步跑起来

```bash
cp .env.example .env
go mod download
go run main.go
```

`.env` 配置：

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=<db_user>
DB_PASSWORD=<db_password>
DB_NAME=zhix
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=<jwt_secret>
PORT=8080
```

> 启动后访问 → `http://localhost:8080`

**Docker 一键启动：**

```bash
docker build -t zhix-backend .
docker run -p 8080:8080 --env-file .env zhix-backend
```

<br>

### 🎯 能干什么

```
✦ RESTful API，接口清晰规范
✦ JWT 身份认证，三级权限体系：管理员 · 用户 · 游客
✦ 文章增删改查，点赞 · 收藏 · 浏览全支持
✦ 用户成长体系，等级自动计算
✦ Redis 缓存加速，高并发不慌
✦ 云闪付 H5 支付集成
✦ 视频转码支持
```

<br>

### 📡 API 接口

<details>
<summary><b>🔐 认证</b></summary>
<br/>

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/auth/register` | 注册 |
| `POST` | `/api/auth/login` | 登录 |

</details>

<details>
<summary><b>📝 文章</b></summary>
<br/>

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
<br/>

| 方法 | 路径 | 权限 | 说明 |
|------|------|------|------|
| `PUT` | `/api/user/avatar` | 登录 | 更新头像 |
| `GET` | `/api/stats/me` | 登录 | 个人统计 |

</details>

<details>
<summary><b>🎨 素材 & 其他</b></summary>
<br/>

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

<br>

### 👥 用户成长体系

```
🧑 普通用户
新手用户 ──→ 普通用户 ──→ 活跃用户 ──→ 资深用户 ──→ 传奇用户
  Lv.1         Lv.2         Lv.3         Lv.4         Lv.5 ✨

✍️ 管理员
实习编辑 ──→ 编辑 ──→ 资深编辑 ──→ 高级编辑 ──→ 首席编辑
  Lv.1      Lv.2     Lv.3        Lv.4        Lv.5 👑
```

> 浏览、点赞、收藏、发布均可积分，等级自动计算。

<br>

### 🚀 部署

通过 GitHub Actions 自动构建，推送即部署，详见 [Backend 部署指南](../docs/项目backend部署指南1.0版.md)。

<br>

---

<div align="center">

**这个项目背后的人，比你想象的更有意思。**

[![👀 去看看作者在搞什么](https://img.shields.io/badge/👀_去看看作者在搞什么_→_macfans.app-FF3B30?style=for-the-badge)](https://www.macfans.app/)

</div>

---

<br>

## 🇺🇸 English Version

<br>

<div align="center">

**Not just another backend. This one means business.**

Go · AWS · Built for performance and scale.

</div>

<br>

### 🧱 Tech Stack

<div align="center">

![Go](https://img.shields.io/badge/Go_1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00B4D8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL_15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis_7-DC382D?style=for-the-badge&logo=redis&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-Auth-black?style=for-the-badge&logo=jsonwebtokens)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![AWS ECS](https://img.shields.io/badge/AWS_ECS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)

</div>

<br>

### 📁 Project Structure

```
backend/
├── config/          # DB & Redis initialization
├── controllers/     # Core business logic
├── middleware/      # JWT auth & permission guards
├── models/          # Data model definitions
├── routes/          # Route registration
├── Dockerfile       # Container config
└── main.go          # Entry point
```

<br>

### ⚡ Up in 3 Steps

```bash
cp .env.example .env
go mod download
go run main.go
```

`.env` config:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=<db_user>
DB_PASSWORD=<db_password>
DB_NAME=zhix
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=<jwt_secret>
PORT=8080
```

> Runs at → `http://localhost:8080`

**Docker one-liner:**

```bash
docker build -t zhix-backend .
docker run -p 8080:8080 --env-file .env zhix-backend
```

<br>

### 🎯 What It Does

```
✦ Clean RESTful API design
✦ JWT authentication with 3-tier roles: Admin · User · Guest
✦ Full article CRUD — like, bookmark, view all supported
✦ User growth system with auto-calculated levels
✦ Redis caching for high-concurrency workloads
✦ UnionPay H5 payment integration
✦ Video transcoding support
```

<br>

### 📡 API Reference

<details>
<summary><b>🔐 Auth</b></summary>
<br/>

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/auth/register` | Register |
| `POST` | `/api/auth/login` | Login |

</details>

<details>
<summary><b>📝 Articles</b></summary>
<br/>

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| `GET` | `/api/articles` | Optional | Article list |
| `GET` | `/api/articles/homepage` | Public | Homepage articles |
| `GET` | `/api/articles/:id` | Public | Article detail |
| `GET` | `/api/articles/GetArticle/:l1/:l2/:l3` | Public | Get by path |
| `POST` | `/api/articles` | 🔒 Admin | Create article |
| `PUT` | `/api/articles/:id` | 🔒 Admin | Update article |
| `DELETE` | `/api/articles/:id` | 🔒 Admin | Delete article |
| `POST` | `/api/articles/:id/like` | Login | Like |
| `POST` | `/api/articles/:id/view` | Public | View |
| `POST` | `/api/articles/:id/favorite` | Login | Bookmark |
| `DELETE` | `/api/articles/:id/favorite` | Login | Remove bookmark |
| `GET` | `/api/articles/:id/favorite/check` | Login | Bookmark status |

</details>

<details>
<summary><b>👤 User & Stats</b></summary>
<br/>

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| `PUT` | `/api/user/avatar` | Login | Update avatar |
| `GET` | `/api/stats/me` | Login | Personal stats |

</details>

<details>
<summary><b>🎨 Assets & Misc</b></summary>
<br/>

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/api/avatar/random` | Random avatar |
| `GET` | `/api/cover/random` | Random cover image |
| `GET` | `/api/cover/cartoon` | Random illustration cover |
| `GET` | `/api/cover/video` | Random video cover |
| `POST` | `/api/payment/unionpay-h5` | UnionPay H5 payment |
| `POST` | `/api/transcode/mp4-to-fmp4` | Video transcode (Admin) |

</details>

<br>

### 👥 User Growth System

```
🧑 Regular Users
Newcomer ──→ Member ──→ Active ──→ Veteran ──→ Legend
  Lv.1        Lv.2      Lv.3       Lv.4        Lv.5 ✨

✍️ Admins
Intern ──→ Editor ──→ Senior Editor ──→ Lead Editor ──→ Chief Editor
  Lv.1      Lv.2         Lv.3              Lv.4            Lv.5 👑
```

> Points earned from views, likes, bookmarks, and posts. Levels calculated automatically.

<br>

### 🚀 Deployment

Auto-deployed via GitHub Actions on every push. See [Backend Deployment Guide](../docs/项目backend部署指南1.0版.md).

<br>

---

<div align="center">

**The person behind this project is worth knowing.**

[![🔗 Step Into the Author's World →](https://img.shields.io/badge/🔗_Step_Into_the_Author's_World_→-6C63FF?style=for-the-badge)](https://www.macfans.app/)

</div>

---

<div align="center">
<br>

`// made with focus · ZhiX Team`

</div>
