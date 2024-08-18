# Bước 1: Build stage
FROM golang:1.20 AS builder

WORKDIR /app

# Chỉ copy những file cần thiết trước, để tận dụng cache của Docker
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn vào
COPY . ./

# Build ứng dụng với các flags để tối ưu hóa kích thước và loại bỏ debug info
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /docker

# Bước 2: Minimal runtime stage
FROM scratch

# Copy binary từ build stage sang
COPY --from=builder /docker /docker

# Thiết lập cổng
EXPOSE 8080

# Thiết lập lệnh khởi động
CMD ["/docker"]