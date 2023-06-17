# Sett up makefile variables
DGRAPH  := dgraph/standalone
RATEL 	:= dgraph/ratel:latest

# dev-docker: pulls all required docker images.
dev-docker:
	docker pull $(DGRAPH)
	docker pull $(RATEL)

# dev-dgraph-local: run the conatiners in docker from the stored images.
dev-dgraph-local:
	docker run --name dgraph-local -d -it -p "8080:8080" -p "9080:9080" -p "8090:8090" -v ~/dgraph $(DGRAPH)
	docker run --name ratel-local --platform linux/amd64 -d -p "8000:8000"  $(RATEL)

# dev-dgraph-seed: will popoulate the DGraph database via its GraphQL endpoint with a schema nad some dummy data.
dev-dgraph-seed:
	@echo -e "Setting Graph schema..."
	curl -X POST 'localhost:8080/admin/schema' --data-binary '@./schema/schema.graphql'
	@echo -e "seeding Graph database..."	
	curl -X POST 'http://localhost:8080/graphql' --header 'Content-Type: application/graphql' --data-binary '@./schema/sample.data.graphql' 
	@echo -e "seeding Graph database complete"

# dev-run: will execute our local Go code without compiling.
dev-run:
	go run ./app/services/api