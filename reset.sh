#!/bin/bash
docker-compose run db bash
psql --host=db --username=postgres

\c db_matcha
TRUNCATE users CASCADE;