CREATE DATABASE nourybot CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE nourybot;

-- create a 'channel' table.
CREATE TABLE channel (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    twitchid VARCHAR(50) NOT NULL,
    added DATETIME NOT NULL
);

-- add an index on the created column.
CREATE INDEX idx_channel_created ON channel(added);

-- add dummy records
INSERT INTO channel (username, twitchid, added) VALUES (
    'whereismymnd',
    '31437432',
    UTC_TIMESTAMP()
);

INSERT INTO channel (username, twitchid, added) VALUES (
    'nourybot',
    '596581605',
    UTC_TIMESTAMP()
);

INSERT INTO channel (username, twitchid, added) VALUES (
    'xnoury',
    '197780373',
    UTC_TIMESTAMP()
);

-- Important: Make sure to swap 'pass' and 'username' with a password/user of your own choosing.
CREATE USER 'username'@'localhost';
GRANT SELECT, INSERT ON nourybot.* TO 'username'@'localhost';

ALTER USER 'username'@'localhost' IDENTIFIED BY 'pass';