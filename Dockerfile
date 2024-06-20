FROM golang:1.21.4 AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
COPY go.sum .

# RUN go env -w GOPROXY=https://proxy.golang.com.cn,https://goproxy.cn,direct
COPY . .
RUN go mod tidy


RUN go build -o fim_gateway/gateway fim_gateway/gateway.go


RUN go build -o fim_chat/chat_rpc/chatrpc fim_chat/chat_rpc/chatrpc.go

RUN go build -o fim_file/file_rpc/filerpc fim_file/file_rpc/filerpc.go

RUN go build -o fim_group/group_rpc/grouprpc fim_group/group_rpc/grouprpc.go

RUN go build -o fim_settings/settings_rpc/settingsrpc fim_settings/settings_rpc/settingsrpc.go

RUN go build -o fim_user/user_rpc/userrpc fim_user/user_rpc/userrpc.go


RUN go build -o fim_auth/auth_api/auth fim_auth/auth_api/auth.go

RUN go build -o fim_chat/chat_api/chat fim_chat/chat_api/chat.go

RUN go build -o fim_file/file_api/file fim_file/file_api/file.go

RUN go build -o fim_group/group_api/group fim_group/group_api/group.go

RUN go build -o fim_logs/logs_api/logs fim_logs/logs_api/logs.go

RUN go build -o fim_settings/settings_api/settings fim_settings/settings_api/settings.go

RUN go build -o fim_user/user_api/user fim_user/user_api/user.go

















