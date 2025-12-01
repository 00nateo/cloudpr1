FROM alpine:3.22.1
COPY ["main.sh", "/usr/local/bin/"]
RUN \
    apk add --no-cache \ 
        bash \
        go &&\
    sed -i 's/\r$//' /usr/local/bin/main.sh &&\
    chmod +x /usr/local/bin/main.sh

WORKDIR /app
COPY ["go.mod", "go.sum", "/app/"]
RUN go mod download
COPY ["delete.go", "/app/"]
RUN go build -o chatbotservice delete.go && apk del go
EXPOSE 8083
USER root 
ENTRYPOINT ["/usr/local/bin/main.sh"]
