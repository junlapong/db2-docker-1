TAG=11.5.5.1

docker:
	@docker build -t db2:${TAG} .

start:
	@docker run -h db2 --name db2 \
	        --restart=always --privileged=true \
			-p 50001:50000 --env-file .env \
			db2:${TAG}

logs:
	@docker logs -f db2

shell:
	@docker exec -ti db2 bash -c "su - db2inst1"

# db2
# db2 => connect to testdb
# db2 => ?
# db2 => LIST ACTIVE DATABASES

# --detach
# -v ${PWD}/database:/database \
# jdbc:/database/config/db2inst1/sqllib/java \
