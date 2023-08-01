create database jackryan;
use jackryan;
show tables;
create table Accounts (
	id integer not null auto_increment unique,
    username varchar(255) not null,
    email varchar(255) unique,
    passwordHash varchar(255),
    saltId integer not null,
    dateCreated datetime default current_timestamp,
    passwordChangedTime datetime default current_timestamp,
    primary key(Id)
);

create table Salts (
	id integer not null auto_increment unique,
    salt varchar(255),
    dateCreated datetime default current_timestamp,
    primary key(id)
);

alter table Accounts
add foreign key(saltId) references Salts(id);

insert into Salts(salt) values('salt1');
insert into Accounts(username, passwordHash, saltId) values('demo','salthash1',1);

show tables;
select * from Salts;
select * from Accounts;

update Accounts set email = "demo@email.com"  where id = 1;

create table Rooms(
	id integer not null unique auto_increment,
    name varchar(255),
    accountSid integer not null,
    primary key(id),
    foreign key(accountSid) references Accounts(id)
);

create table Hubs(
	id integer not null unique auto_increment,
    name varchar(255),
    accountSid integer not null,
    roomId integer not null,
    primary key(id),
    foreign key(accountSid) references Accounts(id),
    foreign key(roomId) references Rooms(id)
);


create table Devices(
	id integer not null unique auto_increment,
    name varchar(255),
    mode enum('output','input') not null,
    status enum('offline','online') not null default 'offline',
    value binary,
    hubId integer not null,
    roomId integer not null,
    accountSid integer not null,
    foreign key(hubId) references Hubs(id),
    foreign key(roomId) references Rooms(id),
    foreign key(accountSid) references Accounts(id),
    primary key(id)
);

insert into Rooms(name, accountSid) values('Room 1', 1);
insert into Hubs(name, accountSid, roomId) values('Hub 1',1,1);
insert into Devices(name, mode, value, hubId, roomId, accountSid) values('device 1','output',false,1,1,1);

select * from Accounts;
select * from Salts;
select * from Rooms;
select * from Hubs;
select * from Devices;


