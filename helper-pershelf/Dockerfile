# 1️⃣ Go ortamı
FROM golang:1.24 

# 2️⃣ Konteynerde çalışma klasörü
WORKDIR /app

# 3️⃣ Dosyaları kopyala
COPY . .

# 4️⃣ Modülleri indir
RUN go mod tidy

# 5️⃣ Uygulamayı çalıştır
CMD ["go", "run", "main/main.go"]
