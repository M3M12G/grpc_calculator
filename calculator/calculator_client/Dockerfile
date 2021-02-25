FROM golang
#run on root folder of project
#Setting environment variables, which provides running our golang binary file on linux
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./calculator/calculator_client /app/calculator/calculator_client
COPY ./calculator/calculatorpb /app/calculator/calculatorpb
WORKDIR /app
RUN go build /app/calculator/calculator_client && chmod +x /app/calculator/calculator_client
ENTRYPOINT ["./calculator_client"]
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./calculator_client"]
EXPOSE 7777