FROM golang:1.17.3 as builder

WORKDIR /workload_service

COPY go.* ./
RUN go mod download
COPY *.csv ./ && COPY /internal/workload_service ./
#RUN go build -v -o wl
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wl .

FROM alpine:latest                       
WORKDIR /workload_service
COPY --from=builder ./workload_service .  


EXPOSE 9092

CMD [ "/workload_service/wl" ]