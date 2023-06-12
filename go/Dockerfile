
# syntax=docker.io/docker/dockerfile:1.4
FROM cartesi/rootfs:0.17.0 as build-stage

ARG GOVERSION=1.20.5.linux-amd64

WORKDIR /opt/build

RUN wget https://go.dev/dl/go${GOVERSION}.tar.gz && tar -C /usr/local -xzf go${GOVERSION}.tar.gz

ENV GOOS=linux
ENV GOARCH=riscv64
ENV CGO_ENABLED=1
ENV CC=riscv64-cartesi-linux-gnu-gcc
ENV PATH=/usr/local/go/bin:${PATH}

COPY src .

RUN make


# runtime stage: produces final image that will be executed
FROM --platform=linux/riscv64 riscv64/ubuntu:22.04

LABEL io.sunodo.sdk_version=0.1.0
LABEL io.cartesi.rollups.ram_size=128Mi

RUN <<EOF
apt-get update
apt-get install -y --no-install-recommends busybox-static=1:1.30.1-7ubuntu3
rm -rf /var/lib/apt/lists/*
EOF

COPY --from=sunodo/machine-emulator-tools:0.11.0-ubuntu22.04 / /
ENV PATH="/opt/cartesi/bin:${PATH}"

WORKDIR /opt/cartesi/dapp
COPY --from=build-stage /opt/build/dapp .

ENV ROLLUP_HTTP_SERVER_URL="http://127.0.0.1:5004"

ENTRYPOINT ["rollup-init"]
CMD ["/opt/cartesi/dapp/dapp"]
