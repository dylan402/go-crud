CREATE TABLE IF NOT EXISTS public.todos (
	id serial primary key not null,
	title text not null,
	description text not null,
	done bool not null default false
);
