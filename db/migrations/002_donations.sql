-- +goose Up
insert into donation_type (name, min_amount,description)
values ('300 рублей или более', 300, ''),
       ('400 рублей или более',400,''),
       ('1000 рублей или более', 1000,''),
       ('Другая сумма', 1,'');



-- +goose Down

