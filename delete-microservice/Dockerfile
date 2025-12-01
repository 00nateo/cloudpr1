FROM alpine:3.22.1
COPY ["main.sh", "/usr/local/bin/"]
RUN \
    apk add --no-cache \ 
        bash \
        go &&\
    chmod +x /usr/local/bin/main.sh

WORKDIR /app
COPY ["go.mod", "go.sum", "/app/"]
RUN go mod download
COPY ["server.go", "/app/"]
RUN go build -o chatbotservice server.go && apk del go
EXPOSE 8080
USER root 
ENTRYPOINT ["/usr/local/bin/main.sh"]
