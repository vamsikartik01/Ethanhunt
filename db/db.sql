#create database jack_db;
use jack_db;

CREATE TABLE Accounts (
    id INTEGER NOT NULL AUTO_INCREMENT UNIQUE,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    passwordHash VARCHAR(255),
    saltId INTEGER NOT NULL,
    dateCreated DATETIME DEFAULT CURRENT_TIMESTAMP,
    passwordChangedTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (Id)
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

update Accounts set email = "demo@email.com"  where id = 1;

CREATE TABLE Preferences (
    id INT NOT NULL AUTO_INCREMENT,
    accountSid INT NOT NULL,
    note TEXT,
    name VARCHAR(255) DEFAULT 'Bengaluru',
    city VARCHAR(255) DEFAULT 'Karnataka',
    state VARCHAR(255) DEFAULT 'IN',
    lat VARCHAR(255) default '12.9767936',
    lon VARCHAR(255) default '77.590082',
    PRIMARY KEY (id),
    FOREIGN KEY (accountSid) REFERENCES Accounts(id)
);

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
    refId varchar(255) not null,
    primary key(id),
    foreign key(accountSid) references Accounts(id)
);
create table Devices(
	id integer not null unique auto_increment,
    name varchar(255),
    mode enum('output','input') not null,
    status enum('offline','online') not null default 'offline',
    type varchar(255),
    value boolean default false,
    isFavorite boolean default false,
    hubPort integer not null,
    hubId integer,
    roomId integer not null,
    accountSid integer not null,
    foreign key(hubId) references Hubs(id),
    foreign key(roomId) references Rooms(id),
    foreign key(accountSid) references Accounts(id),
    primary key(id)
);
insert into Rooms(name, accountSid) values('Room 1', 1);
insert into Hubs(name, accountSid, refId) values('Hub 1',1,'facsghsfbjm,dsbfjsdf');
insert into Devices(name, mode, hubPort, hubId, roomId, accountSid) values('device 1','output',1,1,1,1);
select * from Accounts;
select * from Salts;
select * from Rooms;
select * from Hubs;
select * from Devices;