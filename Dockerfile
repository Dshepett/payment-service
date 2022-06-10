FROM golang:alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN export CGO_ENABLED=0 && go build -o /application ./cmd


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY .env .env
COPY ./docs ./docs

COPY --from=build /application /application

EXPOSE 8080

ENTRYPOINT ["./application"]