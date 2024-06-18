#!/bin/bash

psql -U $PGUSER -d $PGDATABASE -f ./server/test/mock_data/mock_data.sql
