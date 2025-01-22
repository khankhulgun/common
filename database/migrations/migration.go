package migrations

import (
	"github.com/khankhulgun/common/models"
	"github.com/lambda-platform/lambda/DB"
	"log"
)

func Migrate() {

	DB.DB.AutoMigrate(
		&models.AboutUs{},
		&models.Banner{},
		&models.AboutWeb{},
		&models.ContactUs{},
		&models.Cooperation{},
		&models.LutHelpType{},
		&models.Help{},
		&models.Feedback{},
		&models.LutFeedbackType{},
		&models.News{},
		&models.LutNewsType{},
		&models.Privacy{},
		&models.RulesRegulations{},
		&models.TermsOfService{},
		&models.LutServiceType{},
		&models.Faq{},
		&models.LutPartnerType{},
		&models.LutTermService{},
		&models.LutFaqType{},
	)

	MigrateLookupTables()

	createView := `
CREATE OR REPLACE VIEW public.v_cooperation AS
SELECT 
   cooperation.partner_type_id,
   cooperation.title,
   cooperation.logos,
   cooperation.links,
   cooperation.deleted_at,
   cooperation.updated_at,
   cooperation.created_at,
   cooperation.id,
   lut_partner_type.partner_type
FROM 
   public.cooperation
LEFT JOIN 
   public.lut_partner_type 
ON 
   cooperation.partner_type_id = public.lut_partner_type.id;
CREATE OR REPLACE VIEW public.v_faq AS
 SELECT faq.queation,
    faq.answer,
    faq.images,
    faq.deleted_at,
    faq.updated_at,
    faq.created_at,
    faq.id,
    faq.question_type_id,
    lut_faq_type.faq_type
   FROM faq
     LEFT JOIN lut_faq_type ON faq.question_type_id = lut_faq_type.id;

CREATE OR REPLACE VIEW public.v_feedback AS
 SELECT feedback.user_id,
    feedback.feedback_type_id,
    feedback.feedback,
    feedback.hariu_tailbar,
    feedback.image,
    feedback.created_at,
    feedback.updated_at,
    feedback.deleted_at,
    feedback.id,
    lut_feedback_type.feedback_type
   FROM feedback
     LEFT JOIN lut_feedback_type ON feedback.feedback_type_id = lut_feedback_type.id;


CREATE OR REPLACE VIEW public.v_help AS

 SELECT help.id,
    help.role_id,
    help.title,
    help.link,
    help.file,
    help.help_type_id,
    help.urls,
    help."order",
    roles.display_name,
    lut_help_type.help_type
   FROM help
     LEFT JOIN roles ON help.role_id = roles.id
     LEFT JOIN lut_help_type ON help.help_type_id = lut_help_type.id;


CREATE OR REPLACE VIEW public.v_news AS
 SELECT news.id,
    news.news_type_id,
    news.title,
    news.description,
    news.photo,
    news.is_slide,
    news.is_publish,
    news.created_at,
    news.updated_at,
    news.deleted_at,
    news.is_special,
    lut_news_type.news_type
   FROM news
     LEFT JOIN lut_news_type ON news.news_type_id = lut_news_type.id;
CREATE OR REPLACE VIEW public.v_term_service AS

 SELECT terms_of_service.id,
    terms_of_service.user_id,
    terms_of_service.terms_of_service,
    terms_of_service.created_at,
    terms_of_service.updated_at,
    terms_of_service.deleted_at,
    terms_of_service.term_of_service_type_id,
    lut_term_service.term_service
   FROM terms_of_service
     LEFT JOIN lut_term_service ON terms_of_service.term_of_service_type_id = lut_term_service.id;


	`
	err := DB.DB.Exec(createView).Error
	if err != nil {
		log.Fatalf("Failed to create view: %v", err)
	}

}
