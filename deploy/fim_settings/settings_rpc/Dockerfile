FROM fim_server AS builder

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app


COPY settingsrpc.yaml .

COPY --from=builder /build/fim_settings/settings_rpc/settingsrpc .

CMD ["./settingsrpc","-f","settingsrpc.yaml"]