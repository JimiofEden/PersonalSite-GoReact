-- =============================================
-- Author:      Adam Hollock
-- Create date: 2021-09-21
-- =============================================
CREATE TEMP TABLE temp_SkillType
(
    SkillTypeId INTEGER NOT NULL,
    SkillTypeName CHARACTER(50) NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE
) 
ON COMMIT DROP;

BEGIN TRANSACTION;

INSERT INTO temp_SkillType
(SkillTypeId, SkillTypeName)
VALUES
(1, 'Backend'),
(2, 'Frontend'),
(3, 'Database'),
(4, 'Server'),
(5, 'Misc.');

WITH auditConstants (userName, currentTime) as (
   values (CURRENT_USER, CURRENT_TIMESTAMP)
)
MERGE INTO dbo.SkillType as TARGET
USING temp_SkillType as SOURCE
    ON (TARGET.SkillTypeId = SOURCE.SkillTypeId)
WHEN MATCHED   THEN              --- <--- use this if want to check on another column
UPDATE SET 
    TARGET.SkillTypeName = SOURCE.SkillTypeName,
    
    TARGET.LastModifiedBy = userName,
    TARGET.LastModifiedDateTime = currentTime
WHEN NOT MATCHED BY TARGET THEN 
    INSERT (SkillTypeId, SkillTypeName

        , Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime
        )
    VALUES (SOURCE.SkillTypeId, SOURCE.SkillTypeName

        , SOURCE.Deleted, @user, @updatedtime, @user, @updatedtime
        )
;--WHEN NOT MATCHED BY SOURCE THEN
--DELETE;

COMMIT;