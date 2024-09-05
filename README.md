# talentpitch

## Tabla de contenido

* [Configuración](#Configuración)
    * [Pre-requisitos ](#Pre-requisitos)
    * [Instalación](#Instalación)
* [Pruebas unitarias](#Pruebas_unitarias)

## Configuración 🚀

Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos
de desarrollo y pruebas.

### Pre-requisitos 📋

* El proyecto está desarrollado en el lenguaje de programación GO. Para poder utilizar este aplicativo es necesario
  instalar [Go.](https://golang.org/doc/install)

* El proyecto necesita una base de datos Postgres para facilidad puede usar docker es necesario
  instalar [Docker.](https://docs.docker.com/engine/install/)

* El proyecto necesita Make 
  instalar [Make.](https://www.gnu.org/software/make/)

### Instalación

* Clonar el repositorio

````
git clone https://github.com/alexjiec2695/talentpitch.git
````

* Para la ejecución del proyecto tenemos que correr el siguiente comando en la consola, este comando corre un
  docker-compose que configura y deja preparada la base de datos Postgres y levanta la aplicación

```
make run
```

## Pruebas unitarias

* Para correr las pruebas unitarias debemos de correr el siguiente comando

```
make test
```

## Curls

* Obtener todo

````
curl --location 'http://localhost:8081/users?page=1'

curl --location 'http://localhost:8081/challenges?page=1'

curl --location 'http://localhost:8081/videos?page=1'
````

* Crear

````
curl --location 'http://localhost:8081/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "",
    "email": ""
}'

curl --location 'http://localhost:8081/challenges' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "",
    "description": ""
}'

curl --location 'http://localhost:8081/videos' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "",
    "url": ""
}'
````

* Obtener detalle

````
curl --location 'http://localhost:8081/users/ID'

curl --location 'http://localhost:8081/challenges/ID'

curl --location 'http://localhost:8081/videos/ID'
````

* Actulizar

````
curl --location --request PUT 'http://localhost:8081/users/ID' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "",
    "email": ""
}'

curl --location --request PUT 'http://localhost:8081/challenges/ID' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "",
    "description": ""
}'

curl --location --request PUT 'http://localhost:8081/videos/ID' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "",
    "url": ""
}'
````

* Eliminar

````
curl --location --request DELETE 'http://localhost:8081/users/ID'

curl --location --request DELETE 'http://localhost:8081/challenges/ID'

curl --location --request DELETE 'http://localhost:8081/videos/ID'
````

## Construido con 🛠️

* [Go](https://golang.org/) - Lenguaje de programación base del proyecto Falcon.
* [Gin ](https://github.com/gin-gonic/gin) - Librería web usada para la definición de los endpoints REST.
* [Testify](https://github.com/stretchr/testify) - Librería que permite realizar pruebas unitarias.
* [Gorm](https://gorm.io/index.html) - Librería usada para acceder base de datos.
