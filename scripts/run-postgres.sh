#!/bin/bash
docker run --name postgres-wgblog -p 5432:5432 -e POSTGRES_DB=postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres postgres:16.1