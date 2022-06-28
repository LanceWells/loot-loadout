FROM migrate/migrate:4
WORKDIR /migrations
ADD /migrations .
