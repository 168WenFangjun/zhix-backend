<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:000000,50:1a1a2e,100:16213e&height=220&section=header&text=ZhiX%20Backend&fontSize=72&fontColor=00d4ff&fontAlignY=40&desc=极志社区%20·%20后端引擎&descAlignY=62&descSize=22&animation=fadeIn" width="100%"/>

<br/>

[![Go](https://img.shields.io/badge/Go_1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)][go]
[![Gin](https://img.shields.io/badge/Gin-Framework-00B4D8?style=for-the-badge&logo=go&logoColor=white)][gin]
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL_15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)][postgres]
[![Redis](https://img.shields.io/badge/Redis_7-DC382D?style=for-the-badge&logo=redis&logoColor=white)][redis]
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)][docker]
[![JWT](https://img.shields.io/badge/JWT-Auth-black?style=for-the-badge&logo=jsonwebtokens)][jwt]

</div>

---

<div align="center">

### 🔥 写这个项目的人，比这个项目更有意思

> 不只有代码 · 还有你没见过的东西
>
> 好奇心是最好的入场券

**👉 [点这里，去作者主页逛逛][homepage-cn] 👈**

---

### 🔥 The person behind this project is more interesting than the project itself

> Not just code · There's something here you haven't seen before
>
> Curiosity is the best ticket in

**👉 [Click here to visit the author's homepage][homepage-en] 👈**

</div>

---

## 🧬 这是什么

> 一个跑在云上的 Go 后端服务。
> 极志社区的数据中枢——用户、文章、收藏、支付，全在这里流转。
> **快 · 稳 · 轻 · 没有废话。**

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
DB_HOST=localhost          # 数据库地址
DB_PORT=5432               # 数据库端口
DB_USER=<db_user>          # 用户名
DB_PASSWORD=<db_password>  # 密码
DB_NAME=zhix               # 数据库名
REDIS_HOST=localhost       # Redis 地址
REDIS_PORT=6379            # Redis 端口
JWT_SECRET=<jwt_secret>    # JWT 密钥
PORT=8080                  # 服务端口
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

<table>
<tr>
<td align="center">

**看到这里了，说明你是认真的。**

认真的人，值得去认识认真的作者 →

**🌐 [前往作者主页][homepage-cn]**

</td>
<td align="center">

**You made it to the end — respect.**

Serious devs deserve to meet a serious author →

**🌐 [Visit Author's Homepage][homepage-en]**

</td>
</tr>
</table>

<br/>

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:16213e,50:1a1a2e,100:000000&height=120&section=footer" width="100%"/>

</div>

[homepage-cn]: https://www.macfans.app/
[homepage-en]: https://www.macfans.app/
[go]: https://golang.org/
[gin]: https://gin-gonic.com/
[postgres]: https://www.postgresql.org/
[redis]: https://redis.io/
[docker]: https://www.docker.com/
[jwt]: https://jwt.io/
