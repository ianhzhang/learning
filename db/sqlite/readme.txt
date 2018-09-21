sqlite3 test.db

create table family (id integer primary key, ln text, fn text)
insert into family(ln, fn) values("Zhang",  "Ian");
insert into family (ln, fn) values("Zhang", "Kevin");
insert into family (ln, fn) values("Zhang", "Kevin");
select * from family;
update family set fn="Lucas" where id=3;



delete from family where id=3;





http://www.sqlitetutorial.net/sqlite-autoincrement/