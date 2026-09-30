<div align="center">

<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="90" alt="Go">

<h1>Go HTTP JSON API</h1>

<p>
A simple REST API built with Go's standard <code>net/http</code> package and JSONPlaceholder.
</p>


<p>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  </a>
  <a href="https://github.com/claudiovictors/go-http-json-api">
    <img src="https://img.shields.io/github/license/claudiovictors/go-http-json-api?style=for-the-badge" alt="License">
  </a>
  <a href="https://github.com/claudiovictors/go-http-json-api/commits/main">
    <img src="https://img.shields.io/github/last-commit/claudiovictors/go-http-json-api?style=for-the-badge" alt="Last Commit">
  </a>
  <a href="https://github.com/claudiovictors/go-http-json-api">
    <img src="https://img.shields.io/github/repo-size/claudiovictors/go-http-json-api?style=for-the-badge" alt="Repository Size">
  </a>
  <a href="https://github.com/claudiovictors/go-http-json-api/stargazers">
    <img src="https://img.shields.io/github/stars/claudiovictors/go-http-json-api?style=for-the-badge" alt="GitHub Stars">
  </a>
</p>

<p>
  <a href="https://github.com/claudiovictors">
    <img src="https://img.shields.io/badge/GitHub-claudiovictors-181717?style=flat-square&logo=github&logoColor=white" alt="GitHub">
  </a>
  <a href="https://linkedin.com/in/claudio-victor-710986291">
    <img src="https://img.shields.io/badge/LinkedIn-Cláudio%20Victor-0A66C2?style=flat-square&logo=linkedin&logoColor=white" alt="LinkedIn">
  </a>
  <a href="https://instagram.com/claudio_victor.dev">
    <img src="https://img.shields.io/badge/Instagram-claudio__victor.dev-E4405F?style=flat-square&logo=instagram&logoColor=white" alt="Instagram">
  </a>
</p>

</div>

## Overview

This project is a small REST API built with Go's standard library.

It was created to practice building HTTP APIs without external web frameworks, focusing on HTTP routing, request handlers, JSON encoding and decoding, path parameters, HTTP status codes, and communication with external APIs.

The application uses JSONPlaceholder as its external data source.

## Technologies

* Go
* `net/http`
* `encoding/json`
* JSON
* REST API
* HTTP

## API Endpoints

### Get All Posts

```http
GET /posts
```

```bash
curl http://localhost:8080/posts
```

Returns all available posts from JSONPlaceholder.

### Get a Post

```http
GET /posts/{id}
```

Example:

```bash
curl http://localhost:8080/posts/1
```

### Create a Post

```http
POST /posts
Content-Type: application/json
```

Example:

```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"Learning Go","body":"Building an HTTP API with Go","userId":1}'
```

Example response:

```json
{
  "title": "Learning Go",
  "body": "Building an HTTP API with Go",
  "userId": 1,
  "id": 101
}
```

## Running Locally

Clone the repository:

```bash
git clone https://github.com/claudiovictors/go-http-json-api.git
cd go-http-json-api
```

Run the server:

```bash
go run .
```

The API will be available at:

```text
http://localhost:8080
```

## Concepts Practiced

* HTTP servers with `net/http`
* HTTP routing with `http.NewServeMux`
* HTTP handlers
* Path parameters
* JSON decoding with `json.Decoder`
* JSON encoding with `json.Marshal`
* HTTP headers
* HTTP status codes
* Request and response bodies
* Making HTTP requests from Go
* Consuming external REST APIs

## Important Note

This project uses JSONPlaceholder as a fake REST API for testing and development.

Posts created through `POST /posts` are simulated and are not permanently stored. Therefore, a post created through the API cannot be expected to remain available through a later `GET /posts/{id}` request.

## Future Improvements

* PostgreSQL persistence
* Request validation
* Middleware
* Authentication
* Automated tests
* Pagination
* Filtering
* API documentation
* Repository layer
* Service layer
* Environment-based configuration