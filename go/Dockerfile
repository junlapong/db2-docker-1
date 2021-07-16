FROM golang:1.16.5-buster

RUN apt-get update && apt-get -y install build-essential libxml2 vim telnet

WORKDIR /app

RUN git clone https://github.com/ibmdb/go_ibm_db && \
    cd go_ibm_db/installer && go run setup.go

COPY . .
# COPY go.mod .
# COPY go.sum .

# environment
ENV DB2HOME          /app/clidriver
ENV CGO_CFLAGS       -I$DB2HOME/include
ENV CGO_LDFLAGS      -L$DB2HOME/lib
ENV LD_LIBRARY_PATH  $DB2HOME/lib
ENV DB2DSN           HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=pa55w0rd

RUN go mod download
# RUN . ./db2env.sh
RUN go build -o ./bin/app .

# CMD ["/app/bin/app"]
CMD ["ping", "db2"]
