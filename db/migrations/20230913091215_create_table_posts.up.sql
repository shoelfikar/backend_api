create table posts(
    id int primary key auto_increment,
    title varchar(200) not null,
    content text not null,
    category varchar(100),
    created_date timestamp default now(),
    updated_date timestamp default now(),
    status varchar(100) default 'Draft'not null
)