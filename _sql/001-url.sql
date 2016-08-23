CREATE TABLE url (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    reference VARCHAR(13) NOT NULL,
    url VARCHAR(2000) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY (reference)
) CHARACTER SET=utf8mb4 COMMENT='Stores URLs and their references.';
