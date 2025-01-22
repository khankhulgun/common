package migrations

import (
	"github.com/lambda-platform/lambda/DB"
	"log"
)

func MigrateLookupTables() {
	LutPartnerType := `
	CREATE TABLE IF NOT EXISTS "public"."lut_partner_type" (
		   "id" SERIAL PRIMARY KEY,
    "partner_type" VARCHAR(255) COLLATE "pg_catalog"."default" NULL
	);
	`
	if err := DB.DB.Exec(LutPartnerType).Error; err != nil {
		log.Fatalf("Failed to create lut_partner_type table: %v", err)
	}

	LutFeedbackType := `
	CREATE TABLE IF NOT EXISTS "public"."lut_feedback_type" (
		   "id" SERIAL PRIMARY KEY,
    "feedback_type" VARCHAR(255) COLLATE "pg_catalog"."default" NULL,
    "Icon" TEXT COLLATE "pg_catalog"."default" NULL
	);
	`
	if err := DB.DB.Exec(LutFeedbackType).Error; err != nil {
		log.Fatalf("Failed to create lut_feedback_type table: %v", err)
	}

	LutHelpType := `
	CREATE TABLE IF NOT EXISTS "public"."lut_help_type" (
		    "id" SERIAL PRIMARY KEY,
    "help_type" VARCHAR(255) COLLATE "pg_catalog"."default" NULL,
    "OrderType" INTEGER NULL
	);
	`
	if err := DB.DB.Exec(LutHelpType).Error; err != nil {
		log.Fatalf("Failed to create lut_help_type table: %v", err)
	}

	LutNewsType := `
	CREATE TABLE IF NOT EXISTS "public"."lut_news_type" (
		   "id" SERIAL PRIMARY KEY,
    "lut_news_type" VARCHAR(255) COLLATE "pg_catalog"."default" NULL
	);
	`
	if err := DB.DB.Exec(LutNewsType).Error; err != nil {
		log.Fatalf("Failed to create lut_news_type table: %v", err)
	}

	LutServiceType := `
	CREATE TABLE IF NOT EXISTS "public"."lut_service_type" (
		   "id" SERIAL PRIMARY KEY,
    "service_type" VARCHAR(255) COLLATE "pg_catalog"."default" NULL
	);
	`
	if err := DB.DB.Exec(LutServiceType).Error; err != nil {
		log.Fatalf("Failed to create lut_service_type table: %v", err)
	}

	LutTermService := `
	CREATE TABLE IF NOT EXISTS "public"."lut_term_service" (
		   "id" SERIAL PRIMARY KEY,
    "term_service" VARCHAR(255) COLLATE "pg_catalog"."default" NULL
	);
	`
	if err := DB.DB.Exec(LutTermService).Error; err != nil {
		log.Fatalf("Failed to create lut_term_service table: %v", err)
	}

	LutFaqType := `
	CREATE TABLE IF NOT EXISTS "public"."lut_faq_type" (
		   "id" SERIAL PRIMARY KEY,
    "faq_type" VARCHAR(255) COLLATE "pg_catalog"."default" NULL
	);
	`
	if err := DB.DB.Exec(LutFaqType).Error; err != nil {
		log.Fatalf("Failed to create lut_faq_type table: %v", err)
	}
}
