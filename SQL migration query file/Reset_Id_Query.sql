-- Column: public.events.id
ALTER TABLE public.events_views
DROP CONSTRAINT event_id;

ALTER TABLE public.rating
DROP CONSTRAINT event_id;

ALTER TABLE public.ticket
DROP CONSTRAINT event_id;

ALTER TABLE public.event_category
DROP CONSTRAINT event_id;

ALTER TABLE public.events
DROP COLUMN id;

-- ALTER TABLE IF EXISTS public.events DROP COLUMN IF EXISTS id;
DROP SEQUENCE IF EXISTS events_id_seq;
ALTER TABLE IF EXISTS public.events
DROP COLUMN IF EXISTS id;
ALTER TABLE IF EXISTS public.events
    ADD COLUMN id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 );
SELECT setval('events_id_seq', 1);

-- Assuming "events_pkey" is the primary key constraint name on the "events" table
ALTER TABLE public.events
ADD CONSTRAINT events_pkey PRIMARY KEY (id);

-- Now, add the foreign key constraints to other tables
ALTER TABLE public.events_views
ADD CONSTRAINT event_id FOREIGN KEY (event_id)
    REFERENCES public.events (id)
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE public.rating
ADD CONSTRAINT event_id FOREIGN KEY (event_id)
    REFERENCES public.events (id)
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE public.ticket
ADD CONSTRAINT event_id FOREIGN KEY (event_id)
    REFERENCES public.events (id)
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE public.event_category
ADD CONSTRAINT event_id FOREIGN KEY (event_id)
    REFERENCES public.events (id)
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;