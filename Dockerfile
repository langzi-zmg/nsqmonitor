FROM ccr.ccs.tencentyun.com/dhub.wallstcn.com/alpine:3.5
ENV CONFIGOR_ENV ivktest
ADD conf/ /conf
ADD server /
ENTRYPOINT [ "/server" ]