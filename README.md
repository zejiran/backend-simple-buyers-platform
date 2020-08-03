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

1. Install [Docker](https://www.docker.com/) and [jq](https://stedolan.github.io/jq/).
2. ```cd api-rest```.
2. ```make dgraph``` for install packages and run Dgraph.
3. On another terminal run ```npm run inject-dgraph-schema```.
4. Our GraphQL native database is ready to use.
Run queries in a tool like [GraphQL Playground](https://legacy.graphqlbin.com/new)
with ```http://localhost:9000/graphql``` endpoint.

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](LICENSE)**
- Copyright 2020 © Juan Alegría.

<img src='https://i.ibb.co/sWSrvyF/logo.png' width="40%" alt="Logo">
