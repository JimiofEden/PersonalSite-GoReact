-- Table: dbo.Skill

-- DROP TABLE dbo.Skill;

CREATE TABLE IF NOT EXISTS dbo.Skill
(
    Id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    SkillName character(50) COLLATE pg_catalog."default" NOT NULL,
    SkillTypeId integer NOT NULL,
    Url character(255) COLLATE pg_catalog."default",
    Comment character(255) COLLATE pg_catalog."default",
    Sequence integer NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedDateTime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CreatedBy CHARACTER(127) NOT NULL DEFAULT CURRENT_USER,
    LastModifiedDateTime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastModifiedBy CHARACTER(127) NOT NULL DEFAULT CURRENT_USER,
    CONSTRAINT PK_Skill PRIMARY KEY (Id),
    UNIQUE (SkillName, SkillTypeId),
    CONSTRAINT FK_SkillTypeId FOREIGN KEY (SkillTypeId)
        REFERENCES dbo.SkillType (SkillTypeId) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE dbo.Skill
    OWNER to jimi;

COMMENT ON TABLE dbo.Skill
    IS 'Creates a table to contain Skills';