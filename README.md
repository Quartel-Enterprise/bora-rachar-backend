# Backoffice

## Configurando Ambiente de Desenvolvimento

### Inicializando bancos SQL

Basta executar o banco em um daemon docker, não precisa popular dados manualmente.

```
docker run -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=bora_rachar -p 3306:3306 -d mysql --default-authentication-plugin=mysql_native_password
```

### Populando o banco de dados

É possível pré-popular o banco com alguns dados e as tabelas necessárias para a aplicação:

```
npm i
npm run build
npm run db:migrate
npm run db:seed
```

### Executando a API

```
npm i
npm run dev
```

## Diagrama do Banco de Dados

-- TODO: Add image

## Extensão

1. Para criar uma nova migration:

```
npm run db:migration:create
```

# Project Structure

### Structure for backoffice-api:

```
    .
    ├── build                     # Compiled files (alternatively `dist`)
    └── src                       # Source files
```

> Source files

```
  src
    ├── controllers               # Controllers to handle the API request
    ├── database                  # Databases and data access logic
    ├── domain                    # Domain objects of database
    ├── mapper                    # Maps a domain object to database format and vice versa
    ├── routes                    # API definitions
    ├── types                     # Generic types
    ├── useCases                  # UseCase to handle API request
    └── utils                     # Utils functions/class
```

> Data access logic

```
  database
    ├── migrations                # Version control of database
    ├── models                    # A model represent a table
    ├── repository                # Persistence and retrieval data
    └── seeders                   # Database defauld (fake)
```
