docker pull postgres:14
docker create -p 5433:5432 -e POSTGRES_PASSWORD=postgres --name credit_card_management postgres:14