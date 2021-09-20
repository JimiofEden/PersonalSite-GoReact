-- Table: dbo.Skill

-- DROP TABLE dbo.Skill;

CREATE TABLE IF NOT EXISTS dbo.Skill
(
    Id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    SkillName character(50) COLLATE pg_catalog."default" NOT NULL,
    SkillType character(50) COLLATE pg_catalog."default" NOT NULL,
    Url character(255) COLLATE pg_catalog."default",
    Comment character(255) COLLATE pg_catalog."default",
    CONSTRAINT "PK_Skill" PRIMARY KEY (Id)
)

TABLESPACE pg_default;

ALTER TABLE dbo.Skill
    OWNER to jimi;

COMMENT ON TABLE dbo.Skill
    IS 'Creates a table to contain Skills';