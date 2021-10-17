-- Initialization of public schema relates to Library 0.x

-- enum_book_status
INSERT INTO enum_book_status
VALUES
    (1, 'available'),
    (2, 'borrowed');
