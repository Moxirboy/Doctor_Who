FROM golang:1.21 as builder

WORKDIR /app

RUN go install github.com/cespare/reflex@v0.3.1

COPY reflex.conf .
COPY start.sh .

ENTRYPOINT ["reflex", "-c", "/app/reflex.conf"]
