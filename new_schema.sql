 
drop table if exists championships; 
create table championships (
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
description VARCHAR ( 255 ),
state INT,
Date_start TIMESTAMP,
date_end TIMESTAMP,
registration_end TIMESTAMP,
gender CHAR(1));

drop table if exists category;
create table category(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
age_start TIMESTAMP,
age_end TIMESTAMP);

drop table if exists event;
create table event(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
description VARCHAR ( 255 ),
type INT,
has_serie INT,
idcategory uuid,
gender VARCHAR(1),
family_type VARCHAR(50),
track INT,
final INT,
min_mark VARCHAR(10),
sub_type VARCHAR(20));

drop table if exists schools;
create table schools(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
bd VARCHAR(12),
name VARCHAR(200),
alias VARCHAR(25),
run VARCHAR(12),
rutholder VARCHAR(12),
gender VARCHAR(1),
state INT,
restriction INT,
idcommune uuid,
idprovince uuid,
idregion uuid);

drop table if exists users;
create table users(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
email varchar(100),
type varchar(25),
password varchar(255),
created_on TIMESTAMP,
last_login TIMESTAMP,
first_name varchar(100),
last_name varchar(100),
phone varchar(20),
rut varchar(12),
birthdate TIMESTAMP,
gender varchar(1),
nationality varchar(25),
graduate int);

drop table if exists usersschools;
create table usersschools(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
idusers uuid,
idschools uuid);

drop table if exists regions;
create table regions(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
ordinal varchar(25),
name varchar(50));

drop table if exists provinces;
create table provinces(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
name varchar(50),
idregion uuid);

drop table if exists communes;
create table communes(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
name varchar(50),
idprovince uuid);

drop table if exists registration;
create table registration(
id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
iduser uuid,
idscholl uuid,
idchampionships uuid,
idevent uuid,
status int,
headline int);

