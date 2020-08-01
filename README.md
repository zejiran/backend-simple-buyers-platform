# Backend Simple Buyers Platform

> “From emotions to materials, it's all about buying and selling”
― Mehnaz Ansari

## Index

- [Overview](#overview)
- [Used Technologies](#used-technologies)
- [Features](#featuresendpoints)
    - [Load data to database](#1-load-data-to-database)
    - [List buyers](#2-list-buyers)
    - [Query buyers](#3-query-buyers)
- [Usage](#usage)
- [License](#license)

## Overview

API REST used by 
<a href='https://github.com/zejiran/frontend-simple-buyers-platform'>zejiran/frontend-simple-buyers-platform</a>
that is a simple buyer platform on Vue.js.

## Used Technologies

- Language: Go.
- Database: Dgraph.
- API Router: ```chi```.
- Interface: Vue.js & Vuetify.

## Features/endpoints

### 1. Load data to database:

- Allows loading data from an endpoint based on a specific date.
- Download, process and store in a local database. 
- By default, loads actual day.

##### Data to be load:

1. List of products of the day.
2. List of buyers of the day.
3. List of transactions of the day.

### 2. List buyers:

- List all people that have bought on the platform.
- Use a local database with loaded data, day does not matter.

### 3. Query buyers:

Gets ID of buyer and return:

- Shopping history.
- Other buyers using same IP.
- Some product recommendations that people also bought.

## Usage

1. Install <a href='https://www.docker.com/'>Docker</a>.
2. ```docker pull dgraph/dgraph:v20.03.0``` for grab last Dgraph version.
3. ```mkdir -p ~/dgraph``` for storing Dgraph data.
4. Run Dgraph in Docker:
```
# Run Dgraph zero
docker run -it -p 5080:5080 -p 6080:6080 -p 8080:8080 \
  -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph --name dgraph \
  dgraph/dgraph:v20.03.0 dgraph zero

# In another terminal, now run Dgraph alpha
docker exec -it dgraph dgraph alpha --lru_mb 2048 --zero localhost:5080 --whitelist 0.0.0.0/0

# And in another, run ratel (Dgraph UI)
docker exec -it dgraph dgraph-ratel
```

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](LICENSE)**
- Copyright 2020 © Juan Alegría.

<img src='https://i.ibb.co/sWSrvyF/logo.png' width="40%" alt="Logo">
