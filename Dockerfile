FROM golang:1.19-alpine
LABEL maintener="Hanif Naufal <hnaufal123@gmail.com>"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .
CMD ["./out/dist"]