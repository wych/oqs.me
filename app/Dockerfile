FROM golang:1.13

WORKDIR /app

ENV GOOS=linux GOARCH=amd64

COPY . ./

ENTRYPOINT [ "go", "build", "-o" ]

CMD [ "/cmd/app" ]
