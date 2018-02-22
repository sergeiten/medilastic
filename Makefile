GO_PACKAGES=$(shell ls -d */ | grep -v vendor)

default: docker

.PHONY: docker
docker:
	docker-compose up --build

.PHONY: docker-force
docker-force:
	docker-compose up --force-recreate --build --remove-orphans

.PHONY: docker-production
docker-production:
	docker-compose up --force-recreate --build --remove-orphans -d

.PHONY: kibana
kibana:
	docker stop kibana && docker rm kibana
	docker run -d --network=medilastic_default -p 5601:5601 --name=kibana docker.elastic.co/kibana/kibana-oss:6.1.1

.PHONY: quality
quality:
	go test -v -race ./...
	go vet ./...
	golint -set_exit_status $(go list ./...)
	megacheck ./...
	gocyclo -over 12 $(GO_PACKAGES)