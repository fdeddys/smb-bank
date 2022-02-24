
# FROM golang:1.11.0
# LABEL ver="1.0"
# LABEL description ="Application SMB"
# WORKDIR /app/
# COPY app_be /app/
# COPY conf/app.conf /app/conf/
# COPY migration/. /app/migration/
# CMD ["./app_be"]

FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app_be
EXPOSE 8888
CMD ./app_be
