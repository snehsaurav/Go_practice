--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3
-- Dumped by pg_dump version 13.3

-- Started on 2021-07-06 01:09:55 IST

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

--
-- TOC entry 7 (class 2615 OID 16394)
-- Name: pgagent; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA pgagent;


ALTER SCHEMA pgagent OWNER TO postgres;

--
-- TOC entry 3391 (class 0 OID 0)
-- Dependencies: 7
-- Name: SCHEMA pgagent; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA pgagent IS 'pgAgent system tables';


--
-- TOC entry 2 (class 3079 OID 16384)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 3392 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


--
-- TOC entry 3 (class 3079 OID 16395)
-- Name: pgagent; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgagent WITH SCHEMA pgagent;


--
-- TOC entry 3393 (class 0 OID 0)
-- Dependencies: 3
-- Name: EXTENSION pgagent; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgagent IS 'A PostgreSQL job scheduler';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 221 (class 1259 OID 16677)
-- Name: story; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.story (
    id integer NOT NULL,
    story_id integer,
    title character varying,
    sentences character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.story OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16675)
-- Name: fruits_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fruits_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fruits_id_seq OWNER TO postgres;

--
-- TOC entry 3394 (class 0 OID 0)
-- Dependencies: 220
-- Name: fruits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fruits_id_seq OWNED BY public.story.id;


--
-- TOC entry 3228 (class 2604 OID 16680)
-- Name: story id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.story ALTER COLUMN id SET DEFAULT nextval('public.fruits_id_seq'::regclass);


--
-- TOC entry 3179 (class 0 OID 16396)
-- Dependencies: 205
-- Data for Name: pga_jobagent; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--

INSERT INTO pgagent.pga_jobagent (jagpid, jaglogintime, jagstation) VALUES (354, '2021-07-05 20:22:45.544753+05:30', 'SNEH-MACBOOK-PRO');


--
-- TOC entry 3180 (class 0 OID 16407)
-- Dependencies: 207
-- Data for Name: pga_jobclass; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3181 (class 0 OID 16419)
-- Dependencies: 209
-- Data for Name: pga_job; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3183 (class 0 OID 16471)
-- Dependencies: 213
-- Data for Name: pga_schedule; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3184 (class 0 OID 16501)
-- Dependencies: 215
-- Data for Name: pga_exception; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3185 (class 0 OID 16516)
-- Dependencies: 217
-- Data for Name: pga_joblog; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3182 (class 0 OID 16445)
-- Dependencies: 211
-- Data for Name: pga_jobstep; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3186 (class 0 OID 16533)
-- Dependencies: 219
-- Data for Name: pga_jobsteplog; Type: TABLE DATA; Schema: pgagent; Owner: postgres
--



--
-- TOC entry 3385 (class 0 OID 16677)
-- Dependencies: 221
-- Data for Name: story; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.story (id, story_id, title, sentences, created_at, updated_at) VALUES (61, 1, 'verloop verloop', ' verloop verloop verloop verloop verloop verloop verloop verloop verloop verloop verloop verloop verloop verloop verloop', '2021-07-05 22:38:46.861988', '2021-07-05 22:38:46.861989');
INSERT INTO public.story (id, story_id, title, sentences, created_at, updated_at) VALUES (62, 1, NULL, 'verloop', '2021-07-05 22:38:55.795427', '2021-07-05 22:38:55.795428');


--
-- TOC entry 3395 (class 0 OID 0)
-- Dependencies: 220
-- Name: fruits_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fruits_id_seq', 62, true);


--
-- TOC entry 3253 (class 2606 OID 16685)
-- Name: story id_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.story
    ADD CONSTRAINT id_pkey PRIMARY KEY (id);


-- Completed on 2021-07-06 01:09:55 IST

--
-- PostgreSQL database dump complete
--

