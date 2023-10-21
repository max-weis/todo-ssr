CREATE TABLE todos
(
    text TEXT PRIMARY KEY CHECK (LENGTH(text) <= 256) NOT NULL,
    done BOOLEAN                                      NOT NULL
);
