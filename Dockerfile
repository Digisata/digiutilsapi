FROM golang:1.19-alpine
LABEL maintener="Hanif Naufal <hnaufal123@gmail.com>"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go mod vendor
COPY . .
RUN go build -o ./out/dist .
CMD ["./out/dist"]