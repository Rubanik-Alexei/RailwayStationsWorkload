FROM golang:1.17.3 as builder

WORKDIR /redis_service

COPY go.* ./
RUN go mod download
COPY ./internal/redis_service ./
#RUN go build -v -o rds
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rds ./cmd/redis

FROM alpine:latest                       
WORKDIR /redis_service
COPY --from=builder ./redis_service .  

EXPOSE 9091

CMD [ "/redis_service/rds" ]