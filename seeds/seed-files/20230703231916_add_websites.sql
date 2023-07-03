-- +goose Up
-- +goose StatementBegin
INSERT INTO website (id, domain, name, row_inserted, row_last_updated)
VALUES
    (1, 'examplesite.com', 'Example Site', NOW(), NULL),
    (2, 'testsite.net', 'Test Site', NOW(), NULL),
    (3, 'dummy-site.co.uk', 'Dummy Site', NOW(), NULL);
-- +goose StatementEnd

