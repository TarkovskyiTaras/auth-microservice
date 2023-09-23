
create table users (
                       id INTEGER primary key,
                       phone_number varchar(255) not null,
                       password      varchar(255) not null
);

create table refresh_tokens(
                               id            serial not null unique,
                               user_id       int references users (id) on delete cascade not null,
                               token         varchar(255) not null,
                               expires_at    timestamp not null
);

INSERT INTO users (id, phone_number, password) VALUES (123, '0933115485', 'password');
SELECT id FROM users WHERE phone_number='0933115485' AND password='password';
