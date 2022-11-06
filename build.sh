docker pull postgres:14
docker create -p 5433:5432 -e POSTGRES_PASSWORD=postgres --name credit_card_management_db postgres:14
docker build -t golang_backend .
docker create -p 8888:8888 --add-host=host.docker.internal:host-gateway --name credit_card_management_backend golang_backend