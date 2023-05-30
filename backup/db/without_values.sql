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
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.items (item_id, order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) FROM stdin;
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) FROM stdin;
\.


--
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.payment (payment_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) FROM stdin;
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
-- PostgreSQL database dump complete
--

