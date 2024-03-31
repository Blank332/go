CREATE TABLE car
(
    guid UUID DEFAULT uuid_generate_v4() NOT NULL,
    name character varying NOT NULL,
    brand character varying NOT NULL,
    type character varying,
    year integer,
    description character varying,
    CONSTRAINT pk_car PRIMARY KEY (guid)
)
WITH (
    OIDS=FALSE
);


ALTER TABLE car OWNER TO postgres;
