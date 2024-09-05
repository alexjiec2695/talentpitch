FROM golang:1.23
WORKDIR /rest-api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN mkdir -p /rest-api/assets/docs/pqr
RUN go build -o ./out/ /rest-api
EXPOSE ${PORT}
CMD ["./out/api"]