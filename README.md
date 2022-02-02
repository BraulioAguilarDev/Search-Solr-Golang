# Search Engine

A simple search REST API using Apache Solr & Golang

## Quick install

We need to create our envs

```sh
$ cp .env.example .env
$ cp config/config.example.yml config.yml
```

## Test code

Starting service in Docker

```sh
# Build docker image
$ make docker

# Execute containers
$ make dc-up

# Upload data info
$ make seed 
```

## Running
`http://localhost:9090/api/v1/search?category=dev`


## Endpoints

### Search

`GET: /api/v1/search?locality=remoto&category=dev`

Successfull

```json
{
    "success": true,
    "data": [
        {
            "title": "Lead Application Engineer",
            "text": "We are hiring a Lead Engineer to drive the development of our Golang Core API... ",
            "salary": 60000,
            "locality": "Remoto",
            "category": "dev"
        },
        {
            "title": "Data Engineer",
            "text": "A Big Four is currently seeking a  Senior Associate  in Digital Lighthouse...",
            "salary": 55000,
            "locality": "Remoto - Solo México",
            "category": "data"
        },
        {
            "title": "Golang Developer",
            "text": "We are looking for a great Go developer who possesses a strong understanding of..",
            "salary": 50000,
            "locality": "Queretaro",
            "category": "dev"
        },
        {
            "title": "GoLang Engineer",
            "text": "We are seeking a GoLang Engineer who will be working on the modernization of the our platform...",
            "salary": 80000,
            "locality": "México",
            "category": "dev"
        }
    ]
}
```

Error

```json
{
    "success": false,
    "message": "failed"
}
```