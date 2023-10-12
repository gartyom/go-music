FROM golang:latest AS builder

WORKDIR /

COPY ./app /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine 

ARG APP_DIR
ARG STATIC_DIR

WORKDIR ${APP_DIR}

COPY --from=builder /app/app ${APP_DIR}/app

COPY --from=builder /app/static ${STATIC_DIR}

CMD ["./app"]
