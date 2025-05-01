FROM golang:1.23

WORKDIR /app/api
COPY api/go.mod api/go.sum ./
RUN go mod download
COPY api/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main.exe

WORKDIR /app
COPY Makefile .
COPY migrations migrations
COPY entrypoints entrypoints
RUN chmod +x entrypoints/api-entrypoint.sh

EXPOSE 8080

ENTRYPOINT [ "entrypoints/api-entrypoint.sh" ]