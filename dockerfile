# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# กำหนด working directory ภายใน container
WORKDIR /app

# Copy go.mod และ go.sum ไปยัง container
COPY go.mod go.sum ./

# ดาวน์โหลด dependencies ที่จำเป็น
RUN go mod tidy

# Copy โค้ดทั้งหมดไปยัง container
COPY . .

# Build แอป Go (ให้ binary เป็นไฟล์ชื่อว่า app)
RUN go build -o app .

# Stage 2: Run the Go binary in a smaller container
FROM alpine:latest

# กำหนด working directory ภายใน container
WORKDIR /root/

# Copy binary ที่ build จาก stage ก่อนหน้า
COPY --from=builder /app/app .

# เปิดพอร์ตที่แอปพลิเคชันจะฟัง (เช่น 8080)
EXPOSE 8080

# สั่งให้ container รันแอป Go เมื่อ container เริ่มทำงาน
CMD ["./app"]
