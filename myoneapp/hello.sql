--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

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
-- Name: balance_sheet; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.balance_sheet (
    id integer NOT NULL,
    account_type character varying(50) NOT NULL,
    account_name character varying(255) NOT NULL,
    debit numeric(10,2),
    credit numeric(10,2),
    balance numeric(10,2) NOT NULL,
    transaction_date date NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.balance_sheet OWNER TO postgres;

--
-- Name: balance_sheet_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.balance_sheet_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.balance_sheet_id_seq OWNER TO postgres;

--
-- Name: balance_sheet_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.balance_sheet_id_seq OWNED BY public.balance_sheet.id;


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
-- Name: balance_sheet id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.balance_sheet ALTER COLUMN id SET DEFAULT nextval('public.balance_sheet_id_seq'::regclass);


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
-- Data for Name: balance_sheet; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.balance_sheet (id, account_type, account_name, debit, credit, balance, transaction_date, created_at) FROM stdin;
\.


--
-- Data for Name: bill_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bill_details (id, bill_number, bill_date, bill_total_amount, seller_name, seller_pan_num, customer_name, customer_location, customer_phone_number, customer_pan_container, items) FROM stdin;
3	301	2081-05-29	60325	तजले भेज सप्लायर्स	६०१०८६४८९	khwopa  banquet 	kamalbinayak   	9860819308	609518573	[{"sa": "Cauliflower (काउली)", "dar": "100", "kra": "1", "rakam": "6000", "sankhya": "60"}, {"sa": "Radish (मूला)", "dar": "45", "kra": "2", "rakam": "1125", "sankhya": "25"}, {"sa": "Carrot (गाजर)", "dar": "120", "kra": "3", "rakam": "2400", "sankhya": "20"}, {"sa": "Cucumber (काक्रो)30", "dar": "90", "kra": "4", "rakam": "2700", "sankhya": "30"}, {"sa": "Garlic (लसुन)", "dar": "150", "kra": "5", "rakam": "450", "sankhya": "3"}, {"sa": "Chili (खुर्सानी)", "dar": "140", "kra": "6", "rakam": "420", "sankhya": "3"}, {"sa": "Dhaniya (धनियाँ)", "dar": "500", "kra": "7", "rakam": "1000", "sankhya": "2"}, {"sa": "Kagati (कागती)", "dar": "280", "kra": "8", "rakam": "1680", "sankhya": "6"}, {"sa": "Tomato (टमाटर)", "dar": "90", "kra": "9", "rakam": "2250", "sankhya": "25"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "140", "kra": "10", "rakam": "2100", "sankhya": "15"}, {"sa": "Beet (चकुन्दर)", "dar": "100", "kra": "11", "rakam": "200", "sankhya": "2"}, {"sa": "Salad Patta (सलाद पत्ता)", "dar": "700", "kra": "12", "rakam": "700", "sankhya": "1"}, {"sa": "Golo Chyau (गोलो च्याउ)", "dar": "550", "kra": "13", "rakam": "2750", "sankhya": "5"}, {"sa": "Mushroom (च्याउ)", "dar": "250", "kra": "14", "rakam": "11250", "sankhya": "45"}, {"sa": "Beans (सिमी)", "dar": "200", "kra": "15", "rakam": "1000", "sankhya": "5"}, {"sa": "Lapsi Mada (लप्सी माडा)", "dar": "200", "kra": "16", "rakam": "1000", "sankhya": "5"}, {"sa": "Rayo Saag (रायो साग)", "dar": "250", "kra": "17", "rakam": "12500", "sankhya": "50"}, {"sa": "Onion (प्याज)", "dar": "100", "kra": "18", "rakam": "10800", "sankhya": "108"}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
4	302	2082-03-30	61920	तजले भेज सप्लायर्स	६०१०८६४८९	Khwopa  Banquet	Kamalbinayak Bhaktapur	9860819308	609518573	[{"sa": "Potato (आलु)", "dar": "70", "kra": "1", "rakam": "21140", "sankhya": "302"}, {"sa": "Cauliflower (काउली)", "dar": "100", "kra": "2", "rakam": "5000", "sankhya": "50"}, {"sa": "Radish (मूला)", "dar": "45", "kra": "3", "rakam": "900", "sankhya": "20"}, {"sa": "Carrot (गाजर)", "dar": "110", "kra": "4", "rakam": "1650", "sankhya": "15"}, {"sa": "Cucumber (काक्रो)", "dar": "120", "kra": "5", "rakam": "1800", "sankhya": "15"}, {"sa": "Garlic (लसुन)", "dar": "120", "kra": "6", "rakam": "720", "sankhya": "6"}, {"sa": "Chili (खुर्सानी)", "dar": "140", "kra": "7", "rakam": "420", "sankhya": "3"}, {"sa": "Dhaniya (धनियाँ)", "dar": "400", "kra": "8", "rakam": "1200", "sankhya": "3"}, {"sa": "Kagati (कागती)", "dar": "260", "kra": "9", "rakam": "1040", "sankhya": "4"}, {"sa": "Tomato (टमाटर)", "dar": "130", "kra": "10", "rakam": "1300", "sankhya": "10"}, {"sa": "Small Tomato (सानो टमाटर)", "dar": "80", "kra": "11", "rakam": "2000", "sankhya": "25"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "150", "kra": "12", "rakam": "2250", "sankhya": "15"}, {"sa": "Salad Patta (सलाद पत्ता)", "dar": "700", "kra": "13", "rakam": "700", "sankhya": "1"}, {"sa": "Chyau Patey (च्याउ पाते)", "dar": "260", "kra": "14", "rakam": "13000", "sankhya": "50"}, {"sa": "Beans (सिमी)", "dar": "200", "kra": "15", "rakam": "1200", "sankhya": "6"}, {"sa": "Rayo Saag (रायो साग)", "dar": "230", "kra": "16", "rakam": "6900", "sankhya": "30"}, {"sa": "Pudina (पुदिना)", "dar": "100", "kra": "17", "rakam": "700", "sankhya": "7"}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
5	304	2081-03-31	76345	तजले भेज सप्लायर्स	६०१०८६४८९	KHWOPA  BANQUET	KAMALBINAYAK  BHAKTAPUR	9860819308	609518573	[{"sa": "Potato (आलु)", "dar": "70", "kra": "1", "rakam": "4060", "sankhya": "58"}, {"sa": "Cauliflower (काउली)", "dar": "90", "kra": "2", "rakam": "9000", "sankhya": "100"}, {"sa": "Cabbage (बन्दगोभी)", "dar": "50", "kra": "3", "rakam": "500", "sankhya": "10"}, {"sa": "Radish (मूला)", "dar": "45", "kra": "4", "rakam": "1575", "sankhya": "35"}, {"sa": "Carrot (गाजर)", "dar": "110", "kra": "5", "rakam": "4400", "sankhya": "40"}, {"sa": "Cucumber (काक्रो)", "dar": "90", "kra": "6", "rakam": "4050", "sankhya": "45"}, {"sa": "Onion (प्याज)", "dar": "105", "kra": "7", "rakam": "11235", "sankhya": "107"}, {"sa": "Garlic (लसुन)", "dar": "120", "kra": "8", "rakam": "360", "sankhya": "3"}, {"sa": "Chili (खुर्सानी)", "dar": "140", "kra": "9", "rakam": "420", "sankhya": "3"}, {"sa": "Ginger (अदुवा)", "dar": "280", "kra": "10", "rakam": "4200", "sankhya": "15"}, {"sa": "Garlic (लसुन)", "dar": "250", "kra": "11", "rakam": "5375", "sankhya": "21.5"}, {"sa": "Kagati (कागती)", "dar": "280", "kra": "12", "rakam": "1680", "sankhya": "6"}, {"sa": "Small Tomato (सानो टमाटर)", "dar": "80", "kra": "13", "rakam": "3840", "sankhya": "48"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "140", "kra": "14", "rakam": "2100", "sankhya": "15"}, {"sa": "Salad Patta (सलाद पत्ता)", "dar": "600", "kra": "15", "rakam": "600", "sankhya": "1"}, {"sa": "Golo Chyau (गोलो च्याउ)", "dar": "450", "kra": "16", "rakam": "6750", "sankhya": "15"}, {"sa": "Beans (सिमी)", "dar": "200", "kra": "17", "rakam": "2000", "sankhya": "10"}, {"sa": "Rayo Saag (रायो साग)", "dar": "230", "kra": "18", "rakam": "11500", "sankhya": "50"}, {"sa": "Rayo Saag (रायो साग)", "dar": "270", "kra": "19", "rakam": "2700", "sankhya": "10"}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
6	305	2081-04-01	27825	तजले भेज सप्लायर्स	६०१०८६४८९	KHWOPA  BANQUET	KAMALBINAYAK  BHAKTAPUR	9860819308	609518573	[{"sa": "Cauliflower (काउली)", "dar": "100", "kra": "1", "rakam": "4000", "sankhya": "40"}, {"sa": "Cabbage (बन्दगोभी)", "dar": "50", "kra": "2", "rakam": "750", "sankhya": "15"}, {"sa": "Radish (मूला)", "dar": "45", "kra": "3", "rakam": "675", "sankhya": "15"}, {"sa": "Carrot (गाजर)", "dar": "110", "kra": "4", "rakam": "1980", "sankhya": "18"}, {"sa": "Cucumber (काक्रो)", "dar": "90", "kra": "5", "rakam": "1350", "sankhya": "15"}, {"sa": "Garlic (लसुन)", "dar": "130", "kra": "6", "rakam": "390", "sankhya": "3"}, {"sa": "Chili (खुर्सानी)", "dar": "140", "kra": "7", "rakam": "280", "sankhya": "2"}, {"sa": "Kagati (कागती)", "dar": "260", "kra": "8", "rakam": "520", "sankhya": "2"}, {"sa": "Small Tomato (सानो टमाटर)", "dar": "80", "kra": "9", "rakam": "1600", "sankhya": "20"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "105", "kra": "10", "rakam": "1050", "sankhya": "10"}, {"sa": "Salad Patta (सलाद पत्ता)", "dar": "600", "kra": "11", "rakam": "600", "sankhya": "1"}, {"sa": "Chyau Patey (च्याउ पाते)", "dar": "250", "kra": "12", "rakam": "7500", "sankhya": "30"}, {"sa": "Hariyo Simi (हरियो सिमी)", "dar": "220", "kra": "13", "rakam": "880", "sankhya": "4"}, {"sa": "Rayo Saag (रायो साग)", "dar": "250", "kra": "14", "rakam": "6250", "sankhya": "25"}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
7	310	2081-05-07	70400	तजले भेज सप्लायर्स	६०१०८६४८९	KHWOPA  BANQUET	KAMALBINAYAK	9860819308	609518573	[{"sa": "Potato (आलु)", "dar": "70", "kra": "1", "rakam": "35350", "sankhya": "505"}, {"sa": "Cauliflower (काउली)", "dar": "140", "kra": "2", "rakam": "2100", "sankhya": "15"}, {"sa": "Radish (मूला)", "dar": "40", "kra": "3", "rakam": "400", "sankhya": "10"}, {"sa": "Carrot (गाजर)", "dar": "120", "kra": "4", "rakam": "8400", "sankhya": "70"}, {"sa": "Cucumber (काक्रो)", "dar": "110", "kra": "5", "rakam": "1100", "sankhya": "10"}, {"sa": "Garlic (लसुन)", "dar": "110", "kra": "6", "rakam": "110", "sankhya": "1"}, {"sa": "Onion (प्याज)", "dar": "120", "kra": "7", "rakam": "6240", "sankhya": "52"}, {"sa": "Green Chillies (हरियो खुर्सानी)", "dar": "120", "kra": "8", "rakam": "120", "sankhya": "1"}, {"sa": "Dhaniya (धनियाँ)", "dar": "120", "kra": "9", "rakam": "840", "sankhya": "7"}, {"sa": "Ginger (अदुवा)5", "dar": "280", "kra": "10", "rakam": "1400", "sankhya": "5"}, {"sa": "Garlic (लसुन)", "dar": "260", "kra": "11", "rakam": "1300", "sankhya": "5"}, {"sa": "Kagati (कागती)", "dar": "280", "kra": "12", "rakam": "560", "sankhya": "2"}, {"sa": "Tomato (टमाटर)", "dar": "90", "kra": "13", "rakam": "900", "sankhya": "10"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "120", "kra": "14", "rakam": "240", "sankhya": "2"}, {"sa": "Beans (सिमी)", "dar": "140", "kra": "15", "rakam": "420", "sankhya": "3"}, {"sa": "Lapsi Mada (लप्सी माडा)", "dar": "100", "kra": "16", "rakam": "100", "sankhya": "1"}, {"sa": "Rayo Saag (रायो साग)", "dar": "140", "kra": "17", "rakam": "9800", "sankhya": "70"}, {"sa": "Apple (स्याउ)", "dar": "300", "kra": "18", "rakam": "300", "sankhya": "1"}, {"sa": "Peas (केराउ)", "dar": "120", "kra": "19", "rakam": "120", "sankhya": "1"}, {"sa": "Grapes (अंगुर)", "dar": "600", "kra": "20", "rakam": "600", "sankhya": "1"}]
12	309	2081-05-05	26276	तजले भेज सप्लायर्स	६०१०८६४८९	Yoyo Banquet	Suryabinayak	9841692211	123456789	[{"sa": "Potato (आलु)", "dar": "70", "kra": "1", "rakam": "3500", "sankhya": "50"}, {"sa": "Cauliflower (काउली)", "dar": "160", "kra": "2", "rakam": "2400", "sankhya": "15"}, {"sa": "Radish (मूला)", "dar": "50", "kra": "3", "rakam": "525", "sankhya": "10.5"}, {"sa": "Carrot (गाजर)", "dar": "110", "kra": "4", "rakam": "1540", "sankhya": "14"}, {"sa": "Dhaniya (धनियाँ) and lasun", "dar": "350", "kra": "5", "rakam": "700", "sankhya": "2"}, {"sa": "Lapsi Mada (लप्सी माडा)", "dar": "200", "kra": "6", "rakam": "600", "sankhya": "3"}, {"sa": "Cucumber (काक्रो)", "dar": "100", "kra": "7", "rakam": "1850", "sankhya": "18.5"}, {"sa": "Green Chillies (हरियो खुर्सानी)", "dar": "200", "kra": "8", "rakam": "300", "sankhya": "1.5"}, {"sa": "Onion (प्याज)", "dar": "120", "kra": "9", "rakam": "2400", "sankhya": "20"}, {"sa": "Ginger (अदुवा) garlic", "dar": "280", "kra": "10", "rakam": "1120", "sankhya": "4"}, {"sa": "Small Tomato (सानो टमाटर)", "dar": "90", "kra": "11", "rakam": "2115", "sankhya": "23.5"}, {"sa": "Tomato (टमाटर)", "dar": "120", "kra": "12", "rakam": "156", "sankhya": "1.3"}, {"sa": "Beans (सिमी)", "dar": "125", "kra": "13", "rakam": "500", "sankhya": "4"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "100", "kra": "14", "rakam": "200", "sankhya": "2"}, {"sa": "Kagati (कागती)", "dar": "280", "kra": "15", "rakam": "560", "sankhya": "2"}, {"sa": "Celery (अजमोदा)", "dar": "100", "kra": "16", "rakam": "100", "sankhya": "1"}, {"sa": "Mushroom (च्याउ)", "dar": "200", "kra": "17", "rakam": "3000", "sankhya": "15"}, {"sa": "pakucha saag", "dar": "220", "kra": "18", "rakam": "2310", "sankhya": "10.5"}, {"sa": "brinjal (bhanta)", "dar": "80", "kra": "19", "rakam": "200", "sankhya": "2.5"}, {"sa": "Broccoli (ब्रोकोली)", "dar": "220", "kra": "20", "rakam": "2200", "sankhya": "10"}]
13	309	2081-05-05	1280	तजले भेज सप्लायर्स	६०१०८६४८९	yoyo banquet	suryabinayak	984692211	123456789	[{"sa": "Mushroom (च्याउ)", "dar": "500", "kra": "1", "rakam": "1000", "sankhya": "2"}, {"sa": "Pumpkin (फर्सी)", "dar": "70", "kra": "2", "rakam": "280", "sankhya": "4"}, {"sa": "", "dar": "", "kra": "3", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "4", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "5", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "6", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "7", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "8", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "9", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "10", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "11", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
14	310	2081-05-07	18560	तजले भेज सप्लायर्स	६०१०८६४८९	Khwopa Banquet	Kamalbinayak	9860819308	609518573	[{"sa": "Potato (आलु)", "dar": "70", "kra": "1", "rakam": "3500", "sankhya": "50"}, {"sa": "Cauliflower (काउली)", "dar": "140", "kra": "2", "rakam": "2100", "sankhya": "15"}, {"sa": "Radish (मूला)", "dar": "40", "kra": "3", "rakam": "400", "sankhya": "10"}, {"sa": "Carrot (गाजर)", "dar": "120", "kra": "4", "rakam": "840", "sankhya": "7"}, {"sa": "Cucumber (काक्रो)", "dar": "110", "kra": "5", "rakam": "1100", "sankhya": "10"}, {"sa": "hariyo lasun", "dar": "200", "kra": "6", "rakam": "200", "sankhya": "1"}, {"sa": "Onion (प्याज)", "dar": "120", "kra": "7", "rakam": "3000", "sankhya": "25"}, {"sa": "Green Chillies (हरियो खुर्सानी)", "dar": "200", "kra": "8", "rakam": "200", "sankhya": "1"}, {"sa": "Dhaniya (धनियाँ)", "dar": "300", "kra": "9", "rakam": "300", "sankhya": "1"}, {"sa": "Ginger (अदुवा)", "dar": "280", "kra": "10", "rakam": "1400", "sankhya": "5"}, {"sa": "Garlic (लसुन)", "dar": "260", "kra": "11", "rakam": "1300", "sankhya": "5"}, {"sa": "Kagati (कागती)", "dar": "280", "kra": "12", "rakam": "560", "sankhya": "2"}, {"sa": "Tomato (टमाटर)", "dar": "90", "kra": "13", "rakam": "900", "sankhya": "10"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "120", "kra": "14", "rakam": "240", "sankhya": "2"}, {"sa": "Beans (सिमी)", "dar": "140", "kra": "15", "rakam": "420", "sankhya": "3"}, {"sa": "Lapsi Mada (लप्सी माडा)", "dar": "200", "kra": "16", "rakam": "100", "sankhya": "0.5"}, {"sa": "Rayo Saag (रायो साग)", "dar": "140", "kra": "17", "rakam": "980", "sankhya": "7"}, {"sa": "Apple (स्याउ)", "dar": "300", "kra": "18", "rakam": "300", "sankhya": "1"}, {"sa": "Banana (केरा)", "dar": "120", "kra": "19", "rakam": "120", "sankhya": "1"}, {"sa": "Grapes (अंगुर)", "dar": "600", "kra": "20", "rakam": "600", "sankhya": "1"}]
15	1	2081-05-01	3300	तजले भेज सप्लायर्स	६०१०८६४८९	Yoyo Banquet	Suryabinakyak	9841362211	123455678	[{"sa": "pakucha saag", "dar": "210", "kra": "1", "rakam": "3150", "sankhya": "15"}, {"sa": "saag", "dar": "150", "kra": "2", "rakam": "150", "sankhya": "1"}, {"sa": "", "dar": "", "kra": "3", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "4", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "5", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "6", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "7", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "8", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "9", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "10", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "11", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
16	2	2081-05-08	5250	तजले भेज सप्लायर्स	६०१०८६४८९	Khwopa Banquet	Kamalbinayak	9860819308	609518573	[{"sa": "Cauliflower (काउली)", "dar": "120", "kra": "1", "rakam": "1800", "sankhya": "15"}, {"sa": "Radish (मूला)", "dar": "40", "kra": "2", "rakam": "200", "sankhya": "5"}, {"sa": "Carrot (गाजर)", "dar": "120", "kra": "3", "rakam": "600", "sankhya": "5"}, {"sa": "Cucumber (काक्रो)", "dar": "110", "kra": "4", "rakam": "550", "sankhya": "5"}, {"sa": "lasun", "dar": "200", "kra": "5", "rakam": "100", "sankhya": "0.5"}, {"sa": "Green Chillies (हरियो खुर्सानी)", "dar": "160", "kra": "6", "rakam": "80", "sankhya": "0.5"}, {"sa": "Dhaniya (धनियाँ)", "dar": "400", "kra": "7", "rakam": "200", "sankhya": "0.5"}, {"sa": "Kagati (कागती)", "dar": "280", "kra": "8", "rakam": "280", "sankhya": "1"}, {"sa": "Tomato (टमाटर)", "dar": "90", "kra": "9", "rakam": "900", "sankhya": "10"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "100", "kra": "10", "rakam": "300", "sankhya": "3"}, {"sa": "Beans (सिमी)", "dar": "120", "kra": "11", "rakam": "240", "sankhya": "2"}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
18	186	2080-10-29	1014866	तजले भेज सप्लायर्स	६०१०८६४८९	khagi suppliers	kamavinayakl	8951068830	500001310	[{"sa": "Cabbage (बन्दगोभी)", "dar": "40", "kra": "1", "rakam": "13900", "sankhya": "347.5"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "100", "kra": "2", "rakam": "11950", "sankhya": "119.5"}, {"sa": "Carrot (गाजर)", "dar": "70", "kra": "3", "rakam": "41055", "sankhya": "586.5"}, {"sa": "Cucumber (काक्रो)", "dar": "60", "kra": "4", "rakam": "31470", "sankhya": "524.5"}, {"sa": "Cauliflower (काउली)", "dar": "65", "kra": "5", "rakam": "38967.5", "sankhya": "599.5"}, {"sa": "Beans (सिमी)", "dar": "60", "kra": "6", "rakam": "11700", "sankhya": "195"}, {"sa": "Chili (खुर्सानी)", "dar": "90", "kra": "7", "rakam": "3555", "sankhya": "39.5"}, {"sa": "spring garlic", "dar": "120", "kra": "8", "rakam": "14484", "sankhya": "120.7"}, {"sa": "Lapsi  (लप्सी )", "dar": "70", "kra": "9", "rakam": "4130", "sankhya": "59"}, {"sa": "Rayo Saag (रायो साग)", "dar": "40", "kra": "10", "rakam": "11960", "sankhya": "299"}, {"sa": "palak sag", "dar": "65", "kra": "11", "rakam": "18037.5", "sankhya": "277.5"}, {"sa": "Apple (स्याउ)", "dar": "270", "kra": "12", "rakam": "212760", "sankhya": "788"}, {"sa": "Potato (आलु)", "dar": "48", "kra": "13", "rakam": "190992", "sankhya": "3979"}, {"sa": "kino", "dar": "100", "kra": "14", "rakam": "246500", "sankhya": "2465"}, {"sa": "Banana (केरा)", "dar": "85", "kra": "15", "rakam": "59925", "sankhya": "705"}, {"sa": "Radish (मूला)", "dar": "40", "kra": "16", "rakam": "23680", "sankhya": "592"}, {"sa": "Chyau Patey (च्याउ पाते)", "dar": "160", "kra": "17", "rakam": "21600", "sankhya": "135"}, {"sa": "lemon", "dar": "150", "kra": "18", "rakam": "6000", "sankhya": "40"}, {"sa": "Ginger (अदुवा)", "dar": "170", "kra": "19", "rakam": "15300", "sankhya": "90"}, {"sa": "Broccoli (ब्रोकोली)", "dar": "100", "kra": "20", "rakam": "36900", "sankhya": "369"}]
19	186	2080-10-29	7126	तजले भेज सप्लायर्स	६०१०८६४८९	khagi suppliers	kamal vinayak	9851068830	500001310	[{"sa": "cardinier leavs", "dar": "140", "kra": "1", "rakam": "7126", "sankhya": "50.9"}, {"sa": "", "dar": "", "kra": "2", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "3", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "4", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "5", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "6", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "7", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "8", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "9", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "10", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "11", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
20	16	2081-05-31	768867.4	Tajle Veg Suppliers	६०१०८६४८९	khagi suppliers	bhaktapur	951068830	500001310	[{"sa": "Cabbage (बन्दगोभी)", "dar": "40", "kra": "1", "rakam": "15660", "sankhya": "391.5"}, {"sa": "Carrot (गाजर)", "dar": "70", "kra": "2", "rakam": "22295", "sankhya": "318.5"}, {"sa": "Radish (मूला)", "dar": "40", "kra": "3", "rakam": "18860", "sankhya": "471.5"}, {"sa": "cacumber", "dar": "60", "kra": "4", "rakam": "67230", "sankhya": "1120.5"}, {"sa": "lemon", "dar": "150", "kra": "5", "rakam": "6570", "sankhya": "43.8"}, {"sa": "Capsicum (क्याप्सिकम)", "dar": "100", "kra": "6", "rakam": "13769.999999999998", "sankhya": "137.7"}, {"sa": "Ginger (अदुवा)", "dar": "170", "kra": "7", "rakam": "23256.000000000004", "sankhya": "136.8"}, {"sa": "spring ginger", "dar": "120", "kra": "8", "rakam": "12516", "sankhya": "104.3"}, {"sa": "Green Chillies (हरियो खुर्सानी)", "dar": "90", "kra": "9", "rakam": "4365", "sankhya": "48.5"}, {"sa": "Beans (सिमी)", "dar": "60", "kra": "10", "rakam": "3330", "sankhya": "55.5"}, {"sa": "lapsi", "dar": "70", "kra": "11", "rakam": "4305", "sankhya": "61.5"}, {"sa": "Ramtoria (रामतोरिया)", "dar": "60", "kra": "12", "rakam": "27420", "sankhya": "457"}, {"sa": "Potato (आलु)", "dar": "48", "kra": "13", "rakam": "198734.40000000002", "sankhya": "4140.3"}, {"sa": "Apple (स्याउ)", "dar": "270", "kra": "14", "rakam": "300942", "sankhya": "1114.6"}, {"sa": "parwal", "dar": "60", "kra": "15", "rakam": "19200", "sankhya": "320"}, {"sa": "iskus", "dar": "35", "kra": "16", "rakam": "10500", "sankhya": "300"}, {"sa": "bodi", "dar": "60", "kra": "17", "rakam": "7440", "sankhya": "124"}, {"sa": "Cauliflower (काउली)", "dar": "65", "kra": "18", "rakam": "4160", "sankhya": "64"}, {"sa": "karela", "dar": "60", "kra": "19", "rakam": "5430", "sankhya": "90.5"}, {"sa": "cardinier", "dar": "140", "kra": "20", "rakam": "2884", "sankhya": "20.6"}]
21	17	2081-05-31	239508	तजले भेज सप्लायर्स	६०१०८६४८९	khagi suppliers	bhaktapur	9851068830	500001310	[{"sa": "Rayo Saag (रायो साग)", "dar": "40", "kra": "1", "rakam": "1680", "sankhya": "42"}, {"sa": "naspati", "dar": "75", "kra": "2", "rakam": "15450", "sankhya": "206"}, {"sa": "Pumpkin (फर्सी)", "dar": "35", "kra": "3", "rakam": "7000", "sankhya": "200"}, {"sa": "Pudina (पुदिना)", "dar": "1000", "kra": "4", "rakam": "500", "sankhya": "0.5"}, {"sa": "Tomato (टमाटर)", "dar": "60", "kra": "5", "rakam": "15450", "sankhya": "257.5"}, {"sa": "Banana (केरा)", "dar": "140", "kra": "6", "rakam": "167468", "sankhya": "1196.2"}, {"sa": "ghiraula", "dar": "40", "kra": "7", "rakam": "4280", "sankhya": "107"}, {"sa": "Chyau Patey (च्याउ पाते)", "dar": "160", "kra": "8", "rakam": "27680", "sankhya": "173"}, {"sa": "", "dar": "", "kra": "9", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "10", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "11", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "12", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "13", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "14", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "15", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "16", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "17", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "18", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "19", "rakam": "0", "sankhya": ""}, {"sa": "", "dar": "", "kra": "20", "rakam": "0", "sankhya": ""}]
\.


--
-- Data for Name: dailyvegetablesales; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dailyvegetablesales (id, vegetable_name, sale_date, quantity_sold, rate, total_amount, created_at) FROM stdin;
1	Potato (आलु)	2024-09-09	12	12.00	144.00	2024-09-09 20:19:03.925914
2	Green Onion (हरियो प्याज)	2024-09-09	5	25.00	260.00	2024-09-09 20:25:59.597238
6	Cauliflower (काउली)	2024-09-09	60	100.00	6000.00	2024-09-09 23:46:38.886093
7	Radish (मूला)	2024-09-09	25	45.00	1125.00	2024-09-09 23:46:38.886975
4	Carrot (गाजर)	2024-09-09	30	120.00	2900.00	2024-09-09 23:46:38.887798
9	Cucumber (काक्रो)30	2024-09-09	30	90.00	2700.00	2024-09-09 23:46:38.888582
10	Garlic (लसुन)	2024-09-09	3	150.00	450.00	2024-09-09 23:46:38.889231
11	Chili (खुर्सानी)	2024-09-09	3	140.00	420.00	2024-09-09 23:46:38.889895
12	Dhaniya (धनियाँ)	2024-09-09	2	500.00	1000.00	2024-09-09 23:46:38.89052
13	Kagati (कागती)	2024-09-09	6	280.00	1680.00	2024-09-09 23:46:38.891176
5	Tomato (टमाटर)	2024-09-09	40	90.00	3150.00	2024-09-09 23:46:38.891799
15	Capsicum (क्याप्सिकम)	2024-09-09	15	140.00	2100.00	2024-09-09 23:46:38.892444
16	Beet (चकुन्दर)	2024-09-09	2	100.00	200.00	2024-09-09 23:46:38.893088
17	Salad Patta (सलाद पत्ता)	2024-09-09	1	700.00	700.00	2024-09-09 23:46:38.893753
18	Golo Chyau (गोलो च्याउ)	2024-09-09	5	550.00	2750.00	2024-09-09 23:46:38.894376
19	Mushroom (च्याउ)	2024-09-09	45	250.00	11250.00	2024-09-09 23:46:38.895032
20	Beans (सिमी)	2024-09-09	5	200.00	1000.00	2024-09-09 23:46:38.895663
21	Lapsi Mada (लप्सी माडा)	2024-09-09	5	200.00	1000.00	2024-09-09 23:46:38.896474
22	Rayo Saag (रायो साग)	2024-09-09	50	250.00	12500.00	2024-09-09 23:46:38.897153
23	Onion (प्याज)	2024-09-09	108	100.00	10800.00	2024-09-09 23:46:38.897773
31	Dhaniya (धनियाँ)	2024-09-10	3	400.00	1200.00	2024-09-10 00:12:53.432446
33	Tomato (टमाटर)	2024-09-10	10	130.00	1300.00	2024-09-10 00:12:53.433804
40	Pudina (पुदिना)	2024-09-10	7	100.00	700.00	2024-09-10 00:12:53.437922
24	Potato (आलु)	2024-09-10	360	70.00	25200.00	2024-09-10 00:37:09.507652
82	Dhaniya (धनियाँ)	2024-09-11	8	300.00	1300.00	2024-09-11 20:00:32.83919
86	Tomato (टमाटर)	2024-09-11	10	90.00	900.00	2024-09-11 18:55:04.525364
47	Onion (प्याज)	2024-09-10	107	105.00	11235.00	2024-09-10 00:37:09.51269
89	Lapsi Mada (लप्सी माडा)	2024-09-11	1	100.00	100.00	2024-09-11 18:55:04.527673
50	Ginger (अदुवा)	2024-09-10	15	280.00	4200.00	2024-09-10 00:37:09.514633
90	Rayo Saag (रायो साग)	2024-09-11	70	140.00	9800.00	2024-09-11 18:55:04.528381
92	Peas (केराउ)	2024-09-11	1	120.00	120.00	2024-09-11 18:55:04.529636
93	Grapes (अंगुर)	2024-09-11	1	600.00	600.00	2024-09-11 18:55:04.530504
56	Golo Chyau (गोलो च्याउ)	2024-09-10	15	450.00	6750.00	2024-09-10 00:37:09.518794
38	Beans (सिमी)	2024-09-10	16	200.00	3200.00	2024-09-10 00:37:09.519361
78	Cucumber (काक्रो)	2024-09-11	89	35.00	4060.00	2024-09-11 20:00:32.840134
85	Kagati (कागती)	2024-09-11	15	240.00	3700.00	2024-09-11 20:00:32.841478
25	Cauliflower (काउली)	2024-09-10	190	100.00	18000.00	2024-09-10 00:59:01.813424
43	Cabbage (बन्दगोभी)	2024-09-10	25	50.00	1250.00	2024-09-10 00:59:01.814945
26	Radish (मूला)	2024-09-10	70	45.00	3150.00	2024-09-10 00:59:01.81579
27	Carrot (गाजर)	2024-09-10	73	110.00	8030.00	2024-09-10 00:59:01.816506
28	Cucumber (काक्रो)	2024-09-10	75	90.00	7200.00	2024-09-10 00:59:01.817182
29	Garlic (लसुन)	2024-09-10	12	130.00	6845.00	2024-09-10 00:59:01.817874
30	Chili (खुर्सानी)	2024-09-10	8	140.00	1120.00	2024-09-10 00:59:01.818574
32	Kagati (कागती)	2024-09-10	12	260.00	3240.00	2024-09-10 00:59:01.819287
34	Small Tomato (सानो टमाटर)	2024-09-10	93	80.00	7440.00	2024-09-10 00:59:01.819942
35	Capsicum (क्याप्सिकम)	2024-09-10	40	105.00	5400.00	2024-09-10 00:59:01.820608
36	Salad Patta (सलाद पत्ता)	2024-09-10	3	600.00	1900.00	2024-09-10 00:59:01.821259
37	Chyau Patey (च्याउ पाते)	2024-09-10	80	250.00	20500.00	2024-09-10 00:59:01.821912
72	Hariyo Simi (हरियो सिमी)	2024-09-10	4	220.00	880.00	2024-09-10 00:59:01.822624
39	Rayo Saag (रायो साग)	2024-09-10	115	250.00	27350.00	2024-09-10 00:59:01.824533
75	Cauliflower (काउली)	2024-09-11	15	140.00	2100.00	2024-09-11 18:55:04.517225
80	Onion (प्याज)	2024-09-11	52	120.00	6240.00	2024-09-11 18:55:04.520461
83	Ginger (अदुवा)5	2024-09-11	5	280.00	1400.00	2024-09-11 18:55:04.522621
97	Broccoli (ब्रोकोली)	2024-09-11	10	100.00	1000.00	2024-09-11 19:36:32.49783
125	Radish (मूला)	2024-09-26	1123	40.00	45445.00	2024-09-26 03:42:07.533268
131	Cauliflower (काउली)	2024-09-26	60	65.00	52716.00	2024-09-26 03:42:07.509518
88	Beans (सिमी)	2024-09-11	3	100.00	1170.00	2024-09-11 19:36:32.505478
124	Green Chillies (हरियो खुर्सानी)	2024-09-26	1	160.00	272682.00	2024-09-26 02:42:59.929999
123	Dhaniya (धनियाँ)	2024-09-26	51	400.00	7700.00	2024-09-26 02:42:59.933604
128	Kagati (कागती)	2024-09-26	6	280.00	12990.00	2024-09-26 02:42:59.936715
122	Carrot (गाजर)	2024-09-26	31	70.00	66279.00	2024-09-26 03:42:07.50454
91	Apple (स्याउ)	2024-09-11	214	292.00	62496.00	2024-09-11 19:36:32.520484
74	Potato (आलु)	2024-09-11	505	65.00	70027.50	2024-09-11 19:36:32.523145
94	Cabbage (बन्दगोभी)	2024-09-11	81	61.00	5076.00	2024-09-11 20:00:32.836558
87	Capsicum (क्याप्सिकम)	2024-09-11	14	120.00	1830.00	2024-09-11 20:00:32.837934
77	Carrot (गाजर)	2024-09-11	84	100.00	11350.00	2024-09-11 20:00:32.838684
81	Green Chillies (हरियो खुर्सानी)	2024-09-11	5	100.00	520.00	2024-09-11 20:00:32.842256
114	Mushroom (च्याउ)	2024-09-11	20	170.00	3400.00	2024-09-11 20:00:32.842952
115	Ginger (अदुवा)	2024-09-11	10	240.00	2400.00	2024-09-11 20:00:32.844273
76	Radish (मूला)	2024-09-11	51	35.00	1795.00	2024-09-11 20:00:32.844971
79	Garlic (लसुन)	2024-09-11	20	150.00	3510.00	2024-09-11 20:00:32.845348
118	Banana (केरा)	2024-09-11	194	295.00	57230.00	2024-09-11 20:00:32.846
119	Bitter Gourd (करेला)	2024-09-11	73	50.00	3650.00	2024-09-11 20:00:32.846402
126	Green Onion (हरियो प्याज)	2024-09-26	0	120.00	22056.00	2024-09-26 01:51:34.302291
130	palak saag	2024-09-26	38	65.00	2470.00	2024-09-26 01:51:34.306157
129	Rayo Saag (रायो साग)	2024-09-26	438	40.00	18220.00	2024-09-26 03:42:07.525063
127	Cucumber (काक्रो)	2024-09-26	20	60.00	98850.00	2024-09-26 03:42:07.507556
120	Cabbage (बन्दगोभी)	2024-09-26	0	40.00	25760.00	2024-09-26 03:42:07.494627
121	Capsicum (क्याप्सिकम)	2024-09-26	10	100.00	28450.00	2024-09-26 03:42:07.500905
139	bodhi	2024-09-26	40	60.00	2400.00	2024-09-26 01:51:34.316327
140	Bitter Gourd (करेला)	2024-09-26	131	40.00	5240.00	2024-09-26 02:05:24.541074
141	parwal	2024-09-26	70	60.00	4200.00	2024-09-26 02:05:24.542257
142	भिण्डी	2024-09-26	301	60.00	18060.00	2024-09-26 02:05:24.543962
145	mango	2024-09-26	0	110.00	337315.00	2024-09-26 02:05:24.549705
146	lichi	2024-09-26	744	300.00	223200.00	2024-09-26 02:05:24.5533
147	ghiraula	2024-09-26	110	40.00	4400.00	2024-09-26 02:05:24.555486
148	munta	2024-09-26	76	50.00	3800.00	2024-09-26 02:05:24.558376
153	Dhaniya (धनियाँ) and lasun	2024-09-26	2	350.00	700.00	2024-09-26 02:23:56.688234
158	Ginger (अदुवा) garlic	2024-09-26	4	280.00	1120.00	2024-09-26 02:23:56.69916
159	Small Tomato (सानो टमाटर)	2024-09-26	0	90.00	2115.00	2024-09-26 02:23:56.700401
164	Celery (अजमोदा)	2024-09-26	1	100.00	100.00	2024-09-26 02:23:56.706478
167	brinjal (bhanta)	2024-09-26	0	80.00	200.00	2024-09-26 02:23:56.709442
137	Mushroom (च्याउ)	2024-09-26	157	500.00	26400.00	2024-09-26 02:25:59.04241
170	Pumpkin (फर्सी)	2024-09-26	4	70.00	280.00	2024-09-26 02:25:59.045212
176	hariyo lasun	2024-09-26	1	200.00	200.00	2024-09-26 02:34:47.709014
157	Onion (प्याज)	2024-09-26	45	120.00	5400.00	2024-09-26 02:34:47.709525
181	Garlic (लसुन)	2024-09-26	5	260.00	1300.00	2024-09-26 02:34:47.718291
154	Lapsi Mada (लप्सी माडा)	2024-09-26	3	200.00	700.00	2024-09-26 02:34:47.72693
190	Grapes (अंगुर)	2024-09-26	1	600.00	600.00	2024-09-26 02:34:47.732593
166	pakucha saag	2024-09-26	15	210.00	5460.00	2024-09-26 02:37:54.759903
192	saag	2024-09-26	1	150.00	150.00	2024-09-26 02:37:54.777252
197	lasun	2024-09-26	0	200.00	200.00	2024-09-26 02:42:59.926177
143	Tomato (टमाटर)	2024-09-26	514	90.00	31896.00	2024-09-26 02:42:59.939747
132	Beans (सिमी)	2024-09-26	206	60.00	14054.00	2024-09-26 03:42:07.514348
221	Chili (खुर्सानी)	2024-09-26	0	90.00	3555.00	2024-09-26 03:42:07.517674
222	spring garlic	2024-09-26	0	120.00	14484.00	2024-09-26 03:42:07.520584
223	Lapsi  (लप्सी )	2024-09-26	59	70.00	4130.00	2024-09-26 03:42:07.523114
225	palak sag	2024-09-26	0	65.00	18037.50	2024-09-26 03:42:07.527008
136	Apple (स्याउ)	2024-09-26	1170	270.00	315930.00	2024-09-26 03:42:07.528754
135	Potato (आलु)	2024-09-26	4079	48.00	398483.20	2024-09-26 03:42:07.529735
228	kino	2024-09-26	2465	100.00	246500.00	2024-09-26 03:42:07.530903
138	Banana (केरा)	2024-09-26	1010	85.00	90445.00	2024-09-26 03:42:07.532044
231	Chyau Patey (च्याउ पाते)	2024-09-26	135	160.00	21600.00	2024-09-26 03:42:07.5344
232	lemon	2024-09-26	40	150.00	6000.00	2024-09-26 03:42:07.535541
133	Ginger (अदुवा)	2024-09-26	209	170.00	36080.00	2024-09-26 03:42:07.536495
134	Broccoli (ब्रोकोली)	2024-09-26	394	100.00	40600.00	2024-09-26 03:42:07.537357
235	cardinier leavs	2024-09-26	0	140.00	7126.00	2024-09-26 03:51:27.06273
236	Cabbage (बन्दगोभी)	2024-09-27	0	40.00	15660.00	2024-09-27 04:35:16.888278
237	Carrot (गाजर)	2024-09-27	0	70.00	22295.00	2024-09-27 04:35:16.89094
238	Radish (मूला)	2024-09-27	0	40.00	18860.00	2024-09-27 04:35:16.900513
239	cacumber	2024-09-27	0	60.00	67230.00	2024-09-27 04:35:16.901526
240	lemon	2024-09-27	0	150.00	6570.00	2024-09-27 04:35:16.904419
241	Capsicum (क्याप्सिकम)	2024-09-27	0	100.00	13770.00	2024-09-27 04:35:16.905245
242	Ginger (अदुवा)	2024-09-27	0	170.00	23256.00	2024-09-27 04:35:16.906927
243	spring ginger	2024-09-27	0	120.00	12516.00	2024-09-27 04:35:16.908808
244	Green Chillies (हरियो खुर्सानी)	2024-09-27	0	90.00	4365.00	2024-09-27 04:35:16.910245
245	Beans (सिमी)	2024-09-27	0	60.00	3330.00	2024-09-27 04:35:16.911746
246	lapsi	2024-09-27	0	70.00	4305.00	2024-09-27 04:35:16.913091
247	Ramtoria (रामतोरिया)	2024-09-27	457	60.00	27420.00	2024-09-27 04:35:16.914419
248	Potato (आलु)	2024-09-27	0	48.00	198734.40	2024-09-27 04:35:16.915723
249	Apple (स्याउ)	2024-09-27	0	270.00	300942.00	2024-09-27 04:35:16.91687
250	parwal	2024-09-27	320	60.00	19200.00	2024-09-27 04:35:16.918221
251	iskus	2024-09-27	300	35.00	10500.00	2024-09-27 04:35:16.919958
252	bodi	2024-09-27	124	60.00	7440.00	2024-09-27 04:35:16.921197
253	Cauliflower (काउली)	2024-09-27	64	65.00	4160.00	2024-09-27 04:35:16.922382
254	karela	2024-09-27	0	60.00	5430.00	2024-09-27 04:35:16.923938
255	cardinier	2024-09-27	0	140.00	2884.00	2024-09-27 04:35:16.926228
256	Rayo Saag (रायो साग)	2024-09-27	42	40.00	1680.00	2024-09-27 05:13:31.161375
257	naspati	2024-09-27	206	75.00	15450.00	2024-09-27 05:13:31.175969
258	Pumpkin (फर्सी)	2024-09-27	200	35.00	7000.00	2024-09-27 05:13:31.183603
259	Pudina (पुदिना)	2024-09-27	0	1000.00	500.00	2024-09-27 05:13:31.185142
260	Tomato (टमाटर)	2024-09-27	0	60.00	15450.00	2024-09-27 05:13:31.186904
261	Banana (केरा)	2024-09-27	0	140.00	167468.00	2024-09-27 05:13:31.187873
262	ghiraula	2024-09-27	107	40.00	4280.00	2024-09-27 05:13:31.190345
263	Chyau Patey (च्याउ पाते)	2024-09-27	173	160.00	27680.00	2024-09-27 05:13:31.19109
\.


--
-- Data for Name: ledger_entries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ledger_entries (id, date, account, billnumber, debit, credit, balance_amount, created_at) FROM stdin;
1	2024-09-09	Khowpa	102	200.00	200.00	400.00	2024-09-09 20:13:58.042321
2	2024-09-09	Ram Krishna Tajale	102	500.00	1025.00	1525.00	2024-09-09 20:27:14.923558
3	2080-10-29	khagi suppliers	186	1015326.00	1015326.00	0.00	2024-09-26 03:54:32.587809
\.


--
-- Name: balance_sheet_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.balance_sheet_id_seq', 1, false);


--
-- Name: bill_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bill_details_id_seq', 21, true);


--
-- Name: dailyvegetablesales_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dailyvegetablesales_id_seq', 263, true);


--
-- Name: ledger_entries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ledger_entries_id_seq', 3, true);


--
-- Name: balance_sheet balance_sheet_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.balance_sheet
    ADD CONSTRAINT balance_sheet_pkey PRIMARY KEY (id);


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

