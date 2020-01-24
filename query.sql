DROP TABLE ServerService.Server;
DROP TABLE ServerService.Item;
DROP TABLE ServerService.Domain;


CREATE DATABASE IF NOT EXISTS ServerService;

CREATE TABLE IF NOT EXISTS ServerService.Domain(
    IdDomain SERIAL PRIMARY KEY,
    ServersChanged BOOL,
    SSLGrade STRING,
    PreviusSSLGrade STRING,
    Logo STRING,
    Title STRING,
    IsDown BOOL,
    Time TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS ServerService.Server(
    IdServer SERIAL PRIMARY KEY,
    Address STRING,
    SSLGrade STRING,
    Country STRING,
    Owner STRING,

    IdDomain SERIAL REFERENCES ServerService.Domain(IdDomain) 
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS ServerService.Item(
	IdItem SERIAL PRIMARY KEY,
	Site STRING,
	IdDomain SERIAL REFERENCES ServerService.Domain(IdDomain)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);