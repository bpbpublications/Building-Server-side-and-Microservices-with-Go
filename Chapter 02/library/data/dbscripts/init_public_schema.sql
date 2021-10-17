-- Initialization of public schema relates to Library 0.x

-- enum_user_role
INSERT INTO enum_user_role
VALUES
    (1, 'member'),
    (2, 'librarian');

-- enum_book_status
INSERT INTO enum_book_status
VALUES
    (1, 'available'),
    (2, 'borrowed');

-- library_user
INSERT INTO library_user(username, user_password, full_name, user_role)
VALUES
    ('joe', crypt('joe', gen_salt('bf')), 'Average Joe', 1),
    ('smith', crypt('smith', gen_salt('bf')), 'John Smith', 2);

