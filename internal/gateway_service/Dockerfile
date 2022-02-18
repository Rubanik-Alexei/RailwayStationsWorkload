FROM golang:1.17.3 as builder

WORKDIR /gateway_service

COPY go.* ./
RUN go mod download
COPY ./internal/gateway_service ./ 
COPY ./cmd/gateway/main.go ./
#RUN go build -v -o gtw
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gtw .

FROM alpine:latest                       
WORKDIR /gateway_service
COPY --from=builder ./gateway_service .  

EXPOSE 9090


CMD [ "/gateway_service/gtw" ] 