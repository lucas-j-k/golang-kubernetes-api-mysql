-- +goose Up

-- +goose StatementBegin
CREATE TABLE website (
  id INT NOT NULL AUTO_INCREMENT,
  domain VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  row_inserted TIMESTAMP NOT NULL,
  row_last_updated TIMESTAMP NULL,
  PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS website;

-- +goose StatementEnd