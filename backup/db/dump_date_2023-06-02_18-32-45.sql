--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: root
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO root;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: root
--

COMMENT ON SCHEMA public IS '';


--
-- Name: get_all_orders(); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.get_all_orders() RETURNS TABLE(order_uid integer, track_number character varying, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN QUERY SELECT o.order_uid, o.track_number, o.date_created FROM orders o;
END;
$$;


ALTER FUNCTION public.get_all_orders() OWNER TO root;

--
-- Name: get_order(integer); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.get_order(order_id integer) RETURNS TABLE(delivery_id integer, delivery_name character varying, delivery_phone character varying, delivery_zip character varying, delivery_city character varying, delivery_address character varying, delivery_region character varying, delivery_email character varying, payment_id integer, payment_transaction character varying, payment_request_id character varying, payment_currency character varying, payment_provider character varying, payment_amount integer, payment_dt integer, payment_bank character varying, payment_delivery_cost integer, payment_goods_total integer, payment_custom_fee integer, order_uid integer, order_track_number character varying, order_entry character varying, order_delivery_id integer, order_payment_id integer, order_locale character varying, order_internal_signature character varying, order_customer_id character varying, order_delivery_service character varying, order_shardkey character varying, order_sm_id integer, order_date_created timestamp without time zone, order_oof_shard character varying, item_id integer, item_chrt_id integer, item_track_number character varying, item_price numeric, item_rid character varying, item_name character varying, item_sale integer, item_size character varying, item_total_price integer, item_nm_id integer, item_brand character varying, item_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        d.delivery_id,
        d.name AS delivery_name,
        d.phone AS delivery_phone,
        d.zip AS delivery_zip,
        d.city AS delivery_city,
        d.address AS delivery_address,
        d.region AS delivery_region,
        d.email AS delivery_email,
        p.payment_id,
        p.transaction AS payment_transaction,
        p.request_id AS payment_request_id,
        p.currency AS payment_currency,
        p.provider AS payment_provider,
        p.amount AS payment_amount,
        p.payment_dt,
        p.bank AS payment_bank,
        p.delivery_cost AS payment_delivery_cost,
        p.goods_total AS payment_goods_total,
        p.custom_fee AS payment_custom_fee,
        o.order_uid,
        o.track_number AS order_track_number,
        o.entry AS order_entry,
        o.delivery_id AS order_delivery_id,
        o.payment_id AS order_payment_id,
        o.locale AS order_locale,
        o.internal_signature AS order_internal_signature,
        o.customer_id AS order_customer_id,
        o.delivery_service AS order_delivery_service,
        o.shardkey AS order_shardkey,
        o.sm_id AS order_sm_id,
        o.date_created AS order_date_created,
        o.oof_shard AS order_oof_shard,
        i.item_id,
        i.chrt_id AS item_chrt_id,
        i.track_number AS item_track_number,
        i.price AS item_price,
        i.rid AS item_rid,
        i.name AS item_name,
        i.sale AS item_sale,
        i.size AS item_size,
        i.total_price AS item_total_price,
        i.nm_id AS item_nm_id,
        i.brand AS item_brand,
        i.status AS item_status
    FROM orders o
    INNER JOIN delivery d ON o.delivery_id = d.delivery_id
    INNER JOIN payment p ON o.payment_id = p.payment_id
    INNER JOIN items i ON o.order_uid = i.order_uid
    WHERE o.order_uid = order_id;
END;
$$;


ALTER FUNCTION public.get_order(order_id integer) OWNER TO root;

--
-- Name: insert_delivery(text, text, text, text, text, text, text); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.insert_delivery(name text, phone text, zip text, city text, address text, region text, email text) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    delivery_id INTEGER;
BEGIN
    INSERT INTO delivery (name, phone, zip, city, address, region, email)
    VALUES (name, phone, zip, city, address, region, email)
    RETURNING delivery.delivery_id INTO delivery_id;
    
    RETURN delivery_id;
END;
$$;


ALTER FUNCTION public.insert_delivery(name text, phone text, zip text, city text, address text, region text, email text) OWNER TO root;

--
-- Name: insert_item(integer, integer, text, numeric, text, text, integer, text, integer, integer, text, integer); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.insert_item(order_uid integer, chrt_id integer, track_number text, price numeric, rid text, name text, sale integer, size text, total_price integer, nm_id integer, brand text, status integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    item_id INTEGER;
BEGIN
    INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
RETURNING items.item_id INTO item_id;
    
    RETURN item_id;
END;
$$;


ALTER FUNCTION public.insert_item(order_uid integer, chrt_id integer, track_number text, price numeric, rid text, name text, sale integer, size text, total_price integer, nm_id integer, brand text, status integer) OWNER TO root;

--
-- Name: insert_order(text, text, integer, integer, text, text, integer, text, text, integer, timestamp without time zone, text); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.insert_order(track_number text, entry text, delivery_id integer, payment_id integer, locale text, internal_signature text, customer_id integer, delivery_service text, shardkey text, sm_id integer, date_created timestamp without time zone, oof_shard text) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    order_uid integer;
BEGIN
    INSERT INTO orders (track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
    VALUES (track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
    RETURNING orders.order_uid INTO order_uid;
    
    RETURN order_uid;
END;
$$;


ALTER FUNCTION public.insert_order(track_number text, entry text, delivery_id integer, payment_id integer, locale text, internal_signature text, customer_id integer, delivery_service text, shardkey text, sm_id integer, date_created timestamp without time zone, oof_shard text) OWNER TO root;

--
-- Name: insert_order(text, text, integer, integer, text, text, text, text, text, integer, timestamp without time zone, text); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.insert_order(track_number text, entry text, delivery_id integer, payment_id integer, locale text, internal_signature text, customer_id text, delivery_service text, shardkey text, sm_id integer, date_created timestamp without time zone, oof_shard text) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    order_uid integer;
BEGIN
    INSERT INTO orders (track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
    VALUES (track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
    RETURNING orders.order_uid INTO order_uid;
    
    RETURN order_uid;
END;
$$;


ALTER FUNCTION public.insert_order(track_number text, entry text, delivery_id integer, payment_id integer, locale text, internal_signature text, customer_id text, delivery_service text, shardkey text, sm_id integer, date_created timestamp without time zone, oof_shard text) OWNER TO root;

--
-- Name: insert_payment(text, text, text, text, numeric, integer, text, numeric, numeric, numeric); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.insert_payment(transaction text, request_id text, currency text, provider text, amount numeric, payment_dt integer, bank text, delivery_cost numeric, goods_total numeric, custom_fee numeric) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    payment_id INTEGER;
BEGIN
    INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
    VALUES (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
    RETURNING payment.payment_id INTO payment_id;
    
    RETURN payment_id;
END;
$$;


ALTER FUNCTION public.insert_payment(transaction text, request_id text, currency text, provider text, amount numeric, payment_dt integer, bank text, delivery_cost numeric, goods_total numeric, custom_fee numeric) OWNER TO root;

--
-- Name: insert_q(); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.insert_q() RETURNS integer
    LANGUAGE plpgsql
    AS $$DECLARE
	delivery_id INTEGER;
	payment_id INTEGER;
	order_id INTEGER;
	item_id INTEGER;
BEGIN
	START TRANSACTION; -- Start the transaction explicitly

	delivery_id := insert_delivery('John Doe', '123456789', '12345', 'New York', '123 Main St', 'Region', 'john@example.com');
	payment_id := insert_payment('123456', '789', 'USD', 'Provider', 100.0, 32412321, 'Bank', 10.0, 90.0, 5.0);
	order_id := insert_order('123ABC', 'Entry', delivery_id, payment_id, 'en', 'signature', 123, 'Delivery Service', 'shardkey', 456, '2023-06-01', 'oof');
	item_id := insert_item(order_id, 123, 'qwe', 45, 'eeee', 'rr', 30, 'qweea', 12, 23, 'nk', 200);

	COMMIT; -- Commit the transaction
	RETURN item_id;
EXCEPTION
	WHEN OTHERS THEN -- Catch any exceptions
		ROLLBACK; -- Rollback the transaction
		RAISE; -- Rethrow the exception
END;$$;


ALTER FUNCTION public.insert_q() OWNER TO root;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: delivery; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.delivery (
    delivery_id integer NOT NULL,
    name character varying(255),
    phone character varying(255),
    zip character varying(255),
    city character varying(255),
    address character varying(255),
    region character varying(255),
    email character varying(255)
);


ALTER TABLE public.delivery OWNER TO root;

--
-- Name: delivery_delivery_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.delivery ALTER COLUMN delivery_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.delivery_delivery_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: items; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.items (
    item_id integer NOT NULL,
    order_uid integer,
    chrt_id integer,
    track_number character varying(255),
    price numeric,
    rid character varying(255),
    name character varying(255),
    sale integer,
    size character varying(255),
    total_price integer,
    nm_id integer,
    brand character varying(255),
    status integer
);


ALTER TABLE public.items OWNER TO root;

--
-- Name: items_item_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.items ALTER COLUMN item_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.items_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: orders; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.orders (
    order_uid integer NOT NULL,
    track_number character varying(255),
    entry character varying(255),
    delivery_id integer,
    payment_id integer,
    locale character varying(255),
    internal_signature character varying(255),
    customer_id character varying(255),
    delivery_service character varying(255),
    shardkey character varying(255),
    sm_id integer,
    date_created timestamp without time zone,
    oof_shard character varying(255)
);


ALTER TABLE public.orders OWNER TO root;

--
-- Name: orders_order_uid_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.orders ALTER COLUMN order_uid ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_order_uid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: payment; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.payment (
    payment_id integer NOT NULL,
    transaction character varying(255),
    request_id character varying(255),
    currency character varying(255),
    provider character varying(255),
    amount integer,
    payment_dt integer,
    bank character varying(255),
    delivery_cost integer,
    goods_total integer,
    custom_fee integer
);


ALTER TABLE public.payment OWNER TO root;

--
-- Name: payment_payment_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.payment ALTER COLUMN payment_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.payment_payment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: delivery; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.delivery (delivery_id, name, phone, zip, city, address, region, email) FROM stdin;
11	qw	\N	\N	\N	\N	\N	\N
12	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
13	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
14	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
15	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
16	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
17	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
18	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
19	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
27	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
28	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
30	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
31	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
32	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
33	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
34	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
35	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
36	John Doe	123456789	12345	New York	123 Main St	Region	john@example.com
39	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
40	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
41	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
47	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
48	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
52	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
53	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
57	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
58	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
59	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
60	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
61	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
62	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
63	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
64	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.items (item_id, order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) FROM stdin;
1	1	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
2	2	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
3	3	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
4	4	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
5	5	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
6	6	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
8	6	123	qwe	45	eeee	rr	30	qweea	12	23	nk	200
9	13	123	qwe	45	eeee	rr	30	qweea	12	23	nk	200
10	14	123	qwe	45	eeee	rr	30	qweea	12	23	nk	200
11	15	123	qwe	45	eeee	rr	30	qweea	12	23	nk	200
13	17	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
14	18	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
15	19	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
16	25	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
17	26	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
18	30	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
19	31	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
20	32	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
21	33	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
22	34	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
23	35	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
24	36	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
26	38	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
25	38	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
27	39	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
28	39	23	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	31	0	317	2389212	Vivienne Sabo	202
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) FROM stdin;
1	WBILMTESTTRACK	WBIL	13	4	en		test	meest	9	99	2021-11-26 06:22:19	1
2	jopa	WBIL	14	5	en		test	meest	9	99	2021-11-26 06:22:19	1
3	jopa	WBIL	15	6	en		test	meest	9	99	2021-11-26 06:22:19	1
4	jopa	BIILLLO	16	7	en		test	meest	9	99	2021-11-26 06:22:19	1
5	jopa	BIILLLO	17	8	en		test	meest	9	99	2021-11-26 06:22:19	1
6	jopa	BIILLLO	18	9	en		test	meest	9	99	2021-11-26 06:22:19	1
9	123ABC	Entry	19	11	en	signature	123	Delivery Service	shardkey	456	2023-06-01 00:00:00	oof
12	123ABC	Entry	33	14	en	signature	123	Delivery Service	shardkey	456	2023-06-01 00:00:00	oof
13	123ABC	Entry	34	15	en	signature	123	Delivery Service	shardkey	456	2023-06-01 00:00:00	oof
14	123ABC	Entry	35	16	en	signature	123	Delivery Service	shardkey	456	2023-06-01 00:00:00	oof
15	123ABC	Entry	36	17	en	signature	123	Delivery Service	shardkey	456	2023-06-01 00:00:00	oof
17	jopa	BIILLLO	39	19	en		test	meest	9	99	2021-11-26 06:22:19	1
18	jopa	BIILLLO	40	20	en		test	meest	9	99	2021-11-26 06:22:19	1
19	jopa	BIILLLO	41	21	en		test	meest	9	99	2021-11-26 06:22:19	1
25	jopa	BIILLLO	47	27	en		test	meest	9	99	2021-11-26 06:22:19	1
26	jopa	BIILLLO	48	28	en		test	meest	9	99	2021-11-26 06:22:19	1
30	jopa	BIILLLO	52	32	en		test	meest	9	99	2021-11-26 06:22:19	1
31	jopa	BIILLLO	53	33	en		test	meest	9	99	2021-11-26 06:22:19	1
32	jopa	BIILLLO	57	37	en		test	meest	9	99	2021-11-26 06:22:19	1
33	jopa	BIILLLO	58	38	en		test	meest	9	99	2021-11-26 06:22:19	1
34	jopa	BIILLLO	59	39	en		test	meest	9	99	2021-11-26 06:22:19	1
35	jopa	BIILLLO	60	40	en		test	meest	9	99	2021-11-26 06:22:19	1
36	jopa	BIILLLO	61	41	en		test	meest	9	99	2021-11-26 06:22:19	1
37	jopa	BIILLLO	62	42	en		test	meest	9	99	2021-11-26 06:22:19	1
38	jopa	BIILLLO	63	43	en		test	meest	9	99	2021-11-26 06:22:19	1
39	jopa	BIILLLO	64	44	en		test	meest	9	99	2021-11-26 06:22:19	1
\.


--
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.payment (payment_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) FROM stdin;
3	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
4	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
5	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
6	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
7	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
8	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
9	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
10	123456	789	USD	Provider	100	32412321	Bank	10	90	5
11	123456	789	USD	Provider	100	32412321	Bank	10	90	5
13	123456	789	USD	Provider	100	333	Bank	10	90	5
14	123456	789	USD	Provider	100	333	Bank	10	90	5
15	123456	789	USD	Provider	100	32412321	Bank	10	90	5
16	123456	789	USD	Provider	100	32412321	Bank	10	90	5
17	123456	789	USD	Provider	100	32412321	Bank	10	90	5
19	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
20	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
21	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
27	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
28	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
32	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
33	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
37	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
38	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
39	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
40	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
41	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
42	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
43	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
44	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
\.


--
-- Name: delivery_delivery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.delivery_delivery_id_seq', 64, true);


--
-- Name: items_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.items_item_id_seq', 28, true);


--
-- Name: orders_order_uid_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.orders_order_uid_seq', 39, true);


--
-- Name: payment_payment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.payment_payment_id_seq', 44, true);


--
-- Name: delivery delivery_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_pkey PRIMARY KEY (delivery_id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (item_id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid) INCLUDE (order_uid);


--
-- Name: payment payment_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (payment_id);


--
-- Name: orders orders_delivery_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_delivery_id_fkey FOREIGN KEY (delivery_id) REFERENCES public.delivery(delivery_id);


--
-- Name: orders orders_payment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_payment_id_fkey FOREIGN KEY (payment_id) REFERENCES public.payment(payment_id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: root
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


--
-- PostgreSQL database dump complete
--

