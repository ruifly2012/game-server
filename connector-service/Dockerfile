FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

ADD connector-service /app/connector-service

CMD [ "./connector-service" ]