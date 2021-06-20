
-- +migrate Up
create table users(
    id uuid not null default uuid_generate_v4(),
    name text default null,
    password text not null,
    email citext not null constraint proper_email check (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    constraint pkey_tbl primary key (id),
    unique(email)
);
-- +migrate StatementBegin
create or replace function insert_user(
    _name text,
    _password text,
    _email citext
) returns uuid as $f$
declare
    _user_id uuid;
begin
    _user_id := uuid_generate_v4();
    insert into users(id, name, password, email) values(_user_id, _name, _password, _email);
    return _user_id;
end
$f$ language 'plpgsql';
-- +migrate StatementEnd
-- +migrate Down
