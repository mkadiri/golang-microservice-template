BASE_DIRECTORY=${PWD}
GOLANG_BUILDER={{cookiecutter.dockerhub_org}}/golang-builder
GOLANG_TESTER={{cookiecutter.dockerhub_org}}/golang-tester
GOLANG_MICROSERVICE={{cookiecutter.dockerhub_org}}/{{cookiecutter.repo_name}}
PROJECT_URL={{cookiecutter.repo_base_name}}/{{cookiecutter.repo_name}}

build:
	docker build -t ${GOLANG_MICROSERVICE} .

test:
	docker-compose -f docker-compose-test.yml up -d && \
  docker logs -f golang-tester && \
  docker rm -f golang-tester && \
	docker rm -f mysql

run:
	docker-compose down && docker-compose up -d && docker logs -f {{cookiecutter.repo_name}}

build-test:
	cd ./docker/tester && docker build -t {{cookiecutter.dockerhub_org}}/golang-tester .

run-tests:
	docker-compose -f docker-compose-test.yml up -d && \
	docker logs -f mkadiri-golang-tester && \
	docker rm -f mkadiri-golang-tester