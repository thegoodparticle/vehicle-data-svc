CREATE TABLE "owners" (
    "owner_id" text,
	"first_name" varchar(256) NOT NULL,
	"last_name" varchar(256) NOT NULL,
	"mobile_number" varchar(10) NOT NULL,
	"aadhar_number" varchar(12) NOT NULL,
	"address" varchar(1024) NOT NULL,
	"created_at" timestamp with time zone NOT NULL ,
	PRIMARY KEY ("owner_id")
) 

CREATE TABLE "vehicles" (
    "registration_number" text,
	"vehicle_model" varchar(256) NOT NULL,
	"company" varchar(256) NOT NULL,
	"engine_number" varchar(256) NOT NULL,
	"chassis_number" varchar(256) NOT NULL,
	"owner_id" varchar(256) NOT NULL,
	"registration_issue_date" timestamp with time zone NOT NULL,
	"registration_expiry_date" timestamp with time zone NOT NULL ,
	PRIMARY KEY ("registration_number")
) 