FROM golang:alpine as Builder
WORKDIR /app
COPY . /app
RUN go build -o bootstrap .

FROM scratch
WORKDIR /app
COPY --from=builder /app/bootstrap ./bootstrap
COPY ./migrations ./migrations
ENTRYPOINT [ "./bootstrap" ]