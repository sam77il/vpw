--
-- PostgreSQL database dump
--

\restrict j2wqsBxtGXlmBjEEmoEubcduezygEdBbiAxzLblhUCcz0GoEvx6QYBak3cvcsNI

-- Dumped from database version 18.3
-- Dumped by pg_dump version 18.3

-- Started on 2026-05-06 00:21:22

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- TOC entry 223 (class 1259 OID 32919)
-- Name: cart_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cart_items (
    id integer NOT NULL,
    user_id uuid NOT NULL,
    product_id text NOT NULL,
    amount integer,
    metadata jsonb DEFAULT '{}'::jsonb NOT NULL
);


ALTER TABLE public.cart_items OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 32918)
-- Name: cart_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cart_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cart_items_id_seq OWNER TO postgres;

--
-- TOC entry 5045 (class 0 OID 0)
-- Dependencies: 222
-- Name: cart_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cart_items_id_seq OWNED BY public.cart_items.id;


--
-- TOC entry 221 (class 1259 OID 32877)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id text NOT NULL,
    label text CONSTRAINT categories_name_not_null NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 32866)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id text NOT NULL,
    label text CONSTRAINT products_title_not_null NOT NULL,
    description text,
    price numeric(10,2) NOT NULL,
    stock integer DEFAULT 0,
    category_id text,
    options jsonb DEFAULT '{}'::jsonb,
    old_price numeric(10,2)
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 32853)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    role text DEFAULT 'user'::text,
    created_at timestamp without time zone DEFAULT now(),
    gender text,
    first_name text,
    last_name text,
    street text,
    postal_code text,
    city text,
    country text,
    phone_number text,
    company boolean DEFAULT false,
    company_name text,
    company_ustidnr text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 4874 (class 2604 OID 32922)
-- Name: cart_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items ALTER COLUMN id SET DEFAULT nextval('public.cart_items_id_seq'::regclass);


--
-- TOC entry 5039 (class 0 OID 32919)
-- Dependencies: 223
-- Data for Name: cart_items; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 5037 (class 0 OID 32877)
-- Dependencies: 221
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories VALUES ('aluverbundplatten2mm', 'Aluverbundplatten 2mm');
INSERT INTO public.categories VALUES ('aluverbundplatten3mm', 'Aluverbundplatten 3mm');
INSERT INTO public.categories VALUES ('aluverbundplatten4mm', 'Aluverbundplatten 4mm');
INSERT INTO public.categories VALUES ('digitaldruck', 'Digitaldruck');


--
-- TOC entry 5036 (class 0 OID 32866)
-- Dependencies: 220
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products VALUES ('aluverbundplatte-2mm-3050x1500', 'Aluverbundplatte 2mm 3050x1500mm', 'Description lol', 119.74, 100, 'aluverbundplatten2mm', '{"product_type": "aluverbundplatte"}', 119.74);
INSERT INTO public.products VALUES ('aluverbundplatte-3mm-2050x1000', 'Aluverbundplatte 3mm 2050x1000mm', 'Description lol', 60.74, 100, 'aluverbundplatten3mm', '{"product_type": "aluverbundplatte"}', 60.74);
INSERT INTO public.products VALUES ('aluverbundplatte-3mm-2550x1250', 'Aluverbundplatte 3mm 2550x1250mm', 'Description lol', 94.22, 100, 'aluverbundplatten3mm', '{"product_type": "aluverbundplatte"}', 94.22);
INSERT INTO public.products VALUES ('aluverbundplatte-3mm-3050x1000', 'Aluverbundplatte 3mm 3050x1000mm', 'Description lol', 90.37, 100, 'aluverbundplatten3mm', '{"product_type": "aluverbundplatte"}', 90.37);
INSERT INTO public.products VALUES ('aluverbundplatte-3mm-3050x1500', 'Aluverbundplatte 3mm 3050x1500mm', 'Description lol', 135.56, 100, 'aluverbundplatten3mm', '{"product_type": "aluverbundplatte"}', 135.56);
INSERT INTO public.products VALUES ('aluverbundplatte-4mm-2050x1000', 'Aluverbundplatte 4mm 2050x1000mm', 'Description lol', 68.06, 100, 'aluverbundplatten4mm', '{"product_type": "aluverbundplatte"}', 68.06);
INSERT INTO public.products VALUES ('aluverbundplatte-4mm-3050x1000', 'Aluverbundplatte 4mm 3050x1000mm', 'Description lol', 101.26, 100, 'aluverbundplatten4mm', '{"product_type": "aluverbundplatte"}', 101.26);
INSERT INTO public.products VALUES ('aluverbundplatte-4mm-2550x1250', 'Aluverbundplatte 4mm 2550x1250mm', 'Description lol', 105.58, 100, 'aluverbundplatten4mm', '{"product_type": "aluverbundplatte"}', 105.58);
INSERT INTO public.products VALUES ('aluverbundplatte-4mm-3050x1500', 'Aluverbundplatte 4mm 3050x1500mm', 'Description lol', 151.89, 100, 'aluverbundplatten4mm', '{"product_type": "aluverbundplatte"}', 151.89);
INSERT INTO public.products VALUES ('digitaldruck', 'Digitaldruck', 'Description lol', 8.99, 100, 'digitaldruck', '{"product_type": "digitaldruck"}', 8.99);


--
-- TOC entry 5035 (class 0 OID 32853)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('2a30148b-f515-46ab-9b06-facdc1a65ae4', 'test@gmail.com', '$2a$10$MoBRijB.dDK3wanA3g8pQOPAzjcFxf2Uy4rwW5ixfrZaCxt1d6QDK', 'user', '2026-05-05 18:56:44.359072', 'male', 'Max', 'Mustermann', 'Musterstraße 15A', '55555', 'Musterstadt', 'DE', '0176176176', true, 'Sugarweb GmbH', 'DE123123');


--
-- TOC entry 5046 (class 0 OID 0)
-- Dependencies: 222
-- Name: cart_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cart_items_id_seq', 7, true);


--
-- TOC entry 4885 (class 2606 OID 32929)
-- Name: cart_items cart_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_pkey PRIMARY KEY (id);


--
-- TOC entry 4883 (class 2606 OID 32885)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 4881 (class 2606 OID 32876)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 4877 (class 2606 OID 32865)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4879 (class 2606 OID 32863)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4886 (class 2606 OID 32935)
-- Name: cart_items fk_cart_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT fk_cart_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- TOC entry 4887 (class 2606 OID 32930)
-- Name: cart_items fk_cart_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT fk_cart_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


-- Completed on 2026-05-06 00:21:22

--
-- PostgreSQL database dump complete
--

\unrestrict j2wqsBxtGXlmBjEEmoEubcduezygEdBbiAxzLblhUCcz0GoEvx6QYBak3cvcsNI

