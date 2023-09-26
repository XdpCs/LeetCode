# @Title        1757.Recyclable-and-Low-Fat-Products.sql
# @Description  1757.Recyclable-and-Low-Fat-Products solution
# @Create       XdpCs 2023-09-25 19:12
# @Update       XdpCs 2023-09-25 19:12
# sql schema
Create table If Not Exists Products
(
    product_id int,
    low_fats   ENUM ('Y', 'N'),
    recyclable ENUM ('Y','N')
);
Truncate table Products;
insert into Products (product_id, low_fats, recyclable)
values ('0', 'Y', 'N');
insert into Products (product_id, low_fats, recyclable)
values ('1', 'Y', 'Y');
insert into Products (product_id, low_fats, recyclable)
values ('2', 'N', 'Y');
insert into Products (product_id, low_fats, recyclable)
values ('3', 'Y', 'Y');
insert into Products (product_id, low_fats, recyclable)
values ('4', 'N', 'N');

# answer

select product_id
from Products
where low_fats = 'Y'
  and recyclable = 'Y';