# TEST GOLANG BACKEND DAN TEST FRONTEND 

## 🛠️ Installation Guide BACKEND
### 1️⃣ Clone Repository
```bash
git clone https://github.com/YourMasterIsFool/test-payment-gateway-be.git
```


### 2️⃣ Buka directory Backend

```bash
cd backend
```

### 2️⃣ Jalankan Docker Compose
Jika menggunakan Docker, jalankan perintah berikut:
```bash
docker-compose up -d
```

### 3️⃣ Install Dependencies
Jalankan perintah berikut untuk menginstal dependensi:
```bash
go mod tidy
```

### 6️⃣ Migrasi Database
Jalankan migrasi menggunakan goose 
```bash
goose up
```


### 6️⃣ Migrasi Database
Jalankan migrasi test repo
```bash
go test -v backend/domain/repository/test
```

### 8️⃣ Jalankan Server
```bash
go run cmd/main.go
```
Akses aplikasi melalui **[http://localhost:8000](http://localhost:8000)**

Dokumentasi Endpoint **[documentasi](https://app.getpostman.com/join-team?invite_code=db2eb05bc70ff2dd23537baee9fb104e621662464d634d06d7370fde6374b78a&target_code=8b52a10d5f8cd625e7f40b9bb079a6ca)




## 🛠️ Installation Guide Frontend


### 2️⃣ Buka directory Backend

```bash
cd frontend
```


### 3️⃣ Install Dependencies
Jalankan perintah berikut untuk menginstal dependensi:
```bash
yarn install
```

### 6️⃣ Migrasi Database
Running Server
```bash
yarn dev
```

