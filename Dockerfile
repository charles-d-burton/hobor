ARG BUILD_FROM
FROM $BUILD_FROM

RUN apk add --no-cache go

WORKDIR /data

COPY ./ .

RUN go build -o /hobor

ENTRYPOINT ["/hobor"]
