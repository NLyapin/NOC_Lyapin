FROM frrouting/frr:latest

RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    gnmi-gateway

CMD ["bash"]
