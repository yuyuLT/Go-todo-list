ARG BASE_IMAGE=golang
ARG IMAGE_VERSION=alpine

FROM ${BASE_IMAGE}:${IMAGE_VERSION} AS dev

ARG WORK_DIR=/app
ARG MODULE_PATH=example.com/module
ARG PORTS=80

ENV MODULE_PATH=${MODULE_PATH}

WORKDIR ${WORK_DIR}

RUN apk add --update gcc musl-dev \
  && go install github.com/cosmtrek/air@latest

EXPOSE ${PORTS}

ENTRYPOINT ["entrypoint.sh"]

CMD [ "/bin/sh" ]



FROM ${BASE_IMAGE}:${IMAGE_VERSION} AS builder

WORKDIR /tmp

COPY . .

RUN mkdir /app \
  && go mod tidy \
  && go build -o /app/main /tmp/app/main.go



FROM scratch AS prod

ARG WORK_DIR=/app
ARG PORTS=80

WORKDIR ${WORK_DIR}

COPY --from=builder /app /app

EXPOSE ${PORTS}

CMD [ "/app/main" ]
