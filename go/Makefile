TAG=buster
DB2DSN='HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=pa55w0rd'

all: buster

buster:
	@docker build -t go-db2:${TAG} .

alpine:
	@docker build -f Dockerfile.alpine -t go-db2:alpine .

clean:
	@docker rm $(docker ps -a -q)
	@docker image prune

run:
	@docker run --rm -it --name go-db2 --network host -e ${DB2DSN} go-db2:${TAG} /app/bin/app

#docker run --rm -it --network host alpine sh

stop:
	@docker stop go-db2

shell:
	@docker exec -it go-db2 bash
