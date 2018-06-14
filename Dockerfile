# A image to build golang application.
FROM golang:alpine AS builder
WORKDIR /go/src/anbillon.com/gpnm
# Copy the local package files to the container's workspace.
ADD . /go/src/anbillon.com/gpnm
RUN mkdir resources && mv public resources/
RUN apk add --no-cache gcc musl-dev
RUN GOOS=linux go build -a -tags netgo -ldflags '-w'

# Make golang application image.
FROM alpine:3.7
MAINTAINER AnbillonTeam <anbillonteam@gmail.com>
WORKDIR /usr/local/bin/
COPY --from=builder /go/src/anbillon.com/gpnm/gpnm /usr/local/bin/
#RUN mkdir /usr/local/bin/public
COPY --from=builder /go/src/anbillon.com/gpnm/resources /usr/local/bin/
# Run the by default when the container starts.
COPY docker-entrypoint.sh /usr/local/bin/
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["gpnm", "start"]
# The service listens on port 50000.
EXPOSE 50000