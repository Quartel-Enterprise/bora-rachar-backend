# bora-rachar-service
Microsservico responsável por fornecer informações para o app "bora-rachar"


comando para subir banco de dados local utilizando docker:  
```
docker-compose down -v
docker-compose up -d
```

caso alguma mudaça seja feita é necessario gerar gerar novamente as structs e services do swagger
```
oapi-codegen -generate types -o src/cmd/generated-code/openapi_types.gen.go -package generated swagger.yaml
oapi-codegen -generate chi-server -o src/cmd/generated-code/openapi_server.gen.go -package generated swagger.yaml
```