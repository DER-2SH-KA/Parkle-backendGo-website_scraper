FROM alpine:edge
MAINTAINER "DER-2SH-KA"
WORKDIR /app
COPY build/application .
RUN chmod +x ./application
ENTRYPOINT ["./application"]
