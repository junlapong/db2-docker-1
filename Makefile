all:

db-start:
	@docker run -h db2server --name db2server \
	        --restart=always --privileged=true \
			-p 50000:50000 --env-file ./db2/.env \
			-v ${PWD}/database:/database \
			ibmcom/db2:11.5.5.1
db-stop:
	@docker stop db2server && \
	 docker rm db2server && \
	 docker ps -a

logs:
	@docker logs -f db2server

shell:
	@docker exec -ti db2server bash -c "su - db2inst1"

# --detach
# -v ${PWD}/jdbc:/database/config/db2inst1/sqllib/java \