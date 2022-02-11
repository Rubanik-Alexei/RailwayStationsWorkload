#FROM golang:1.17.3-alpine
FROM golang:1.17.3 AS builder 

WORKDIR /gateway_service
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gtw .

#RUN go build -v -o gtw

FROM alpine:latest                       
WORKDIR /gateway_service
COPY --from=builder ./gateway_service .  

EXPOSE 9090

CMD [ "/gateway_service/gtw" ] 