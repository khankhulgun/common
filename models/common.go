package models

import (
	"gorm.io/gorm"
	"time"
)

type AboutUs struct {
	ID          string    `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(255)"`
	Description string    `gorm:"column:description;type:text"`
	Image       string    `gorm:"column:image;type:text"`
	ListOrder   int16     `gorm:"column:list_order;type:int2"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:timestamp"`
}
type Banner struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       *string `gorm:"column:title" json:"title"`
	Description *string `gorm:"column:description" json:"description"`
	Images      *string `gorm:"column:images" json:"images"`
	Link        *string `gorm:"column:link" json:"link"`
}

func (b *Banner) TableName() string {
	return "public.banner"
}

type AboutWeb struct {
	ID              int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title           string         `gorm:"column:title" json:"title"`
	AnalyticsKey    *string        `gorm:"column:analytics_key" json:"analytics_key"`
	Description     *string        `gorm:"column:description" json:"description"`
	MateKey         *string        `gorm:"column:mate_key" json:"mate_key"`
	OgImage         *string        `gorm:"column:og_image" json:"og_image"`
	FaviconImage    *string        `gorm:"column:favicon_image" json:"favicon_image"`
	LogoWideLight   *string        `gorm:"column:logo_wide_light" json:"logo_wide_light"`
	LogoWideDark    *string        `gorm:"column:logo_wide_dark" json:"logo_wide_dark"`
	LogoSquireDark  *string        `gorm:"column:logo_squire_dark" json:"logo_squire_dark"`
	LogoSquireLight *string        `gorm:"column:logo_squire_light" json:"logo_squire_light"`
	CreatedAt       *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (a *AboutWeb) TableName() string {
	return "public.about_web"
}

type ContactUs struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Email     string         `gorm:"column:email" json:"email"`
	Phone     int            `gorm:"column:phone" json:"phone"`
	Address   string         `gorm:"column:address" json:"address"`
	Comment   *string        `gorm:"column:comment" json:"comment"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (c *ContactUs) TableName() string {
	return "public.contact_us"
}

type Cooperation struct {
	PartnerTypeID *int           `gorm:"column:partner_type_id" json:"partner_type_id"`
	Title         *string        `gorm:"column:title" json:"title"`
	Logos         *string        `gorm:"column:logos" json:"logos"`
	Links         *string        `gorm:"column:links" json:"links"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at"`
	ID            int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (c *Cooperation) TableName() string {
	return "public.cooperation"
}

type LutPartnerType struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PartnerType *string `gorm:"column:partner_type" json:"partner_type"`
}

func (l *LutPartnerType) TableName() string {
	return "public.lut_partner_type"
}

type Faq struct {
	ID             string         `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Queation       *string        `gorm:"column:queation" json:"queation"`
	Answer         *string        `gorm:"column:answer" json:"answer"`
	Images         *string        `gorm:"column:images" json:"images"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt      *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt      *time.Time     `gorm:"column:created_at" json:"created_at"`
	QuestionTypeID *int           `gorm:"column:question_type_id" json:"question_type_id"`
}

func (f *Faq) TableName() string {
	return "public.faq"
}

type Feedback struct {
	UserID          string         `gorm:"column:user_id;type:uuid" json:"id"`
	FeedbackTypeID  *int           `gorm:"column:feedback_type_id" json:"feedback_type_id"`
	Feedback        *string        `gorm:"column:feedback" json:"feedback"`
	HariultAwahEseh *string        `gorm:"column:hariult_awah_eseh" json:"hariult_awah_eseh"`
	HariuTailbar    *string        `gorm:"column:hariu_tailbar" json:"hariu_tailbar"`
	Image           *string        `gorm:"column:image" json:"image"`
	CreatedAt       *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID              int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (f *Feedback) TableName() string {
	return "public.feedback"
}

type LutFeedbackType struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FeedbackType *string `gorm:"column:feedback_type" json:"feedback_type"`
	Icon         *string `gorm:"column:icon" json:"icon"`
}

