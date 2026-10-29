create table posts (
  id bigint primary key auto_increment,
  title varchar(200) not null,
  content text not null,
  category varchar(200),
  created_at timestamp not null default now(),
  updated_at timestamp,
  status enum ('publish', 'draft', 'trash') not null
);

