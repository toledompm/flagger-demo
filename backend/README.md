# Backend

This is the backend of our application. It's written in Go and provides a simple API with a single route.

## Getting Started

To run the backend locally, you'll need to have Go installed on your machine. Once you have Go installed, you can start the server by running the following command in the `backend` directory:

```
go run main.go
```

This will start the server on `localhost:8080`.

## API

The backend provides a single route at `/api/color` that returns a random color name in the format:

```json
{
  "color": "red"
}
```