func (l *LutFeedbackType) TableName() string {
	return "public.lut_feedback_type"
}

type Help struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleID     *int    `gorm:"column:role_id" json:"role_id"`
	Title      *string `gorm:"column:title" json:"title"`
	Link       *string `gorm:"column:link" json:"link"`
	File       *string `gorm:"column:file" json:"file"`
	HelpTypeID *int    `gorm:"column:help_type_id" json:"help_type_id"`
	Urls       *string `gorm:"column:urls" json:"urls"`
	Order      *int    `gorm:"column:order" json:"order"`
}

func (h *Help) TableName() string {
	return "public.help"
}

type LutHelpType struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	HelpType  *string `gorm:"column:help_type" json:"help_type"`
	OrderType *int    `gorm:"column:order_type" json:"order_type"`
}

func (l *LutHelpType) TableName() string {
	return "public.lut_help_type"
}

type News struct {
	ID          string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsTypeID  int            `gorm:"column:news_type_id" json:"news_type_id"`
	Title       string         `gorm:"column:title" json:"title"`
	Description string         `gorm:"column:description" json:"description"`
	Photo       *string        `gorm:"column:photo" json:"photo"`
	IsSlide     *int           `gorm:"column:is_slide" json:"is_slide"`
	IsPublish   *int           `gorm:"column:is_publish" json:"is_publish"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	IsSpecial   *int           `gorm:"column:is_special" json:"is_special"`
}

func (n *News) TableName() string {
	return "public.news"
}

type LutNewsType struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsType *string `gorm:"column:news_type" json:"news_type"`
}

func (l *LutNewsType) TableName() string {
	return "public.lut_news_type"
}

type Privacy struct {
	UserID    string         `gorm:"column:user_id;type:uuid" json:"id"`
	Privacy   *string        `gorm:"column:privacy" json:"privacy"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID        string         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (p *Privacy) TableName() string {
	return "public.privacy"
}

type RulesRegulations struct {
	ID                     int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RulesRegulationsTypeID *int           `gorm:"column:rules_regulations_type_id" json:"rules_regulations_type_id"`
	FormTitle              *string        `gorm:"column:form_title" json:"form_title"`
	FormFile               *string        `gorm:"column:form_file" json:"form_file"`
	FormLink               *string        `gorm:"column:form_link" json:"form_link"`
	ProcedureTitle         *string        `gorm:"column:procedure_title" json:"procedure_title"`
	ProcedureFile          *string        `gorm:"column:procedure_file" json:"procedure_file"`
	ProcedureLink          *string        `gorm:"column:procedure_link" json:"procedure_link"`
	CreatedAt              *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FileCheck              *string        `gorm:"column:file_check" json:"file_check"`
	FileLink               *string        `gorm:"column:file_link" json:"file_link"`
	FileMateral            *string        `gorm:"column:file_materal" json:"file_materal"`
}

func (r *RulesRegulations) TableName() string {
	return "public.rules_regulations"
}

type TermsOfService struct {
	ID                  int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID              string         `gorm:"column:user_id;type:uuid" json:"id"`
	TermsOfService      *string        `gorm:"column:terms_of_service" json:"terms_of_service"`
	CreatedAt           *time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	TermOfServiceTypeID *int           `gorm:"column:term_of_service_type_id" json:"term_of_service_type_id"`
}

func (t *TermsOfService) TableName() string {
	return "public.terms_of_service"
}

type LutServiceType struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ServiceType *string `gorm:"column:service_type" json:"service_type"`
}

func (l *LutServiceType) TableName() string {
	return "public.lut_service_type"
}

type LutTermService struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TermService *string `gorm:"column:term_service" json:"term_service"`
}

func (l *LutTermService) TableName() string {
	return "public.lut_term_service"
}

type LutFaqType struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FaqType *string `gorm:"column:faq_type" json:"faq_type"`
}

func (l *LutFaqType) TableName() string {
	return "public.lut_faq_type"
}
