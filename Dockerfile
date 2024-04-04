FROM golang:1.22.1 as base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/emplyee-api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=base /app/emplyee-api /emplyee-api

EXPOSE 8080

ENTRYPOINT ["/emplyee-api"]
