FROM ubuntu:20.04


RUN apt update && \
    apt install -yq --no-install-recommends vim gnupg curl wget ca-certificates && \
    wget -O - https://repo.fortinet.com/repo/6.4/ubuntu/DEB-GPG-KEY | apt-key add - && \
    echo "deb [arch=amd64] https://repo.fortinet.com/repo/6.4/ubuntu/ /bionic multiverse" >> /etc/apt/sources.list && \
    apt update && \
    apt install -yq forticlient && \
    apt clean && \
    apt autoremove -y && \
    rm -rf /var/lib/apt/lists/*

