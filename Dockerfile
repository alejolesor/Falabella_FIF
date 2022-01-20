FROM golang:alpine AS build
WORKDIR /go/src/falabella
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/falabella cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/falabella /go/bin/falabella
EXPOSE 3000
ENTRYPOINT ["/go/bin/falabella"]
