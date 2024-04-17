FROM golang:alpine3.19 as builder
RUN apk update && apk upgrade && apk add build-base git make sed
ENV GO111MODULE=on
RUN go install github.com/silenceper/gowatch@latest
# install migrate to perform migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# swagger addition via swagger-ui inyection
WORKDIR /go/src/github.com/sail3/interfell-vaccinations/swagger
RUN git clone https://github.com/sail3/swagger-ui && \
 cp -r swagger-ui/. . && rm -r swagger-ui/ && sed -i 's+https://petstore.swagger.io/v2/swagger.json+/swagger/swagger.yml+g' swagger-initializer.js
COPY ./oas/oas.yml ./swagger.yml

WORKDIR /go/src/github.com/sail3/interfell-vaccinations
COPY . .

RUN GIT_COMMIT="deuna-opp-hash" && \
  git config --global --add safe.directory /go/src/github.com/sail3/interfell-vaccinations && \
  go build -o service -ldflags "-X 'github.com/sail3/interfell-vaccinations/internal/config.serviceVersion=$GIT_COMMIT'" ./cmd

FROM alpine:3.19

COPY --from=builder /go/src/github.com/sail3/interfell-vaccinations/service /
COPY --from=builder /go/src/github.com/sail3/interfell-vaccinations/swagger /swagger
COPY --from=builder /go/src/github.com/sail3/interfell-vaccinations/migrations /migrations

ENTRYPOINT [ "./service" ]