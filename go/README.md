
```
hub clone ibmdb/go_ibm_db
cd go_ibm_db/installer
go run setup.go

cd -
cd go

# edit db2env.sh DB2HOME=/path/to/clidriver
. ./db2env.sh
go get -u -v github.com/ibmdb/go_ibm_db
go mod tidy
# go run main.go
go build -o bin/app .
bin/app
```

## Anothers

- [Talking to IBM DB2 with Golang](https://link.medium.com/GLaeftAushb)
