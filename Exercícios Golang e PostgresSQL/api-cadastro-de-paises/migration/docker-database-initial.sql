--Create the schema
create schema api_countries;

--Create the tables
drop table api_countries.countries;
create table api_countries.countries (
	id serial primary key,
	name varchar(255) not null,
	acronym varchar(255) not null
);

drop table api_countries.states;
create table api_countries.states (
	id serial primary key,
	name varchar(255) not null, 
	acronym varchar(255) not null,
	country_id integer not null references api_countries.countries(id) on delete cascade
);

drop table api_countries.cities;
create table api_countries.cities (
	id serial primary key,
	name varchar(255) not null,
	state_id integer not null references api_countries.states(id) on delete cascade
);

--Inserting countries
insert into api_countries.countries(name, acronym)
values
('Brasil', 'BR'),
('Estados Unidos', 'EUA'),
('Alemanha', 'DE');

--inserting states
insert into api_countries.states(name, acronym, country_id)
values 
('Acre', 'AC', 1),
('Alagoas', 'AL', 1),
('Espírito Santo', 'ES', 1),
('Goiás', 'GO', 1),
('Alasca', 'AK', 2),
('Alabama', 'Al', 2),
('Arizona', 'AZ', 2),
('Califórnia', 'CA', 2),
('Baviera', 'BY', 3),
('Schleswig-Holstein', 'SH', 3),
('Baixa Saxônia', 'NI', 3),
('Hamburgo', 'HH', 3);

--inserting cities
insert into api_countries.cities(name, state_id)
values
('Rio Branco', 1),
('Cruzeiro do Sul', 1),
('Maceió', 2),
('Palmeira dos Índios', 2),
('Vitória', 3),
('Serra', 3),
('Goiânia', 4),
('Goiás', 4),
('Juneau', 5),
('North Pole', 5),
('Montgomery', 6),
('Auburn', 6),
('Phoenix', 7),
('Sedona', 7),
('Los Angeles', 8),
('São Francisco', 8),
('Munique', 9),
('Lindau', 9),
('Elshorn', 10),
('Eutin', 10),
('Hanôver', 11),
('Celle', 11),
('Asfred', 12),
('Hexs', 12);