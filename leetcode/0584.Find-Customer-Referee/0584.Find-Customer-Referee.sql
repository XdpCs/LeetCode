# @Title        0584.Find-Customer-Referee.sql
# @Description  0584.Find-Customer-Referee solution
# @Create       XdpCs 2023-09-25 19:12
# @Update       XdpCs 2023-09-25 19:12

# sql schema
Create table If Not Exists Customer
(
    id         int,
    name       varchar(25),
    referee_id int
);
Truncate table Customer;
insert into Customer (id, name, referee_id)
values ('1', 'Will', 'None');
insert into Customer (id, name, referee_id)
values ('2', 'Jane', 'None');
insert into Customer (id, name, referee_id)
values ('3', 'Alex', '2');
insert into Customer (id, name, referee_id)
values ('4', 'Bill', 'None');
insert into Customer (id, name, referee_id)
values ('5', 'Zack', '1');
insert into Customer (id, name, referee_id)
values ('6', 'Mark', '2');

# answer
select name
from Customer
where referee_id != 2
   or referee_id is null;