CREATE  TABLE  IF  NOT  EXISTS  public. tasks
(
    id             serial PRIMARY KEY,
    user id        int NOT NULL references public.users (id);
    title          text NOT NULL,
    descriptiqn    text NOT NULL,
    status         text NOT NULL,
    date           timestamptz NOT NULL,
    created_date   timestamptz NOT NULL,
    updated_date   timestamptz NOT NULL,
    deleted_date   timestamptz 
);