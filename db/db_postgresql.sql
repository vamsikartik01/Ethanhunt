create table Accounts (
    id serial,
    username varchar(255) not null,
    email varchar(255) unique,
    passwordHash varchar(255),
    saltId integer not null,
    dateCreated timestamp default current_timestamp,
    passwordChangedTime timestamp default current_timestamp,
    primary key(Id)
);

create table Salts (
    id serial,
    salt varchar(255),
    dateCreated timestamp default current_timestamp,
    primary key(id)
);

create table Rooms(
    id serial,
    name varchar(255),
    accountSid integer not null,
    primary key(id),
    foreign key(accountSid) references Accounts(id)
);

create table Hubs(
    id serial,
    name varchar(255),
    accountSid integer not null,
    roomId integer not null,
    status  varchar(16) check(status in ('offline','online')) not null default 'offline',
    primary key(id),
    foreign key(accountSid) references Accounts(id),
    foreign key(roomId) references Rooms(id)
);

create table Devices(
    id serial,
    name varchar(255),
    mode varchar(16) check(mode in ('output','input')) not null,
    status  varchar(16) check(status in ('offline','online')) not null default 'offline',
    value varchar(16) check(value in ('off','on')) not null default 'off',
    hubId integer not null,
    roomId integer not null,
    accountSid integer not null,
    foreign key(hubId) references Hubs(id),
    foreign key(roomId) references Rooms(id),
    foreign key(accountSid) references Accounts(id),
    primary key(id)
);

insert into Devices(name, mode, hubId, roomId, accountSid) values('device 1','output',1,1,1);