FROM docker.io/library/golang:1.19.5-bullseye AS build

# Install mage
WORKDIR /temp
RUN git clone https://github.com/magefile/mage
WORKDIR /temp/mage
RUN go run bootstrap.go install

# Compile binary
WORKDIR /temp/src
COPY . .
RUN mage -v build:binary

FROM scratch

# Install binary
COPY --from=build /temp/src/build/site /app

ENTRYPOINT ["/app"]