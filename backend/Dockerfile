FROM golang:alpine AS builder

ENV GIN_MODE=release

WORKDIR /server

COPY ./ /server

RUN apk update && apk add --no-cache git
RUN go mod tidy
RUN go build -o api-beatrice /server

FROM gcr.io/distroless/base-debian10

WORKDIR /server

COPY --from=builder /server /server

ENV PORT=3004

EXPOSE $PORT

USER nonroot:nonroot

ENTRYPOINT ["./api-beatrice"]

#*  Command to build the image
# docker build . -t api-beatrice
# docker run -d -p 8080:3004 --name api-beatrice <id_images>  

