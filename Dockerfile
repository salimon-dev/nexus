FROM golang:alpine as Builder
WORKDIR /app
COPY . /app
RUN go build -o bootstrap .

FROM scratch
WORKDIR /app
COPY --from=builder /app/bootstrap ./bootstrap
ENTRYPOINT [ "./bootstrap" ]