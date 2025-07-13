\c rinha;

create table payments (
  correlation_id uuid primary key,
  amount decimal(10, 2) not null,
  created_at timestamp not null
);
