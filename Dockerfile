FROM golang:1.20

ARG DATABASE_URL = "postgres://postgres:D3Cg241fG--F*EcddA1g*3b1BbE2Gc-F@monorail.proxy.rlwy.net:50003/railway"

ARG JWT_KEY = "b501561895ce919484a3253d495bea8be8ce903cf7bc422e78a1671d175b667b"

ARG HASH_SALT="10"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /alfred-server

EXPOSE 8080

CMD ["/alfred-server"]
