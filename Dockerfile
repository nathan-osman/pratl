# The build is completed in two stages:
#
# 1) the JavaScript UI
# 2) the server binary
#
# The server binary embeds the UI.

# Firstly, build the JavaScript UI
FROM node:latest
ADD ui /src
WORKDIR /src
RUN npm install
RUN npm run build

# Secondly, compile the Go binary
FROM golang:latest
ENV CGO_ENABLED=0
ADD . /src
COPY --from=0 /src/build /src/ui/build
WORKDIR /src
RUN go build

# Lastly, create the container with the resultant binary
FROM scratch
COPY --from=1 /src/pratl /usr/local/bin/
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/
ENTRYPOINT ["/usr/local/bin/pratl"]
