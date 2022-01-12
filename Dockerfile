FROM golang:1.17-bullseye AS builder

WORKDIR /var/app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -v -o vxchan .


FROM debian:bullseye

RUN apt-get update && \
    apt-get install --no-install-recommends -y curl ca-certificates

RUN curl -L -O https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-7.16.2-amd64.deb && \
    dpkg -i filebeat-7.16.2-amd64.deb && \
    rm -rf /var/lib/apt/lists/*

COPY ./filebeat.yml /etc/filebeat/filebeat.yml

RUN mkdir -p /etc/vxchan

COPY ./config.yaml /etc/vxchan
COPY ./entrypoint.sh /etc/vxchan
RUN chmod +x /etc/vxchan/entrypoint.sh

COPY --from=builder /var/app/vxchan /usr/local/bin

ENTRYPOINT [ "/etc/vxchan/entrypoint.sh" ]

CMD [ "/usr/local/bin/vxchan", "--configFile", "/etc/vxchan/config.yaml" ]
