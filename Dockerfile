# Bước 1: Build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Cài đặt tzdata và thiết lập múi giờ
RUN apt-get update && \
    apt-get install -y tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata

# Chỉ copy những file cần thiết trước, để tận dụng cache của Docker
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn vào
COPY . ./

# Build ứng dụng với các flags để tối ưu hóa kích thước và loại bỏ debug info
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /docker

# Bước 2: Minimal runtime stage
FROM alpine:latest

# Cài tzdata cho Alpine
RUN apk add --no-cache tzdata

# Copy binary từ build stage sang
COPY --from=builder /docker /docker

# Thiết lập múi giờ
ENV TZ=Asia/Ho_Chi_Minh
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime

# Thiết lập cổng
EXPOSE 8080

# Thiết lập lệnh khởi động
CMD ["/docker"]