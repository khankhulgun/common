package seeds

import (
	"github.com/lambda-platform/lambda/DB"
	"log"
)

func SeedLookupTables() {
	// faq_type
	faq_type := []struct {
		faq_type string
	}{
		{"test 1"},
	}

	for _, style := range faq_type {
		query := `
		INSERT INTO "public"."lut_faq_type" ("faq_type")
		VALUES (?)
		
		`
		if err := DB.DB.Exec(query, style.faq_type).Error; err != nil {
			log.Printf("Failed to seed lut_line_style: %v", err)
		}
	}

	//news
	news_type := []struct {
		news_type string
	}{
		{"test 1"},
	}

	for _, style := range news_type {
		query := `
		INSERT INTO "public"."lut_news_type" ("news_type")
		VALUES (?)
		
		`
		if err := DB.DB.Exec(query, style.news_type).Error; err != nil {
			log.Printf("Failed to seed lut_news_type: %v", err)
		}
	}
}
