FROM fim_server AS builder

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app


COPY grouprpc.yaml .

COPY --from=builder /build/fim_group/group_rpc/grouprpc .

CMD ["./grouprpc","-f","grouprpc.yaml"]