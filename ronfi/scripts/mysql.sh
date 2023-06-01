#!/bin/bash

#HOST="172.31.85.248"
#PORT="3306"
#USER="bsc"
#PASS="ronfi"
HOST="176.9.120.196"
PORT="3306"
USER="bsc"
PASS="ronfi"


DATABASE="rkdb-eth"

newDB(){
  mysql -h $HOST -P $PORT -u$USER -p$PASS -e "
    drop database if exists $DATABASE;
    create database $DATABASE;
    use $DATABASE;
    create table if not exists loops(id int not null auto_increment, loopsId char(66) not null unique, path varchar(1024), poolFee char(32), tokenFee char(32), direction char(32), indexes char(32), counts int, canceled tinyint(1) not null default 0, hasV3 tinyint(1) not null default 0, primary key(id));
    create table if not exists obsall(id int not null auto_increment, tx char(66) not null unique, obsId char(1), loops varchar(2048), primary key(id));
    create table if not exists pair_dir_gas(id int not null auto_increment, pairDir char(44) not null unique, gas int, primary key(id));
    create table if not exists dex_pairs(id int not null auto_increment, pair char(44) not null unique, frequency int, primary key(id));
    create table if not exists obs_routers(id int not null auto_increment, router char(44) not null unique, methodID int unsigned, primary key(id));
    create table if not exists obs_methods(id int not null auto_increment, methodID char(20) not null unique, obsInfo char(10), primary key(id));
    create table if not exists pairs(id int not null auto_increment, pair char(44) not null unique, name char(40), pairIndex int, bothBriToken tinyint(1), keyToken char(44), token0 char(44), token1 char(44), factory char(44), primary key(id));
    create table if not exists pools(id int not null auto_increment, pool char(44) not null unique, name char(40), token0 char(44), token1 char(44), fee int, tickSpacing int, primary key(id));
    create table if not exists tokens(id int not null auto_increment, token char(44) not null unique, symbol char(20), decimals int, primary key(id));
  "
}



newDB
