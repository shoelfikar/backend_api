create table users (
    id int primary key auto_increment not null,
    username varchar(100) unique not null,
    password varchar(255),
    email varchar(254) unique not null,
    status int default 1,
    created_at timestamp default now(),
    created_by varchar(50) default 'system'
);

insert into users (username,password,email,created_by) values('superadmin', '$2a$10$D8PSZgwq/MuQeNebW96qGuwVyg3PMgy9qlXqADgDIF9LsRlA3Fm3u', 'admin@superadmin.com', 'system') 
-- pass superadmin
