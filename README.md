# microservices-wallet
Some troubleshootings :
1.pq: function uuid_generate_v4() does not exist | Insert this sql query for all db -> CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

To deploy :
Create internal folder in the root directory and create into this files:

movement.env
DB_HOST=movements-service-db
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=movements
DB_PORT=5432

user.env
DB_HOST=user-service-db
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=virtualwallet
DB_PORT=5432
REDIS_ADDR=user-service-redis:6379
SECRECT_KEY=YOU_SECRET_KEY
MOVEMENT_GRPC=movements-service:9000
AUTH_GRPC=auth-service:9002
TRANSACTION_GRPC=transaction-service:9003

auth.env
DB_HOST=user-service-db
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=virtualwallet
DB_PORT=5432
REDIS_ADDR=user-service-redis:6379
SECRECT_KEY=YOU_SECRET_KEY
MOVEMENT_GRPC=movements-service:9000

transaction.env
DB_HOST=transaction-service-db
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=virtualwallet
DB_PORT=5432
REDIS_ADDR=transaction-service-redis:6379
SECRECT_KEY=YOU_SECRET_KEY
MOVEMENT_GRPC=movements-service:9000
USER_GRPC=user-service:9001

Open the terminal in the root directory:
docker-compose up
