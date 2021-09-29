-- =============================================
-- Author:      Adam Hollock
-- Create date: 2021-09-21
-- =============================================
CREATE TEMP TABLE temp_StoredLink
(
    LinkName character(50) NOT NULL,
    Url CHARACTER(255) NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedBy CHARACTER(255) NOT NULL DEFAULT CURRENT_USER,
    CreatedDateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastModifiedBy CHARACTER(255) NOT NULL DEFAULT CURRENT_USER,
    LastModifiedDateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) 
ON COMMIT DROP;

BEGIN;

INSERT INTO temp_StoredLink
(LinkName, Url)
VALUES
('twitter', 'https://twitter.com/JimiofEden'),
('resume', './AH-Resume_0721-linkedin.pdf'),
('github', 'https://github.com/jimiofeden'),
('email', 'mailto:jimiofeden@gmail.com')
;

INSERT INTO dbo.StoredLink 
(LinkName, Url

, Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime)
select 
LinkName, Url

, Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime
from temp_StoredLink
ON CONFLICT (LinkName)
DO
    UPDATE SET
    Url = EXCLUDED.Url,

    Deleted = EXCLUDED.Deleted,
    LastModifiedBy = EXCLUDED.LastModifiedBy,
    LastModifiedDateTime = EXCLUDED.LastModifiedDateTime
;

COMMIT
--ROLLBACK