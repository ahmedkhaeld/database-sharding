# Database Sharding
Database Sharding is a process of segmenting the data into partitions that are spread in multiple database
instances This is essentially to speed up query and scale the system.


the idea of traditionally how we do things is we
have one big database with a huge table
And we would you might have some indexes on those tables.
You query and then you get results.
It's always a centralized one database.
Your pin point is one server and querying becomes slower and slower the more data you have because what
gets what you are going to have indexes and those indexes are going to grow.
And as your tables grow or large and that kind of slow down your queries.
And you need more memory, you need more CPU.

* e.g. if I want to make a select URL from this table

<img src="https://user-images.githubusercontent.com/40498170/180182363-dfb82f3d-2642-428d-b49a-8cd708c933ef.png" height="300" width="600">

let's say this is like a URL shortener website, URL_table in the database
contains the urls and their codes, essentially what you do is give me the 
url for this code so i can visit that page<br>
you query that, that's a huge table to say there are one million rows
That's that query is going to get slower and slower, slower the more rows you have
even with indexes, that still is going to get slow.

#### How to speed things up?
using sharding, put the table into multiple databases<br>
split 1 million rows table into 5 database instances s1, s2, s3, s4, s5
* How do we do Shard? How do we distribute? <br>
  You can  do partition based on the zip code, like these are the zip code for the west coast
  This is a zip code for the East goes. Or if you have a bunch of users, maybe users from one to five hundred million goes to this database
  and from five hundred and one to one million goes to this database, that's another kind of sharding.

But what if your key is a string like this URLID = "5FTOJ"
Which database server is this 5FTOJ is on?. how do you know them?
<br> So that's where we're going to talk about **Consistent Hashing**

<img src="https://user-images.githubusercontent.com/40498170/180186310-fbe4c628-68b0-4529-9d77-c647be795ea6.png" width="600" height="300">

#### Consistent Hashing

<img src="https://user-images.githubusercontent.com/40498170/180189224-55175e83-5047-4076-a676-19c888974aaf.png" height="300" width="600">

the idea of consistent things you take an input or a string or any user provided piece of
data "INPUT" you want to query on, <br> and want to know which database instance to query on
essentially that hash("Input") gives back the instance somehow.
<br> every time we submit the input1 string, it will always go to the database instance 5432
and if we submit input2 we go to instance 5433, and input3 goes to 5434
<br> Consistently hashing that make sure consistently to the same server node

it is like take the input and give me the port number back so i connect to the database

---
### Demo: URL shortener
what is doing it shorten the url and then figure out which database to hit to write to that database

i. spin up Docker instances for the postgres, but in order to do that each docker instance that we are going to spin up
has to initialize the sql script that creates the url_table

spin up 3 postgres instances with identical schema  on ports: - 5434, 54322, 54323 when we spin the container instance 
we are going to  initialize the database with our URL table <br>
* write to the sharded databases<br>
* Reads from the sharded databases

`~/mkdir sharding`
`~/sharding$ touch init.sql`
1. create the SQL script that creates the table `init.sql` that gets executed every time we spin up a new docker container
```sql
CREATE TABLE URL_TABLE
(
id serial NOT NULL PRIMARY KEY,
URL text,
URL_ID character(5)
);
```
2. write a docker file because our random postgres image won't do. we
need to create our own image that essentially executes that script.
from the docker postgres instance write the image and 
copy the init.sql into a special folder called `/docker-entry-point-initdb.d`
the moment we copy that and spin up the container, that postgres image
will say, oh there a sql file here, I'm going to execute 
```dockerfile
FROM postgres
COPY init.sql /docker-entrypoint-initdb.d
```
3. build the image from the Dockerfile(our special image) which we are going
to spin up the containers<br>
   `~/sharding$ docker build -t pgshard .`



4. spin up the instances containers <br>
`docker run --name pgshard1 -p 5434:5432 -d -e POSTGRES_PASSWORD=postgres pgshard `<br>
`docker run --name pgshard2 -p 54322:5432 -d -e POSTGRES_PASSWORD=postgres pgshard`<br>
`docker run --name pgshard3 -p 54323:5432 -d -e POSTGRES_PASSWORD=postgres pgshard
   `<br>
we have now three database instances









