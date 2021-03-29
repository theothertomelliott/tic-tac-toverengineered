FROM alpine

WORKDIR /root
COPY ./entrypoint/start.sh .
COPY ./entrypoint/restart.sh .
ENTRYPOINT ["./start.sh", "/root/app"]