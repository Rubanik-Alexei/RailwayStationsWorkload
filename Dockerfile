FROM golang:1.17.3-alpine

WORKDIR /gateway_service
COPY . .
RUN go mod download

RUN go build -v -o gtw

EXPOSE 9090


CMD [ "/gateway_service/gtw" ] 