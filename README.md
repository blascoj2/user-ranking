# User Ranking Service

## SOLUTION

The challenge is solved following **Domain Driven Design** and **Hexagonal Architecture**.

Project skeleton is layered following Hexagonal architecture pattern:
```shell
├── cmd: app entrypoint
├── internal
    ├── application: use cases that orchestrate bussiness logic
    ├── domain: domain entities that contains all the bussiness logic
    └── infrastructure: connection with third parties or repositories
    └── ui: http or message brokers controllers 
```
## API

The interface provided by the service is a RESTfull API. The operations are as follows.

### POST /user/{user_id}/score

**Body** _required_ The user score to be loaded. Can be partial or total. If there's a partial score and we want to update the current user score, user must exists.

**Content Type** `application/json`

Sample:

```json
  {
    "total": 4
  }
```
```json
  {
    "score": "+100"
  }
```
Responses:

* **200 OK** When the score is registered correctly.
* **400 Bad Request** When there is a failure in the request format, expected headers, or the payload can't be unmarshalled.
* **404 NotFound** When we want to update partial score and user doesn't exists.
### Get /ranking?type="top100"

Fetch ranking depending on the ranking type requested.

**Query** _required_ Ranking type

Sample:

```
Top ranking type: GET /ranking?type="top100"
Top ranking type: Get /ranking?type="at100/3"
```

Responses:

* **200 OK** Ranking list.
```json
  [{
    "user_id": "1",
    "score": 100
  }]
```
* **400 Bad Request** When there is a failure in the request format, expected headers, or the payload can't be unmarshalled.

### System Requirements:

1. Operating System : Any(Windows/Linux/Mac)
2. Golang 1.19
4. Testify & Mockgen for testing

### How to run the application?

This application can be easily run in Docker building Dockerfile placed in root folder:

Change directory with the path to your application
```shell
cd /path/to/user-ranking
```
Build the container image

```shell
docker build -t user-ranking .
```

Start your container using the docker run command specifying the mapping ports.
In this case you can acces to the application in http://localhost:8080

```shell
docker run -dp 8080:8080 user-ranking
```
