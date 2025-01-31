DROP TABLE IF EXISTS `memos`;
DROP TABLE IF EXISTS `queries`;
DROP TABLE IF EXISTS `resources`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` TEXT NOT NULL PRIMARY KEY,
  `username` TEXT NOT NULL,
  `password` TEXT NOT NULL,
  `open_id` TEXT NOT NULL DEFAULT '',
  `created_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  `updated_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  UNIQUE(`username`, `open_id`)
);

CREATE TABLE `queries`  (
  `id` TEXT NOT NULL PRIMARY KEY,
  `user_id` TEXT NOT NULL,
  `title` TEXT NOT NULL,
  `querystring` TEXT NOT NULL,
  `created_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  `updated_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  `pinned_at` TEXT NOT NULL DEFAULT '',
  FOREIGN KEY(`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `memos`  (
  `id` TEXT NOT NULL PRIMARY KEY,
  `content` TEXT NOT NULL,
  `user_id` TEXT NOT NULL,
  `created_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  `updated_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  `deleted_at` TEXT NOT NULL DEFAULT '',
  FOREIGN KEY(`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `resources`  (
  `id` TEXT NOT NULL PRIMARY KEY,
  `user_id` TEXT NOT NULL,
  `filename` TEXT NOT NULL,
  `blob` BLOB NOT NULL,
  `type` TEXT NOT NULL,
  `size` INTEGER NOT NULL DEFAULT 0,
  `created_at` TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  FOREIGN KEY(`user_id`) REFERENCES `users`(`id`)
);


INSERT INTO `users`
  (`id`, `username`, `password`, `open_id`)
VALUES
  ('1', 'guest', '123456', 'guest_open_id'),
  ('2', 'mine', '123456', 'mine_open_id');

INSERT INTO `memos`
  (`id`, `content`, `user_id`)
VALUES
  ('1', '👋 Welcome to memos', '1'),
  ('2', '👋 Welcome to memos', '2');
