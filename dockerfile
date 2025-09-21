FROM golang:1.23-alpine AS build
WORKDIR /birthday-backend
ADD --chown=user:group . /birthday-backend/
RUN go mod tidy
RUN go build -o main main.go

FROM alpine:latest
COPY --from=build /birthday-backend/main .
EXPOSE 8002
ENTRYPOINT ["./main"]