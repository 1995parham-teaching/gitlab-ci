FROM alpine:3.18

LABEL maintainer="Parham Alvani <parham.alvani@gmail.com>"

WORKDIR /app/

COPY app .

ENTRYPOINT ["./app"]
