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
- Database: MySQL (It was Dgraph. After trying to use it for a while 
and can't get a good result, I have decided to use a different database.
I would try to use it on this project when more documentation about Dgraph release).
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

1. Make sure you have installed all Go dependencies with ```go install```.
2. Run ```make letsgo``` for generate CSV and JSON files from endpoints.

![Generated](https://i.ibb.co/ZLGq0Xj/jsoncsv.gif)

3. Now you have started our backend server and uploaded data to MySQL database.

![Database](https://i.ibb.co/rcmxfcB/database.png)

4. If your URL to database is different, change ```tester:@tcp(localhost:3306)/``` 
on line 30 api-rest/database/database.go with your own URL and user for MySQL.  
5. Return to [Frontend](https://github.com/zejiran/frontend-simple-buyers-platform#usage),
for setting up UI and see our web app up and running.

## Docker Commands Used
-docker build -t computational-infrastructure/backend-simple-buyers-platform:1.1 .

-docker-compose -f deploy.yaml up

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](LICENSE)**
- Copyright 2020 © Juan Alegría.

<img src='https://i.ibb.co/sWSrvyF/logo.png' width="40%" alt="Logo">
