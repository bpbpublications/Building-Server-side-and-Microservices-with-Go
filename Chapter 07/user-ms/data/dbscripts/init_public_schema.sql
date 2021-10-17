-- Initialization of public schema relates to Library 0.x

-- enum_user_role
INSERT INTO enum_user_role
VALUES
    (1, 'member'),
    (2, 'librarian');

-- library_user
INSERT INTO library_user(username, user_password, full_name, user_role)
VALUES
    ('joe', crypt('joe', gen_salt('bf')), 'Average Joe', 1),
    ('smith', crypt('smith', gen_salt('bf')), 'John Smith', 2);
