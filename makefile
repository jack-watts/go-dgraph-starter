# Sett up makefile variables
DGRAPH  := dgraph/standalone
RATEL 	:= dgraph/ratel:latest

# dev-docker: pulls all required docker images.
dev-docker:
	docker pull $(DGRAPH)
	docker pull $(RATEL)

# dev-dgraph-local: run the conatiners in docker from the stored images
dev-dgraph-local:
	docker run --name dgraph-local -d -it -p "8080:8080" -p "9080:9080" -p "8090:8090" -v ~/dgraph $(DGRAPH)
	docker run --name ratel-local --platform linux/amd64 -d -p "8000:8000"  $(RATEL)