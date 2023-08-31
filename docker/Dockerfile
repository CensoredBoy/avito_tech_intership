FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /avito_task_segments

COPY ./pkg ./pkg
COPY ./cmd ./cmd

RUN go mod init avito_task_segments
RUN go mod tidy
RUN go mod download
RUN go build /avito_task_segments/cmd/main.go

EXPOSE 8080

CMD ["./main"]
