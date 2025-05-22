# Copyright 2025 sriov-network-device-plugin Authors. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.23-alpine as builder

COPY . /usr/src/sriov-network-device-plugin

ENV HTTP_PROXY $http_proxy
ENV HTTPS_PROXY $https_proxy
RUN apk add --no-cache --virtual build-dependencies build-base

WORKDIR /usr/src/sriov-network-device-plugin
RUN make clean && \
    make build

FROM golang:1.20-alpine3.16 as ddp-builder

ADD images/ddptool-1.0.1.12.tar.gz /tmp/ddptool/

ENV HTTP_PROXY $http_proxy
ENV HTTPS_PROXY $https_proxy
RUN apk add --no-cache --virtual build-dependencies build-base linux-headers

WORKDIR /tmp/ddptool
RUN make

FROM alpine:3
RUN apk add --no-cache hwdata-pci
COPY --from=builder /usr/src/sriov-network-device-plugin/build/sriovdp /usr/bin/
COPY --from=ddp-builder /tmp/ddptool/ddptool /usr/bin/
WORKDIR /

LABEL io.k8s.display-name="SRIOV Network Device Plugin"

COPY ./images/entrypoint.sh /

RUN rm -rf /var/cache/apk/*

ENTRYPOINT ["/entrypoint.sh"]
