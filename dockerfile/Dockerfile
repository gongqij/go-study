FROM --platform=$BUILDPLATFORM golang:1.16 AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG GOBUILDTAG
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"
ENV APP_SRC_DIR "/src"
COPY ./ $APP_SRC_DIR
WORKDIR $APP_SRC_DIR
RUN  mkdir -p /src/build/config && \
     cp -rf config/* /src/build/config/ && \
     cp startup.sh /src/build/ && \
     go env -w GO111MODULE=on && \
     go env -w GOPROXY=https://goproxy.cn,direct && \
     go build -o /src/build/app


FROM bitnami/minideb:stretch
ENV APP_DIR "/home/www"
WORKDIR /home/www/
RUN mkdir -p $APP_DIR
COPY --from=builder /src/build  $APP_DIR
RUN chmod +x /home/www/startup.sh
EXPOSE 80
ENTRYPOINT ["/home/www/startup.sh"]
