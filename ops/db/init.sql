create database if not exists clean_architecture_go_v2;
use clean_architecture_go_v2;
create table if not exists user (id varchar(50),email varchar(255),password varchar(255),first_name varchar(100), last_name varchar(100), created_at datetime, updated_at datetime, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
create table if not exists book (id varchar(50),title varchar(255),author varchar(255), pages integer,quantity integer, created_at datetime, updated_at datetime, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
create table if not exists book_user (user_id varchar(50),book_id varchar(50), created_at datetime, PRIMARY KEY (`user_id`,`book_id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;