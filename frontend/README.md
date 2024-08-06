# Frontend

This is the frontend part of the monorepo project. It is a simple React application that periodically sends requests to a backend API and displays a green or red bubble based on the response.

## Setup

To install the necessary dependencies, run:

```
npm install
```

## Running the Application

To start the application, run:

```
npm start
```

The application will be available at `http://localhost:3000`.

## Structure

- `public/index.html`: Entry point of the application. Includes a root div where the React app is injected.
- `src/App.js`: Main component of the application. Contains logic to periodically send requests to the backend API and display bubbles based on the response.
- `src/index.js`: Renders the main App component into the root div in `index.html`.
- `src/components/Bubble.js`: Exports a Bubble component that takes a color prop and renders a bubble of that color.

## Dependencies

The main dependencies of the frontend application are:

- `react`: A JavaScript library for building user interfaces.
- `axios`: Promise based HTTP client for the browser and node.js. Used to send requests to the backend API.

For a full list of dependencies, refer to `package.json`.