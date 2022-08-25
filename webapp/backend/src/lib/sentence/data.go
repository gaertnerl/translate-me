package sentence

import "database/sql"

func PrepareTranslations(db *sql.DB) {
	InsertTranslation(db, EN_DE, "Tomorrow I have to get up early to catch my plane.", "Morgen muss ich früh aufstehen um meinen Flieger zu erwischen.")
	InsertTranslation(db, EN_DE, "Each sommer I go to visit my favourite Lake at least once.", "Jeden Sommer besuche ich meine Lieblingssee mindestens einmal.")
	InsertTranslation(db, EN_DE, "Most sundays I just spend in my garden.", "Die meissten Sonntage verbringe ich einfach in meinem Garten.")
	InsertTranslation(db, EN_DE, "I would like to help you, however I have my own problems right now.", "Ich würde dir gerne helfen, allerdings habe ich gerade meine eigenen Probleme.")
	InsertTranslation(db, EN_DE, "What an amazing song! Can you tell me the name?", "Was für ein toller Song! Kannst du mir den Namen sagen?")
	InsertTranslation(db, EN_DE, "I thank you from the bottom of my heart.", "Ich danke dir von ganzem Herzen.")
}
