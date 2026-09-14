FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /aegis ./cmd/aegis

FROM alpine:3.24
RUN addgroup -S aegis && adduser -S -G aegis aegis && mkdir -p /workspace && chown aegis:aegis /workspace
COPY --from=build /aegis /usr/local/bin/aegis
USER aegis
WORKDIR /workspace
ENTRYPOINT ["aegis"]
CMD ["help"]
