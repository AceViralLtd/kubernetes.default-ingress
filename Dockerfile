FROM golang as compiler 

WORKDIR /app
COPY main.go .

RUN CGO_ENABLED=0 go build -o http-server .

# Actual final image
FROM scratch 

COPY --from=compiler /app/http-server
EXPOSE 8081

CMD ["./http-server"]
