\c rinha;

create table payments (
  correlation_id uuid primary key,
  amount integer not null,
  created_at timestamp not null
);
