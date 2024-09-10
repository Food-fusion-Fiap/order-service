# Food Fusion - Tech Challenge - Order Microservice
Microsserviço responsável pela parte de pedidos e produtos.

# Como rodar localmente
- `docker compose up` para rodar a base (postgres e pgadmin) (arquivo docker-compose.yml)
- em src/infra/db/gorm/gorm.go, commente a string de produção e descomente a string local
- rode a aplicação pela IDE ou pelo comando `go run ./main.go`
