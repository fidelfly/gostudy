CREATE TABLE jwt_auth_token
(
    id               INT AUTO_INCREMENT PRIMARY KEY,
    create_time      DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version          INT DEFAULT 0,
    token            VARCHAR(1000)                      NULL,
    refresh_time     DATETIME                           NULL,
    invalid_time     DATETIME                           NULL,
    deprecated_token VARCHAR(1000)                      NULL,
    client_ip        VARCHAR(100)                       NULL,
    user_agent       VARCHAR(1000)                      NULL,
    CONSTRAINT jwt_auth_token_id_uindex
    UNIQUE (id)
)
