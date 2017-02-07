FROM alpine:latest

RUN echo "Asia/shanghai" >> /etc/timezone

COPY ./main /bin/kk-article

RUN chmod +x /bin/kk-article

COPY ./config /config

COPY ./app.ini /app.ini

ENV KK_ENV_CONFIG /config/env.ini

VOLUME /config

CMD kk-article $KK_ENV_CONFIG

