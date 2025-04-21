CREATE TYPE TaskStatus AS ENUM ('created', 'in_progress', 'canceled', 'completed', 'failed');

CREATE TABLE task_list
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER      NOT NULL,
    title       VARCHAR(128) NOT NULL,
    description VARCHAR(255),
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE TABLE task
(
    id           SERIAL PRIMARY KEY,
    task_list_id INTEGER      NOT NULL REFERENCES task_list (id) ON DELETE CASCADE,
    title        VARCHAR(128) NOT NULL,
    description  VARCHAR(255),
    user_id      INTEGER      NOT NULL,
    status       TaskStatus   NOT NULL DEFAULT 'created',
    created_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE TABLE task_share
(
    task_id INTEGER NOT NULL REFERENCES task (id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL,
    PRIMARY KEY (task_id, user_id)
);

CREATE TABLE task_list_share
(
    task_list_id INTEGER NOT NULL REFERENCES task_list (id) ON DELETE CASCADE,
    user_id      INTEGER NOT NULL,
    PRIMARY KEY (task_list_id, user_id)
);
