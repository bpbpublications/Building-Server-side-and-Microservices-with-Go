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
CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA pg_catalog;

SET search_path = public, pg_catalog;
SET default_tablespace = '';

-- enum_user_role
CREATE TABLE enum_user_role (
    code integer NOT NULL,
    user_role text NOT NULL,
    CONSTRAINT enum_user_status_pk PRIMARY KEY (code)
);

-- library_user
CREATE TABLE library_user (
    user_id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    username text NOT NULL UNIQUE,
    user_password text NOT NULL,
    full_name text NOT NULL,
    user_role integer DEFAULT 1,
    token uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    CONSTRAINT library_user_pk PRIMARY KEY (user_id),
    CONSTRAINT fk_library_user_user_role FOREIGN KEY (user_role)
        REFERENCES enum_user_role (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
