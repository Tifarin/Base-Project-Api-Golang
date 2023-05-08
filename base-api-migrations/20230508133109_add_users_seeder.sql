-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO users (username, names, email, password_hash, id_roles, created_at)
VALUES --passwordnya (hash123)
('johndoe', 'John Doe', 'johndoe@example.com', '673D190B758967621DA243F06C350CE68BE4276174DC886560239FEA923D4A5A', 1, NOW()),
('superadmin', 'Super Admin', 'admin@example.com', '673D190B758967621DA243F06C350CE68BE4276174DC886560239FEA923D4A5A', 2, NOW()),
('bobsmith', 'Bob Smith', 'bobsmith@example.com', '673D190B758967621DA243F06C350CE68BE4276174DC886560239FEA923D4A5A', 1, NOW());
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
