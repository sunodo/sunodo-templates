# syntax=docker.io/docker/dockerfile:1.4
FROM ubuntu:22.04 as builder

RUN <<EOF
apt-get update
apt-get install -y --no-install-recommends \
    busybox-static=1:1.30.1-7ubuntu3 \
    ca-certificates=20230311ubuntu0.22.04.1 \
    curl=7.81.0-1ubuntu1.13 \
    openjdk-18-jre=18.0.2+9-2~22.04
rm -rf /var/lib/apt/lists/*
EOF

ARG MAVEN_VERSION=3.9.5
ARG BASE_URL=https://dlcdn.apache.org/maven/maven-3/${MAVEN_VERSION}/binaries

RUN <<EOF
mkdir -p /usr/share/maven /usr/share/maven/ref
curl -fsSL -o /tmp/apache-maven.tar.gz ${BASE_URL}/apache-maven-${MAVEN_VERSION}-bin.tar.gz
tar -xzf /tmp/apache-maven.tar.gz -C /usr/share/maven --strip-components=1
rm -f /tmp/apache-maven.tar.gz
ln -s /usr/share/maven/bin/mvn /usr/bin/mvn
EOF

WORKDIR /opt/cartesi/dapp
COPY . .
RUN mvn package

FROM --platform=linux/riscv64 riscv64/ubuntu:22.04

LABEL io.sunodo.sdk_version=0.2.0
LABEL io.cartesi.rollups.ram_size=128Mi

ARG MACHINE_EMULATOR_TOOLS_VERSION=0.12.0
RUN <<EOF
apt-get update
apt-get install -y --no-install-recommends \
    busybox-static=1:1.30.1-7ubuntu3 \
    ca-certificates=20230311ubuntu0.22.04.1 \
    curl=7.81.0-1ubuntu1.13 \
    openjdk-18-jre=18.0.2+9-2~22.04
cd /tmp
curl -fsSL https://github.com/cartesi/machine-emulator-tools/releases/download/v${MACHINE_EMULATOR_TOOLS_VERSION}/machine-emulator-tools-v${MACHINE_EMULATOR_TOOLS_VERSION}.tar.gz \
  | tar -C / --overwrite -xvzf -
rm -rf /var/lib/apt/lists/*
EOF

ENV PATH="/opt/cartesi/bin:${PATH}"

WORKDIR /opt/cartesi/dapp
COPY --from=builder /opt/cartesi/dapp/target/my-app-1.0-SNAPSHOT-jar-with-dependencies.jar .
#COPY target/my-app-1.0-SNAPSHOT-jar-with-dependencies.jar .

ENV ROLLUP_HTTP_SERVER_URL="http://127.0.0.1:5004"

ENTRYPOINT ["rollup-init"]
CMD ["java", "-jar", "my-app-1.0-SNAPSHOT-jar-with-dependencies.jar"]
