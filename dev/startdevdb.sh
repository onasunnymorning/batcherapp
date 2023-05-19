#!/bin/bash
docker run --rm --name postgresql-intro-gosamples -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=batch postgres