FROM fim_server AS builder

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app


COPY chatrpc.yaml .

COPY --from=builder /build/fim_chat/chat_rpc/chatrpc .

CMD ["./chatrpc","-f","chatrpc.yaml"]
