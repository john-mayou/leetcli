# --- Stage: base (shared deps) -----
FROM golang:1.23 AS base

# nsjail runtime deps + bash tooling
RUN apt-get -y update && apt-get install -y \
    libc6 libstdc++6 libprotobuf32 libnl-route-3-200 \
    bash grep gawk sed jq \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY api/go.mod api/go.sum api/
RUN cd api && go mod download

# --- Stage: nsjail builder ----
FROM debian:bookworm-slim AS nsjail-builder

RUN apt-get -y update && apt-get install -y \
    autoconf bison flex gcc g++ git libprotobuf-dev libnl-route-3-dev libtool \
    libc6 libstdc++6 libprotobuf32 libnl-route-3-200 make pkg-config protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*
RUN git clone --recursive --branch 3.4 https://github.com/google/nsjail.git
RUN cd /nsjail && make clean && make

# --- Stage: build binary ---
FROM base AS build

COPY api/ api/
RUN cd api && CGO_ENABLED=0 GOOS=linux go build -o main.exe

# --- Stage: dev ---
FROM base AS dev

COPY --from=nsjail-builder /nsjail/nsjail /usr/local/bin/
COPY --from=build /api/main.exe /main.exe

COPY Makefile .
COPY migrations migrations
COPY api/entrypoints entrypoints

EXPOSE 8080

RUN chmod +x entrypoints/dev.sh
ENTRYPOINT [ "entrypoints/dev.sh" ]

# --- Stage: testing ---
FROM base AS test

WORKDIR /app

COPY --from=nsjail-builder /nsjail/nsjail /usr/local/bin/
COPY Makefile .
COPY migrations migrations

# --- Stage: prod ---
FROM base AS runtime

WORKDIR /app

COPY --from=nsjail-builder /nsjail/nsjail /usr/local/bin/
COPY --from=build /api/main.exe /main.exe

COPY Makefile .
COPY migrations migrations
COPY api/entrypoints entrypoints

EXPOSE 8080

RUN chmod +x entrypoints/prod.sh
ENTRYPOINT [ "entrypoints/prod.sh" ]