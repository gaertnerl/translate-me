package database

var (
	Q_NEXT_TRANSLATION                       = "SELECT * FROM translations WHERE id > ? AND translation_type = ? ORDER BY id ASC LIMIT 1;"
	Q_CREATE_TRANSLATION_TABLE_IF_NOT_EXISTS = "CREATE TABLE IF NOT EXISTS translations(" +
		"id INT AUTO_INCREMENT PRIMARY KEY," +
		"translation_type UNSIGNED TINYINT," +
		"sentence_a VARCHAR(255)," +
		"sentence_b VARCHAR(255));"

	Q_INSERT_TRANSLATION      = "INSERT INTO translations(translation_type, sentence_a, sentence_b) VALUES (?, ?, ?);"
	Q_GET_TRANSLATION_WITH_ID = "SELECT * FROM translations WHERE id = ?;"
	Q_CLEAR_TABLE             = "TRUNCATE ?;"
)
