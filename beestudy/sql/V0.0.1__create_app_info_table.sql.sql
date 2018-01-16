CREATE TABLE app_file
(
  id          INT AUTO_INCREMENT PRIMARY KEY,
  name        VARCHAR(200),
  size        INT,
  md5         VARCHAR(200),
  data        LONGBLOB,
  create_time DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  app_code    VARCHAR(100),
  app_desc    NVARCHAR(400),
  app_version VARCHAR(100),
  CONSTRAINT app_file_id_uindex
  UNIQUE (id)
);

CREATE TABLE app
(
  id     INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(100) NOT NULL,
  `desc` VARCHAR(400),
  create_time DATETIME DEFAULT current_timestamp NOT NULL,
  version VARCHAR(100),
  CONSTRAINT app_id_uindex
  UNIQUE (id),
  CONSTRAINT app_code_uindex
  UNIQUE (code)
);