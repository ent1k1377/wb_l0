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

CREATE FUNCTION public.get_all_orders() RETURNS TABLE(order_uid character varying, track_number character varying, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN QUERY SELECT o.order_uid, o.track_number, o.date_created FROM orders o;
END;
$$;


ALTER FUNCTION public.get_all_orders() OWNER TO root;

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
-- Name: items; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.items (
    item_id integer NOT NULL,
    order_uid character varying(255),
    chrt_id integer,
    track_number character varying(255),
    price integer,
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
-- Name: orders; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.orders (
    order_uid character varying(255) NOT NULL,
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
-- Data for Name: delivery; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.delivery (delivery_id, name, phone, zip, city, address, region, email) FROM stdin;
1	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
2	Jane Smith	+9722222222	9876543	Haifa	Ocean Avenue 10	Carmel	jane@example.com
4	Michael Brown	+9724444444	7654321	Netanya	Beach Road 20	Sela	michael@example.com
3	Anna Johnson	+9723333333	5432198	Jerusalem	Park Avenue 5	Givat Ram	anna@example.com
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.items (item_id, order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) FROM stdin;
1	b563feb7b2b84b6test	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
2	c895ghe7a4d12f3test	5678234	HAIFATESTTRACK	120	ab4219087a764ae0btest	Lipsticks	10	0	500	3298412	MAC	201
4	f1c2e3d4a5b6test	1234567	NETATESTTRACK	80	ab4219087a764ae0btest	Foundation	15	0	120	7890123	Oréal	201
3	e9b85fca2d61456test	7890123	JERUTESTTRACK	60	ab4219087a764ae0btest	Eyeliners	5	0	150	4567890	Maybelline	200
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) FROM stdin;
b563feb7b2b84b6test	WBILMTESTTRACK	WBIL	1	1	en		test	meest	9	99	2021-11-26 06:22:19	1
c895ghe7a4d12f3test	HAIFATESTTRACK	HAIF	2	2	en		test	dhl	6	88	2021-11-27 09:10:00	2
f1c2e3d4a5b6test	NETATESTTRACK	NETA	4	4	en		test	fedex	8	66	2021-11-27 17:00:00	4
e9b85fca2d61456test	JERUTESTTRACK	JERU	3	3	en		test	ups	3	77	2021-11-27 13:30:00	3
\.


--
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.payment (payment_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) FROM stdin;
1	b563feb7b2b84b6test		USD	wbpay	1817	1637907727	alpha	1500	317	0
2	c895ghe7a4d12f3test		EUR	paynet	2500	1637912000	beta	2000	500	0
4	f1c2e3d4a5b6test		USD	stripe	920	1637918000	delta	800	120	0
3	e9b85fca2d61456test		GBP	paypal	750	1637915000	gamma	600	150	0
\.


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
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);


--
-- Name: payment payment_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (payment_id);


--
-- Name: items items_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid);


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

