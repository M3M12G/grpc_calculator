FROM golang
#run on root folder of project
#Setting environment variables, which provides running our golang binary file on linux
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./calculator/calculator_server /app/calculator/calculator_server
COPY ./calculator/calculatorpb /app/calculator/calculatorpb
WORKDIR /app
RUN go build /app/calculator/calculator_server && chmod +x /app/calculator/calculator_server
ENTRYPOINT ["./calculator_server"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./calculator_server"]
EXPOSE 7777