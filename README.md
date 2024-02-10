
# Astrology Service

Service for obtaining metadata and pictures of the day APOD and API server for obtaining all album entrie

## Features

- Receiving once a day metadata and picture of the day [APOD](https://api.nasa.gov/) and saving to binary storage
- Http API server for receiving all album entries and entries for the selected day;

## Tech Stack

- Golang
- PostgreSQL
- Docker
- Viper
- Gorilla/mux

## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod download
```

Start the application

```bash
  docker compose up
```

