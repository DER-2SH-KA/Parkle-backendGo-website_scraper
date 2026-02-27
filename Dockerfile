FROM alpine:edge
MAINTAINER "DER-2SH-KA"
WORKDIR /app
COPY build/main .
RUN chmod +x ./main
ENTRYPOINT ["./main"]

