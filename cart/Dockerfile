FROM alpine:latest
ADD cart /cart

#这种情况是因为动态链接库位置错误导致的，alpine镜像使用的是musl libc而不是gun libc。因而动态链接库的位置不一致。
# 由于alpine中缺少动态链接库一下命令是解决此问题
ENV GLIBC_REPO=https://github.com/sgerrand/alpine-pkg-glibc
ENV GLIBC_VERSION=2.30-r0

RUN set -ex && \
    apk --update add libstdc++ curl ca-certificates && \
    for pkg in glibc-${GLIBC_VERSION} glibc-bin-${GLIBC_VERSION}; \
        do curl -sSL ${GLIBC_REPO}/releases/download/${GLIBC_VERSION}/${pkg}.apk -o /tmp/${pkg}.apk; done && \
    apk add --allow-untrusted /tmp/*.apk && \
    rm -v /tmp/*.apk && \
    /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib

ENTRYPOINT ["/cart"]