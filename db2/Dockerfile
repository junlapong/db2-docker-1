# FROM ibmcom/db2
FROM ibmcom/db2:11.5.5.1

RUN set -x && yum install java -y && yum clean all

WORKDIR /database

COPY assets/createdb.sh /database/scripts/createdb.sh
# RUN su - db2inst1 -c /database/scripts/createdb.sh
	
# CMD ["db2start"]
CMD ["ping", "www.google.com"]
