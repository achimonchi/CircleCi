build:
	docker build -t achimonchi/circleci .

build-and-run: build run test

run:
	make remove
	docker run -d --name belajar-circleci -p 4000:4000 achimonchi/circleci

test:
	docker run --name go-test -v $(PWD):/ws \
		--workdir=/ws \
     	golang:1.12.1 \
   		/bin/bash -c "ls "
	docker rm -f go-test

remove: 
	docker rm -f belajar-circleci