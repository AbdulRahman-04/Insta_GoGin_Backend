# 📸 Insta\_GoGin\_Backend

**A blazing-fast, scalable, authenticated backend built with Go, Gin, MongoDB, and JWT — complete with role-based access control, file uploads, and more!**

---

## 🧹 Table of Contents

1. [🚀 Features](#-features)
2. [⚙️ Tech Stack](#️-tech-stack)
3. [💪 Getting Started](#️-getting-started)
4. [📁 API Endpoints](#-api-endpoints)
5. [🔐 Authentication & Roles](#-authentication--roles)
6. [📆 File Uploads & Logging](#-file-uploads--logging)
7. [📁 Folder Structure](#-folder-structure)
8. [✅ Roadmap & Future Plans](#✅-roadmap–future-plans)
9. [🙌 Contributing](#-contributing)
10. [📄 License](#-license)

---

## 🚀 Features

* **Fast, concurrent RESTful API** using Go + Gin
* **User & Admin** sign-up / login / password change / email & phone verification
* **JWT authentication** with middleware-protected private routes
* **Role-Based Access**: only users can create/edit/delete their own posts/stories
* **Secure password reset** via email
* **File upload support** for images/files
* **Detailed logging** for terminal-friendly dev insights

---

## ⚙️ Tech Stack

* Language: `Go 1.21+`
* Web Framework: [Gin](https://github.com/gin-gonic/gin)
* Database: [MongoDB](https://www.mongodb.com/)
* Authentication: JWT (`Authorization: Bearer <token>`)
* Email sending: SMTP (using utils)
* Logging: Custom structured logs
* Image/File Handling: Multipart + S3/local (via utils)

---

## 💪 Getting Started

### 1. Clone & Setup

```bash
git clone https://github.com/AbdulRahman-04/Insta_GoGin_Backend.git
cd Insta_GoGin_Backend
```

### 2. Set Environment Variables

Create a `.env`:

```
MONGO_URI=<your_mongo_conn_string>
JWTKEY=<your_jwt_secret>
URL=<your_app_base_url>
SMTP_HOST=<...>
SMTP_EMAIL=<...>
SMTP_PASS=<...>
```

### 3. Install & Run

```bash
go mod download

go run main.go
```

→ Server runs at `http://localhost:6060`

---

## 📁 API Endpoints

### **Public Routes (No token required)**

| Route                                 | Method | Description                                 |
| ------------------------------------- | ------ | ------------------------------------------- |
| `/api/public/user/register`           | POST   | User signup                                 |
| `/api/public/user/login`              | POST   | User login                                  |
| `/api/public/user/emailverify/:token` | GET    | Email verification                          |
| `/api/public/user/change-password`    | POST   | Change password                             |
| `/api/public/user/forgot-password`    | POST   | Send temporary password via email           |
| \[Admin versions available]           |        | Similar signup/login routes with admin role |
| `/api/public/ping`                    | GET    | Health check (`{"msg": "✅ API is live!"}`)  |

---

### **Private Routes (JWT required)**

* **Posts (User-only):**
  `addpost`, `getallposts`, `getonepost/:id`, `editpost/:id`, `deletepost/:id`, `deleteallposts`

* **Stories (User-only):**
  `addstory`, `getallstories`, `getonestories/:id`, `editonestories/:id`, `deleteonestories/:id`, `deleteallstories`

* **User Management (Admin-only):**
  `GET /users`, `GET /users/:id`

---

## 🔐 Authentication & Roles

* JWTs provide `id`, `role`, and `email` claims
* Middleware enforces:

  * `middleware.AuthMiddleware()` ➔ token presence
  * `middleware.OnlyUser()` ➔ user-level routes
  * `middleware.OnlyAdmin()` ➔ admin-level routes

---

## 📆 File Uploads & Logging

* Images/files supported via `multipart/form-data`
* Upload helper: `utils.UploadFile(c)`
* Logging: Custom log structure with timestamps built into controllers

---

## 📁 Folder Structure

```
├── controllers
│   ├── public   # Signup, login, verification
│   └── private  # Posts & Stories CRUD
├── middleware   # Auth & role-based access
├── models       # Post, Story, User data types
├── routes       # Public & private router setup
├── utils        # Mongo client, file uploader, email sender
└── main.go      # Application entrypoint
```

---

## ✅ Roadmap & Future Plans

* ⚡ Will be Adding **Redis caching** for hot content
* 🧌 Implementing **pagination & sorting**
* 🧑‍💻 Add **rate limiting** & **input sanitization**
* 🛡️ Improve **error handling** & **structured logging**
* 📝 Add **unit/integration tests**

---

## 🙌 Contributing

1. Fork & clone the repo
2. Create a feature branch (`git checkout -b feat/your-feature`)
3. Commit your changes (`git commit -am 'Your feature added'`)
4. Push & raise a Pull Request
5. Review & discuss; merge when approved!

---

## This Backend Project is solely made by Syed Abdul Rahman

---

**Let’s build clean, fast, and scalable Go backends — one PR at a time!** 🚀
