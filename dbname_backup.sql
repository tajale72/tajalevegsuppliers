--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Homebrew)

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
-- Name: bill_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bill_details (
    id integer NOT NULL,
    bill_number character varying(255) NOT NULL,
    bill_date date NOT NULL,
    bill_total_amount character varying(255) NOT NULL,
    seller_name character varying(255) NOT NULL,
    seller_pan_num character varying(255) NOT NULL,
    customer_name character varying(255) NOT NULL,
    customer_location character varying(255) NOT NULL,
    customer_phone_number character varying(20) NOT NULL,
    customer_pan_container character varying(255) NOT NULL,
    items jsonb
);


ALTER TABLE public.bill_details OWNER TO postgres;

--
-- Name: bill_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bill_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bill_details_id_seq OWNER TO postgres;

--
-- Name: bill_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bill_details_id_seq OWNED BY public.bill_details.id;


--
-- Name: dailyvegetablesales; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dailyvegetablesales (
    id integer NOT NULL,
    vegetable_name character varying(255) NOT NULL,
    sale_date date NOT NULL,
    quantity_sold integer NOT NULL,
    rate numeric(10,2) NOT NULL,
    total_amount numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.dailyvegetablesales OWNER TO postgres;

--
-- Name: dailyvegetablesales_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dailyvegetablesales_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dailyvegetablesales_id_seq OWNER TO postgres;

--
-- Name: dailyvegetablesales_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dailyvegetablesales_id_seq OWNED BY public.dailyvegetablesales.id;


--
-- Name: ledger_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ledger_entries (
    id integer NOT NULL,
    date date NOT NULL,
    account character varying(255) NOT NULL,
    billnumber character varying(255) NOT NULL,
    debit numeric(10,2) DEFAULT 0.00,
    credit numeric(10,2) DEFAULT 0.00,
    balance_amount numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.ledger_entries OWNER TO postgres;

--
-- Name: ledger_entries_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ledger_entries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ledger_entries_id_seq OWNER TO postgres;

--
-- Name: ledger_entries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ledger_entries_id_seq OWNED BY public.ledger_entries.id;


--
-- Name: bill_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bill_details ALTER COLUMN id SET DEFAULT nextval('public.bill_details_id_seq'::regclass);


--
-- Name: dailyvegetablesales id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailyvegetablesales ALTER COLUMN id SET DEFAULT nextval('public.dailyvegetablesales_id_seq'::regclass);


--
-- Name: ledger_entries id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ledger_entries ALTER COLUMN id SET DEFAULT nextval('public.ledger_entries_id_seq'::regclass);


--
-- Data for Name: bill_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bill_details (id, bill_number, bill_date, bill_total_amount, seller_name, seller_pan_num, customer_name, customer_location, customer_phone_number, customer_pan_container, items) FROM stdin;
17	1	2081-05-13	135	तजले भेज सप्लायर्स	६०१०८६४८९	sdfsdf	sdfsdf	sdfdsf	455737378	[{"sa": "Green Onion (हरियो प्याज)", "dar": "10", "kra": "1", "rakam": "135", "sankhya": "13.5"}, {"sa": "", "dar": "", "kra": "2", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "3", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "4", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "5", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "6", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "7", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "8", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "9", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "10", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "11", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
18	2	2081-05-27	513.4	तजले भेज सप्लायर्स	६०१०८६४८९	Subani Tajale	Imadol	98415869914	123456794	[{"sa": "Pomegranate (अनार)", "dar": "10", "kra": "1", "rakam": "100", "sankhya": "10"}, {"sa": "Bitter Gourd (करेला)", "dar": "20", "kra": "2", "rakam": "410", "sankhya": "20.5"}, {"sa": "Banana (केरा)", "dar": "10", "kra": "3", "rakam": "3.4000000000000004", "sankhya": "0.34"}, {"sa": "", "dar": "", "kra": "4", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "5", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "6", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "7", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "8", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "9", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "10", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "11", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
\.


--
-- Data for Name: dailyvegetablesales; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dailyvegetablesales (id, vegetable_name, sale_date, quantity_sold, rate, total_amount, created_at) FROM stdin;
1	Potato (आलु)	2024-09-07	69	23.00	1587.00	2024-09-07 12:43:40.554384
2	Green Chillies (हरियो खुर्सानी)	2024-09-07	36	12.00	432.00	2024-09-07 12:43:40.556013
7	Peas (केराउ)	2024-09-07	12	12.00	144.00	2024-09-07 18:46:04.941989
8	Garlic (लसुन)	2024-09-07	15	15.00	225.00	2024-09-07 18:46:04.946357
9	Green Onion (हरियो प्याज)	2024-09-09	0	10.00	135.00	2024-09-09 22:51:46.793491
10	Pomegranate (अनार)	2024-09-11	10	10.00	100.00	2024-09-11 21:23:08.570346
11	Bitter Gourd (करेला)	2024-09-11	0	20.00	410.00	2024-09-11 21:23:08.581848
12	Banana (केरा)	2024-09-11	0	10.00	3.40	2024-09-11 21:23:08.583941
\.


--
-- Data for Name: ledger_entries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ledger_entries (id, date, account, billnumber, debit, credit, balance_amount, created_at) FROM stdin;
3	2024-09-06	asdsad	12	12.00	12.00	24.00	2024-09-07 23:31:06.077114
4	2024-09-10	subu	222	20.00	20.00	40.00	2024-09-08 21:29:49.030449
5	2024-09-13	ewweew	1213	2000.00	30000.00	32000.00	2024-09-08 21:37:52.335645
6	2024-09-18	sacs	124	12.00	12.00	24.00	2024-09-08 21:39:02.349857
7	2024-09-17	assadc	21	12.00	12.00	24.00	2024-09-08 21:42:32.024332
8	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:08.961266
9	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:24.051412
10	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:24.441843
11	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:24.891765
12	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:26.670511
13	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:27.030235
14	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:28.653084
15	2024-09-12	qwq	102	10.00	40.00	50.00	2024-09-11 16:39:29.274906
16	2024-09-12	qwq	102	10.25	40.24	50.00	2024-09-11 16:42:00.509498
\.


--
-- Name: bill_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bill_details_id_seq', 18, true);


--
-- Name: dailyvegetablesales_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dailyvegetablesales_id_seq', 12, true);


--
-- Name: ledger_entries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ledger_entries_id_seq', 16, true);


--
-- Name: bill_details bill_details_bill_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bill_details
    ADD CONSTRAINT bill_details_bill_number_key UNIQUE (bill_number);


--
-- Name: bill_details bill_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bill_details
    ADD CONSTRAINT bill_details_pkey PRIMARY KEY (id);


--
-- Name: dailyvegetablesales dailyvegetablesales_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailyvegetablesales
    ADD CONSTRAINT dailyvegetablesales_pkey PRIMARY KEY (id);


--
-- Name: dailyvegetablesales dailyvegetablesales_vegetable_name_sale_date_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dailyvegetablesales
    ADD CONSTRAINT dailyvegetablesales_vegetable_name_sale_date_key UNIQUE (vegetable_name, sale_date);


--
-- Name: ledger_entries ledger_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ledger_entries
    ADD CONSTRAINT ledger_entries_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

