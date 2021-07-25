CREATE TABLE kurs.bank_notes (
                                 id SERIAL PRIMARY KEY,
                                 symbol varchar NOT NULL,
                                 sell double precision NOT NULL,
                                 buy double precision NOT NULL,
                                 indexing_date date NOT NULL
);

CREATE TABLE kurs.e_rates (
                              id SERIAL PRIMARY KEY,
                                 symbol varchar NOT NULL,
                                 sell double precision NOT NULL,
                                 buy double precision NOT NULL,
                                 indexing_date date NOT NULL
);

CREATE TABLE kurs.tt_counter (
                                 id SERIAL PRIMARY KEY,
                                 symbol varchar NOT NULL,
                                 sell double precision NOT NULL,
                                 buy double precision NOT NULL,
                                 indexing_date date NOT NULL
);
