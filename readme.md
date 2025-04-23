## Start cassandra DB

```
podman run --name cassandra -p 9042:9042 -p 9160:9160 cassandra:latest
```

## start csqlsh
```
podman exec -it cassandra cqlsh
```
## create schema

```
CREATE KEYSPACE hotel WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TYPE hotel.address (
    street text,
    city text,
    state_or_province text,
    postal_code text,
    country text
);


CREATE TABLE hotel.hotels (
    id text PRIMARY KEY,
    address frozen<address>,
    name text,
    phone text,
    pois set<text>
)
```

## sample data

```
INSERT INTO hotel.hotels ( id , address , name , phone ) VALUES (  'mirador', { street : 'Rte des Alpes 10', city: 'Montreux', country: 'Switzerland' }, 'Hotel Mirador','1234' );
```

