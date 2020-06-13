BASE_DIRECTORY=${PWD}
GOLANG_BUILDER=mkadiri/golang-builder
GOLANG_TESTER=mkadiri/golang-tester
GOLANG_MICROSERVICE_IMAGE=mkadiri/{{cookiecutter.repo_name}}
PROJECT_URL={{cookiecutter.repo_base_name})/{{cookiecutter.repo_name}}

build:
	docker build -t ${GOLANG_MICROSERVICE_IMAGE} .

test:
	cd ${BASE_DIRECTORY}/docker/tester && \
	docker build -t ${GOLANG_TESTER} . && \
	cd ${BASE_DIRECTORY} && \
	docker-compose -f docker-compose-test.yml up -d && \
    docker logs -f mkadiri-golang-tester && \
    docker rm -f mkadiri-golang-tester

run:
	docker-compose down && docker-compose up -d && docker logs -f mkadiri-{{cookiecutter.repo_name}}

build-test:
	cd ./docker/tester && docker build -t mkadiri/golang-tester .

run-tests:
	docker-compose -f docker-compose-test.yml up -d && \
	docker logs -f mkadiri-golang-tester && \
	docker rm -f mkadiri-golang-tester