FROM openjdk:18-jdk-slim as BUILD

COPY . /src
WORKDIR /src
RUN ./gradlew --version
RUN ./gradlew --no-daemon jar

FROM openjdk:18-slim as kotlin

COPY --from=BUILD /src/build/libs/jcasbin-sample-1.0-SNAPSHOT.jar /bin/runner/run.jar
WORKDIR /bin/runner

CMD ["java","-jar","run.jar"]

FROM golang:1.18-alpine AS go-builder
WORKDIR /app
COPY ./golang ./

ENV CGO_ENABLED=0
RUN go get -d -v ./cmd
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/bin/casbin-sample ./cmd/main.go

FROM scratch AS go
COPY --from=go-builder /usr/bin/casbin-sample /usr/bin/casbin-sample
