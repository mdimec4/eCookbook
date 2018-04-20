# eCookbook

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
 
 ## Allrecipes middleware
 
 Cookbook has the ability to process allrecepies.com recipes by simply inserting alrecipes.com URL or decipe ID. For this it relies on external webservice called https://github.com/mdimec4/allrecipes. You need to set ``ALLRECIPES_MIDDLEWARE_HOST_PORT`` accordingly.  

## Docker
 Use provided ```Dockerfile``` to buld docker image.
 To build for ARM, you need to first generate ```Dockerfile_arm``` with ```generate_Dockerfile_arm.sh```.


```sh
docker run -d --restart=always -e POSTGRES_PASSWORD=chef -e POSTGRES_USER=chef -e POSTGRES_DB=cookbook --name eCookbook_postgres postgres
docker run -d --restart=always -e CHEF_DB_HOST=db2_1 -p 4006:4006 --link eCookbook_postgres:db2_1 --name eCookbook mihad/ecookbook
```

Pre build docker image is also avaliable on [dockerhub](https://hub.docker.com/r/mihad/ecookbook/]), but it is not neceserily up to date.
