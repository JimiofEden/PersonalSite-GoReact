-- =============================================
-- Author:      Adam Hollock
-- Create date: 2021-09-21
-- =============================================
CREATE TABLE IF NOT EXISTS dbo.SkillType
(
    Id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    SkillTypeId INTEGER NOT NULL,
    SkillTypeName CHARACTER(50) NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedDateTime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CreatedBy CHARACTER(127) NOT NULL DEFAULT CURRENT_USER,
    LastModifiedDateTime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastModifiedBy CHARACTER(127) NOT NULL DEFAULT CURRENT_USER,
    PRIMARY KEY (Id),
    UNIQUE (SkillTypeId)
);

ALTER TABLE dbo.SkillType
    OWNER to jimi;


COMMENT ON TABLE dbo.SkillType
    IS 'Creates lookup table with different skilltype references';