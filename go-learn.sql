--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Homebrew)
-- Dumped by pg_dump version 14.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: peoples; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.peoples (
    id integer NOT NULL,
    name character varying NOT NULL,
    age integer NOT NULL,
    address character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.peoples OWNER TO postgres;

--
-- Name: peoples_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.peoples ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.peoples_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: peoples; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.peoples (id, name, age, address, created_at, updated_at) FROM stdin;
1	Taufan Samudra	26	Jakarta	0001-01-01 07:07:12+07:07:12	2023-01-23 20:59:37.282653+07
\.


--
-- Name: peoples_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.peoples_id_seq', 2, true);


--
-- Name: peoples peoples_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.peoples
    ADD CONSTRAINT peoples_pk PRIMARY KEY (id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: taufansa
--

GRANT ALL ON SCHEMA public TO postgres;


--
-- PostgreSQL database dump complete
--

