CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE purchase (
    purchase_id uuid default uuid_generate_v4() not null primary KEY,
    purchase_user VARCHAR not null,
    purchase_company VARCHAR not null,
    purchase_department VARCHAR not null,
    purchase_item VARCHAR not null,
    purchase_prquantity int not null,
    purchase_price int not null,
    purchase_code text not null,
    is_deleted bool default false not null,
    created_at timestamp default now() not null,
    created_by VARCHAR not null,
    updated_at timestamp null ,
    updated_by VARCHAR null
);
create unique index purchase_purchase_id_uindex
    on purchase (purchase_id);


create or replace function purchase_code()
returns text as
$$
declare
old_code text := (select purchase_code from purchase order by purchase_code desc limit 1);
    id_number char(4) := '0001';
    new_code text;
    num integer;
begin
    if old_code is null then
        new_code := 'P'||id_number;
return new_code;
end if;


        num := cast(right(old_code, 4) as integer) + 1;
        id_number :=
        case
            when num < 10 then '000'||num
            when num < 100 then '00'||num
            when num < 1000 then '0'||num
            when num < 10000 then cast(num as text)
end;

    new_code := 'P'||id_number;
return new_code;
end;
$$
language plpgsql;