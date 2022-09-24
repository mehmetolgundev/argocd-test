FROM golang:1.18
WORKDIR /app
COPY . .
CMD ["go","run","."]
