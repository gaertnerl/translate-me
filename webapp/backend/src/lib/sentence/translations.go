package sentence

import (
	"database/sql"
	"strconv"

	"github.com/gaertnerl/translate-me.git/webserver/lib/database"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
)

func getTranslationWithId(db *sql.DB, translation_id TranslationId) Translation {
	var translation Translation
	rows, err := db.Query(database.Q_GET_TRANSLATION_WITH_ID, translation_id)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if !rows.Next() {
		id := strconv.Itoa(int(translation_id))
		panic("no translation with id " + id + " found")
	}

	err = rows.Scan(&translation.Translation_id, &translation.Translation_type, &translation.Sentence_1, &translation.Sentence_2)
	if err != nil {
		panic(err.Error())
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	return translation
}

func InsertTranslation(db *sql.DB, translate_type TranslationType, sentence_a string, sentence_b string) {
	_, err := db.Query(database.Q_CREATE_TRANSLATION_TABLE_IF_NOT_EXISTS)
	if err != nil {
		panic("cannot create translation db table, reason: \n" + err.Error())
	}
	_, err = db.Query(database.Q_INSERT_TRANSLATION, translate_type, sentence_a, sentence_b)
	if err != nil {
		panic("cannot insert translation, reason: \n" + err.Error())
	}
}

// sencond parameter is false if no next sentence exists
func GetNextTranslation(db *sql.DB, id TranslationId, translationType TranslationType) (Translation, bool) {

	var next Translation
	rows, err := db.Query(database.Q_NEXT_TRANSLATION, id, translationType)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if !rows.Next() {
		return next, false
	}

	err = rows.Scan(&next.Translation_id, &next.Translation_type, &next.Sentence_1, &next.Sentence_2)
	if err != nil {
		panic(err.Error())
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	return next, true
}

func EvaluateUserTranslation(db *sql.DB, service similarity.SimilarityService, ut UserTranslation) (UserTranslationEvaluation, error) {
	var ute UserTranslationEvaluation
	score, err := similarity.GetSimilarityScore(service, ut.UserTranslated, ut.Correct)
	if err != nil {
		return ute, nil
	}
	ute.Score = score
	return ute, nil
}
