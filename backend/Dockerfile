FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o main

FROM alpine

WORKDIR /app

COPY --from=build /app/main ./

EXPOSE 8080

CMD [ "./main" ]