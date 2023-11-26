create table users
(
    id_user         UUID primary key default gen_random_uuid(),
    name       varchar(100) not null,
    same_info varchar(100)

) ;

create table slug
(
    id_slug        UUID primary key default gen_random_uuid(),
    name_slug      varchar(100) not null,
    CONSTRAINT slugnickunicum UNIQUE (name_slug)

) ;

create table slugtraker
(
    id       UUID primary key default gen_random_uuid(),
    id_user uuid  REFERENCES users (id_user),
    id_slug uuid  REFERENCES slug (id_slug),
    CONSTRAINT newtrakerconstrain UNIQUE  (id_user,id_slug)

) ;




