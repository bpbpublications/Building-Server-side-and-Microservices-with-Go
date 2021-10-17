-- Initial public schema relates to Library 0.x

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;

SET search_path = public, pg_catalog;
SET default_tablespace = '';

-- report
CREATE TABLE report (
    report_id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    book_id uuid NOT NULL,
    report_date timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT report_pk PRIMARY KEY (report_id)
);

CREATE INDEX report_report_date
ON report (report_date);
