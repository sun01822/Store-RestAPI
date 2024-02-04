create database store;
use store;


select * from Products;
select * from users;
SELECT * FROM `users` WHERE id = 1 AND `users`.`deleted_at` IS NULL;

select * from Addresses;
select * from geo_locations;
select * from names;

drop table addresses;
drop table users;
drop table names;
drop table geo_locations;

update users
set name_id=0
where id = 1;