FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/orbit .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build /app/orbit /orbit
COPY --from=build /src/configs /configs
COPY --from=build /src/admin/static /admin/static
COPY --from=build /src/templates /templates

EXPOSE 9000 2121

CMD ["/orbit"]
