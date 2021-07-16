export DB2HOME=$HOME/git-space/docker/docker-was-db2/clidriver
# export DB2HOME=/app/clidriver
export CGO_CFLAGS=-I$DB2HOME/include
export CGO_LDFLAGS=-L$DB2HOME/lib

#Linux:
# export LD_LIBRARY_PATH=$DB2HOME/lib

#Mac:
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:$DB2HOME/lib

export DB2DSN='HOSTNAME=localhost;PORT=50001;PROTOCOL=TCPIP;DATABASE=testdb;UID=db2inst1;PWD=pa55w0rd'
