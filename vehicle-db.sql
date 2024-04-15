CREATE TABLE "vehicles" (
    "registration_number" text,
    "vehicle_model" varchar(255) NOT NULL,
    "company" varchar(255) NOT NULL,
    "engine_number" varchar(255) NOT NULL,
    "chassis_number" varchar(255) NOT NULL,
    "owner_license" varchar(255) NOT NULL,
    "registration_date" timestamp with time zone NOT NULL ,
    PRIMARY KEY ("registration_number")
) 

CREATE TABLE "drivers" (
    "license_number" text,
    "driver_name" varchar(255) NOT NULL,
    "license_date" timestamp with time zone NOT NULL ,
    PRIMARY KEY ("license_number")
) 