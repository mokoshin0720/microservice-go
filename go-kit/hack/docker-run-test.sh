docker network create we-mental || true

docker-compose up -d mysql-test
docker-compose run --rm dockerize -wait tcp://mysql-test:3306 -timeout 20s
docker-compose run --rm migration-test

docker-compose up -d cognito-idp
docker-compose run --rm dockerize -wait tcp://cognito-idp:5001 -timeout 20s
