# --- Stage 1: nsjail builder ----
FROM debian:bookworm-slim AS nsjail-builder

RUN apt-get -y update && apt-get install -y \
    autoconf bison flex gcc g++ git libprotobuf-dev libnl-route-3-dev libtool \
    libc6 libstdc++6 libprotobuf32 libnl-route-3-200 make pkg-config protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*
RUN git clone --recursive --branch 3.4 https://github.com/google/nsjail.git
RUN cd /nsjail && make clean && make

# --- Stage 2: runtime image -----
FROM golang:1.23

# nsjail runtime deps + bash tooling
RUN apt-get -y update && apt-get install -y \
    libc6 libstdc++6 libprotobuf32 libnl-route-3-200 \
    bash grep gawk sed jq \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY api/go.mod api/go.sum api/
RUN cd api && go mod download
COPY api/ api/
RUN cd api && CGO_ENABLED=0 GOOS=linux go build -o main.exe

COPY Makefile .
COPY migrations migrations
COPY entrypoints entrypoints
COPY --from=nsjail-builder /nsjail/nsjail /usr/local/bin/

EXPOSE 8080

RUN chmod +x entrypoints/api-entrypoint.sh
ENTRYPOINT [ "entrypoints/api-entrypoint.sh" ]