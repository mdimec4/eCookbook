docker_tag=

ifdef DOCKER_TAG
	docker_tag=:$(DOCKER_TAG)
endif

dockerfile='Dockerfile'

all: docker
docker:
	tar -czh ../bin \
		$(dockerfile) \
		../cmd \
		../spa\
		../Makefile \
	| docker build -t chef$(docker_tag) -

modify_for_arm: 
	cp ./Dockerfile ./Dockerfile_arm
	sed -i 's/ubuntu:/armv7\/armhf-ubuntu:/' docker/Dockerfile_arm
	dockerfile='Dockerfile_arm'


arm: modify_for_arm docker

