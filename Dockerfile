FROM openjdk:18-jdk-slim as BUILD

COPY . /src
WORKDIR /src
RUN ./gradlew --version
RUN ./gradlew --no-daemon jar

FROM openjdk:18-slim

COPY --from=BUILD /src/build/libs/jcasbin-sample-1.0-SNAPSHOT.jar /bin/runner/run.jar
WORKDIR /bin/runner

CMD ["java","-jar","run.jar"]