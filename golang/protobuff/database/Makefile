build:
	sudo docker build . -t postgres-platzi-go

run:
	sudo docker run -p 5433:5432 --name postgres-platzi -e POSTGRES_USER=postgres -e POSTGRES_DB=platzi_go -e POSTGRES_PASSWORD=root -d postgres-platzi-go

all: build run
