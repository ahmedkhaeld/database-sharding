# 🗄️ Database Sharding

Welcome to the Database Sharding repository! This project demonstrates database sharding using Go as the back-end language.

## 📋 Table of Contents
- Introduction
- Features
- Installation
- Usage

## 🌟 Introduction
Database sharding is a process of segmenting data into partitions spread across multiple database instances to speed up queries and scale systems.

Traditionally, a single large database with huge tables and indexes leads to slow queries as data grows. Sharding solves this by distributing the data.

**Example: URL Shortener**

For a URL shortener website, the URL_table contains URLs and their codes. As the table grows, queries slow down despite indexes.

**How to speed things up?**

Using sharding, split the table into multiple databases. For instance, a table with 1 million rows can be split into 5 database instances (s1, s2, s3, s4, s5).

**Consistent Hashing**

Consistent hashing helps determine which database instance to query by hashing the input (e.g., URL ID) to consistently route to the same server node.

## ✨ Features
- Efficiently manage large datasets with sharding
- Written in Go for high performance
- Example implementation included

## 🚀 Installation
To get started with the Database Sharding project, follow these steps:

1. Clone the repository:
```
git clone https://github.com/ahmedkhaeld/database-sharding.git
cd database-sharding

```
2. Install dependencies:

```
go mod download

```
3. Run the project:
```
go run main.go
```

## Setting Up Sharded Databases with Docker

1. Create the SQL script init.sql:

```sql
   
CREATE TABLE URL_TABLE (
  id serial NOT NULL PRIMARY KEY,
  URL text,
  URL_ID character(5)
);
```

2. Create a Dockerfile:
```
FROM postgres
COPY init.sql /docker-entrypoint-initdb.d
```
3. docker build -t pgshard .
```
docker build -t pgshard .
```

4. Spin up the instances:
```bash
docker run --name pgshard1 -p 5434:5432 -d -e POSTGRES_PASSWORD=postgres pgshard
docker run --name pgshard2 -p 54322:5432 -d -e POSTGRES_PASSWORD=postgres pgshard
docker run --name pgshard3 -p 54323:5432 -d -e POSTGRES_PASSWORD=postgres pgshard

```

## 📚 Usage
Once the application is running, you can explore the sharding functionalities through the provided endpoints and configurations.

