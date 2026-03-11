<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:0f0c29,50:302b63,100:24243e&height=200&section=header&text=ZhiX%20Backend&fontSize=60&fontColor=ffffff&fontAlignY=38&desc=极志社区%20·%20后端引擎&descAlignY=58&descSize=20&animation=fadeIn" width="100%"/>

[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-Framework-008ECF?style=flat-square&logo=go&logoColor=white)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-4169E1?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-7-DC382D?style=flat-square&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)
[![JWT](https://img.shields.io/badge/JWT-Auth-000000?style=flat-square&logo=jsonwebtokens&logoColor=white)]()

<br/>

### 💡 想看看写这个的人是谁？

**→ [走，去作者主页看看](https://www.macfans.app/) ←**

*不只有代码，还有更多有意思的东西在那里等你。*

<sub>→ Curious about the author? [Visit Homepage](https://www.macfans.app/) ←</sub>

</div>

---

<div align="center">

*快 · 稳 · 轻 · 没有废话*

</div>

---

## 🧬 这是什么

> 跑在云上的 Go 后端服务，极志社区的数据中枢。  
> 用户、文章、收藏、支付——所有流转，全在这里。

---

## 🛠️ 技术栈

| | 技术 | 版本 |
|:---:|:---:|:---:|
| 🐹 | Go | 1.23 |
| 🌐 | Gin | latest |
| 🗄️ | PostgreSQL | 15 |
| ⚡ | Redis | 7 |
| 🔐 | JWT | — |
| 📦 | Docker | — |

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

**① 配置环境**

```bash
cp .env.example .env
# 填写你的数据库 & Redis 配置
```

**② 本地跑起来**

```bash
go mod download
go run main.go
# ✅ http://localhost:8080
```

**③ Docker 一键启动**

```bash
docker build -t zhix-backend .
docker run -p 8080:8080 --env-file .env zhix-backend
```

---

## 🔑 环境变量

```env
DB_HOST=localhost        # 数据库地址
DB_PORT=5432             # 数据库端口
DB_USER=postgres         # 用户名
DB_PASSWORD=your_pwd     # 密码
DB_NAME=zhix             # 数据库名
REDIS_HOST=localhost     # Redis 地址
REDIS_PORT=6379          # Redis 端口
JWT_SECRET=your-secret   # JWT 密钥
PORT=8080                # 服务端口
```

---

## 📡 API 接口

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

---

## 👥 用户成长体系

```
🧑 普通用户
新手用户 ──→ 普通用户 ──→ 活跃用户 ──→ 资深用户 ──→ 传奇用户
  Lv.1         Lv.2         Lv.3         Lv.4         Lv.5 ✨

✍️ 管理员
实习编辑 ──→ 编辑 ──→ 资深编辑 ──→ 高级编辑 ──→ 首席编辑
  Lv.1      Lv.2     Lv.3        Lv.4        Lv.5 👑
```

> 浏览、点赞、收藏、发布均可积分，等级自动计算。

---

## 🧪 测试

```bash
go test -v ./...
```

---

<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:24243e,50:302b63,100:0f0c29&height=120&section=footer" width="100%"/>

**如果这个项目对你有帮助，不妨去 [作者主页](https://www.macfans.app/) 逛逛 👀**

*也许你会发现更多有意思的东西。*

<sub>If you find this project helpful, feel free to visit the <a href="https://www.macfans.app/">Author's Homepage</a> for more. ✨</sub>

</div>
