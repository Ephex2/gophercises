# URL Shortener; Gophercise number 2
This project is meant to be a URL shortener, or more accurately, a URL redirector. It can take config files that are json,
yaml, or hard-coded variables that will build a mapping between pathes on the web server and URLs. 

Note: the yaml files must be in the same format as the testyaml.yml file (only path entries with url subentries, all below a pairs header).
Ditto for json; look at the test.json file for the pattern to follow.

flags:
- -yamlpath is used to specify the path to a yaml file with a similar format to testyaml.yml
- -jsonpath is used to specify the path to a json file with a similar format to test.json
- -sql is used to use a local postgres SQL server.


## Note on using the -sql flag:
The program expects a local variable called DATABASE_URL that contains a connection string to the postgres database.

The database must have a table called ```pairs``` which contains two rows, ```path``` and ```url```.
Here is the create script generated from pgAdmin. Note that I was running a local instance and that the table was created in the ```public``` schema:

```sql
CREATE TABLE IF NOT EXISTS public.pairs
(
    path character varying(40) COLLATE pg_catalog."default" NOT NULL,
    url character varying(200) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pairs_pkey PRIMARY KEY (path)
)
```

Of course, feel free to change the type of the columns, so long as they can still hold the desired strings ( I am relatively new to postgres as of this writing ).
  
