FROM fim_server AS builder

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app


COPY auth.yaml .

COPY --from=builder /build/fim_auth/auth_api/auth .

CMD ["./auth","-f","auth.yaml"]