# syntax=docker/dockerfile:1
FROM scratch
COPY target/hello.wasm /hello.wasm
ENTRYPOINT [ "hello.wasm" ]