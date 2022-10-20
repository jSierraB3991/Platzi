build:
	sudo docker build . -t postgres-platzi-go

run:
	sudo docker run -p 5433:5432 --name postgres-platzi -d postgres-platzi-go

all: build run
