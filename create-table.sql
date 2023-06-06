CREATE TABLE IF NOT EXISTS account (
    id SERIAL,
    username VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL
);

ALTER TABLE account ADD CONSTRAINT pk_account PRIMARY KEY (id);
CREATE UNIQUE INDEX idx_accountemail ON account (email);


CREATE TABLE IF NOT EXISTS post (
    id SERIAL,
    title VARCHAR(256) NOT NULL,
    body TEXT NOT NULL,
    userId INTEGER NOT NULL
);

ALTER TABLE post ADD CONSTRAINT pk_post PRIMARY KEY (id);
ALTER TABLE post ADD CONSTRAINT fk_bookuserid FOREIGN KEY (userid) REFERENCES account (id) ON DELETE CASCADE;
