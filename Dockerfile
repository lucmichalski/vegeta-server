# Build stage
FROM golang:1.15-alpine AS build-env
MAINTAINER Luc Michalski <lmichalski@evolutive-business.com>

ENV ROOT=/vegeta-server
ADD . $ROOT
WORKDIR $ROOT
RUN apk add --no-cache --no-progress make && \
    make build && \
    apk del make 

# Final stage
FROM gcr.io/distroless/static
COPY --from=build-env /vegeta-server/bin/vegeta-server .
ENTRYPOINT ["./vegeta-server"]
