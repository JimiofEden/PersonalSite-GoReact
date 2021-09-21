-- =============================================
-- Author:      Adam Hollock
-- Create date: 2021-09-21
-- =============================================
CREATE TEMP TABLE temp_SkillType
(
    SkillTypeId INTEGER NOT NULL,
    SkillTypeName CHARACTER(50) NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedBy CHARACTER(255) NOT NULL DEFAULT CURRENT_USER,
    CreatedDateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastModifiedBy CHARACTER(255) NOT NULL DEFAULT CURRENT_USER,
    LastModifiedDateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) 
ON COMMIT DROP;

BEGIN;

INSERT INTO temp_SkillType
(SkillTypeId, SkillTypeName)
VALUES
(1, 'Backend'),
(2, 'Frontend'),
(3, 'Database'),
(4, 'Server'),
(5, 'Misc.');



-- -- Postgres does not yet support merging
-- MERGE INTO dbo.SkillType TARGET
-- USING temp_SkillType SOURCE
--     ON (TARGET.SkillTypeId = SOURCE.SkillTypeId)
-- WHEN MATCHED   THEN              --- <--- use this if want to check on another column
-- UPDATE SET 
--     TARGET.SkillTypeName = SOURCE.SkillTypeName,
    
--     TARGET.LastModifiedBy = SOURCE.LastModifiedBy,
--     TARGET.LastModifiedDateTime = SOURCE.LastModifiedDateTime
-- WHEN NOT MATCHED BY TARGET THEN 
--     INSERT (SkillTypeId, SkillTypeName

--         , Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime
--         )
--     VALUES (SOURCE.SkillTypeId, SOURCE.SkillTypeName

--         , SOURCE.Deleted, SOURCE.CreatedBy, SOURCE.CreatedDateTime, SOURCE.LastModifiedBy, SOURCE.LastModifiedDateTime
--         )
-- ;--WHEN NOT MATCHED BY SOURCE THEN
-- --DELETE;

INSERT INTO dbo.SkillType 
(SkillTypeId, SkillTypeName

, Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime)
select 
SkillTypeId, SkillTypeName

, Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime
from temp_SkillType
ON CONFLICT (SkillTypeId)
DO
    UPDATE SET
    SkillTypeName = EXCLUDED.SkillTypeName,

    Deleted = EXCLUDED.Deleted,
    LastModifiedBy = EXCLUDED.LastModifiedBy,
    LastModifiedDateTime = EXCLUDED.LastModifiedDateTime
;

COMMIT
--ROLLBACK