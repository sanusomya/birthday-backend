FROM alpine:latest
ADD --chown=user:group main /
EXPOSE 8002
ENTRYPOINT ["./main"]