-- +goose Up
-- +goose StatementBegin
CREATE TABLE `logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `job_id` VARCHAR(255) NOT NULL,
  `job_name` VARCHAR(255) NOT NULL,
  `recipient` VARCHAR(255) NOT NULL,
  `content` TEXT NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `logs`;
-- +goose StatementEnd
