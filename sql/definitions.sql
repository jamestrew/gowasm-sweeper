create table difficulties (
  id int primary key,
  description text
);


create table leaderboard (
  id bigint generated by default as identity primary key,
  name text not null,
  difficulty_id int not null references difficulties (id),
  time int not null,
  data jsonb,
  inserted_at timestamp with time zone default timezone('utc'::text, now()) not null
);

insert into difficulties (id, description)
values
  (0, 'beginner'),
  (1, 'intermediate'),
  (2, 'expert'),
  (3, 'custom');

insert into leaderboard (name, difficulty_id, time)
values
  ('john', 0, 12),
  ('jason', 0, 15),
  ('supabase', 0, 7),
  ('gopher', 0, 14),
  ('john', 1, 58),
  ('jason', 1, 98),
  ('foobar', 1, 111),
  ('foobar', 1, 154),
  ('smashing pumps', 2, 154),
  ('john', 2, 304),
  ('gopher', 2, 349),
  ('foobar', 2, 493),
  ('foobar', 2, 348);
