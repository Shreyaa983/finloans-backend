# finloans-backend
Finloans is a production-style backend system built with Golang, designed to manage user authentication, role-based workflows, and financial eligibility modules. 
---

## 📌 Features

- ✅ JWT-based authentication (Login, Register)
- ✅ Role-based authorization (Admin vs User)
- ✅ User profile management (view/update)
- ✅ Loan eligibility check (mock logic)
- ✅ Loan application & status APIs
- ✅ PDF generation for loan sanction letter
- ✅ Swagger-based auto API documentation
- ✅ GitHub Actions pipeline for build

---

## 🧱 Tech Stack

| Layer         | Tools Used |
|---------------|------------|
| Language      | Golang (`gin-gonic`, `gorm`) |
| Authentication| JWT (`github.com/dgrijalva/jwt-go`) |
| Database      | MySQL (or SQLite for local dev) |
| Docs          | Swagger (`swaggo/gin-swagger`) |

---

## 👩‍💼 For the Recruiter

This project is a **self-built demonstration** of the backend API systems I developed during my internship at **Fundsmama(FinTech)**.

It reflects the kind of work I contributed to in a real-world environment, including:

- 🔐 **Secure authentication** (JWT-based)
- 👤 **User profile and password management**
- 🔐 **Role-based access control**
- 🧠 **Loan eligibility logic-basic**
- 📄 **PDF document generation**
- 🧩 **Modular routing**

> **Note:** No proprietary code or company-specific logic has been used.  
> This project is built entirely from scratch to showcase my backend development skills and my familiarity with systems in a FinTech context.

---

## 📁 Project Structure

```
finloans-backend/
├── controllers/       # Route handler logic
├── middlewares/       # Auth & other middleware
├── models/            # Database models
├── public/            # Static files (PDFs, etc.)
├── routes/            # Route definitions
├── utils/             # Helper utilities
├── main.go            # Entry point
└── go.mod             # Go module definitions
```

---
## 🔐 API Endpoints

| Method | Route                              | Description                        |
|--------|------------------------------------|------------------------------------|
| POST   | `/api/register`                    | Register a new user                |
| POST   | `/api/login`                       | Login and receive JWT              |
| GET    | `/api/auth/get-profile`            | Get current user's profile         |
| POST   | `/api/auth/check-eligibility`      | Run loan eligibility logic         |
| POST   | `/api/auth/apply-loan`             | Submit loan application            |
| GET    | `/api/auth/my-loans`               | Get current user's loan list       |
| GET    | `/swagger/index.html`              | View Swagger docs                  |


---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/Shreyaa983/finloans-backend.git
cd finloans-backend
````

### 2. Configure Environment

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=finloans_db
JWT_SECRET=your_super_secret_key
```

### 3. Run Locally

```bash
go mod tidy
go run main.go
```

### 4. Swagger Docs

Visit: `http://localhost:8080/swagger/index.html`

---

## 🔄 CI/CD – GitHub Actions

* Auto builds and tests on every push to `main`

---

## 📚 Learning Goals / Resume Highlights

* Golang backend project with clean modular design
* JWT Auth, Role-based Access, PDF generation
* CI/CD with GitHub Actions
* API documentation using Swagger
* Aligned with FinTech backend interview requirements

---

## 🤝 Contributors

* [Shreya Shukla](https://github.com/Shreyaa983)

---

## 📬 Feedback or Questions?

Feel free to open an [issue](https://github.com/Shreyaa983/finloans-backend.git) or message me on [LinkedIn](https://linkedin.com/in/shreya--9833--).

---

## ⭐ Star This Repo

If you found this project helpful or inspiring, please consider ⭐ starring the repo. It motivates open-source work!

```
