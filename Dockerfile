FROM alpine

COPY app /
COPY config.yml /

ENTRYPOINT ["./app"]