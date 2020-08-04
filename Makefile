# Dgraph Go SDK
sdk:
	go get github.com/dgraph-io/dgo/

# Dgraph on Docker
docker:
	docker pull dgraph/standalone:master
	mkdir -p ~/dgraph

# Run Dgraph zero
zero:
	docker run -it -p 5080:5080 -p 6080:6080 -p 8080:8080 \
	  -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph --name dgraph \
	  dgraph/dgraph:v20.03.0 dgraph zero

# Run Dgraph alpha
alpha:
	docker exec -it dgraph dgraph alpha --lru_mb 2048 --zero localhost:5080 --whitelist 0.0.0.0/0

# Run ratel (Dgraph UI)
ratel:
	docker exec -it dgraph dgraph-ratel

#Add GraphQL Schema
graphschema:
	cd api-rest/data/schema/schema.graphql
	curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'

dgraph:
	npm install
	npm run dgraph

letsgo:
	go run api-rest/main.go

database:
	go run api-rest/database.go

#Dgraph GraphQL
graph:
	docker run -it -p 8080:8080 dgraph/standalone:master
