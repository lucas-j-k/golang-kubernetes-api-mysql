-- +goose Up

-- +goose StatementBegin
CREATE TABLE note (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  label VARCHAR(255) NOT NULL,
  body TEXT NOT NULL,
  row_inserted TIMESTAMP NOT NULL,
  row_last_updated TIMESTAMP NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS note;
-- +goose StatementEnd
