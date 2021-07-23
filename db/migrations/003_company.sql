-- +goose Up
insert into company (name, termination_amount, expiration_date)
values ('Компания 1', 30000, '2022-01-01');

-- +goose Down

