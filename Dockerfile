FROM alpine
WORKDIR /app
COPY bin/mr-wordcount ./
CMD ["/app/mr-wordcount"]