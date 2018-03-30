# Chef

Cookbook on Visionect e-ink device. 

## Setup chef parametes using  Enviorment variables

 | variable                         | default value  |
 |----------------------------------|----------------|
 | CHEF_DB_USER                     | chef           |
 | CHEF_DB_NAME                     | cookbook       |
 | CHEF_DB_PASS                     | chef           |
 | CHEF_DB_SSLMODE                  | disable        |
 | CHEF_DB_HOST                     | localhost      |
 | CHEF_DB_PORT                     | 5432           |
 | CHEF_LISTEN_ADDR                 | :4006          |
 | ALLRECIPES_MIDDLEWARE_HOST_PORT  | localhost:4007 |

## Docker
 Use provided ```Dockerfile``` to buld docker image.
 To build for ARM, you need to first generate ```Dockerfile_arm``` with ```generate_Dockerfile_arm.sh```.


```sh
docker run -d --restart=always -e POSTGRES_PASSWORD=chef -e POSTGRES_USER=chef -e POSTGRES_DB=cookbook --name chef_postgres postgres
docker run -d --restart=always -e CHEF_DB_HOST=db2_1 -p 4006:4006 --link chef_postgres:db2_1 --name chef chef
```

