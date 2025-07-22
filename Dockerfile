ARG BUILD_FROM
FROM $BUILD_FROM

RUN apk add --no-cache go

WORKDIR /data

COPY ./ .

RUN go build -tags excludetinygo -o /hobor

ENTRYPOINT ["/hobor"]
