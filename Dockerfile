FROM golang:1.19-alpine
ENV TRACE_LEVEL=DEBUG
WORKDIR ${GOPATH}/src/user-ranking
COPY . .
RUN apk add --no-cache bash git openssh curl
RUN cd cmd/api && go install -v
EXPOSE 8080
CMD api -trace_level=${TRACE_LEVEL}