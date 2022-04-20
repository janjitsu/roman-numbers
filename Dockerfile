FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY app/ go.* /src/
RUN go mod download
COPY . .
WORKDIR /src/app/http/
RUN CGO_ENABLED=0 go build -o /bin/server

FROM scratch
COPY --from=build /bin/server /bin/server
EXPOSE 80
ENTRYPOINT ["/bin/server"]
