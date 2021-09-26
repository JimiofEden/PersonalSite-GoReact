-- =============================================
-- Author:      Adam Hollock
-- Create date: 2021-09-21
-- =============================================
CREATE TEMP TABLE temp_Skill
(
    SkillName character(50) NOT NULL,
    SkillTypeId integer NOT NULL,
    Url CHARACTER(255) NOT NULL,
    Comment CHARACTER(255) NOT NULL,
    Sequence integer NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedBy CHARACTER(255) NOT NULL DEFAULT CURRENT_USER,
    CreatedDateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastModifiedBy CHARACTER(255) NOT NULL DEFAULT CURRENT_USER,
    LastModifiedDateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) 
ON COMMIT DROP;

BEGIN;

WITH c (SkillTypeName) as (
   values ('Backend')
)
INSERT INTO temp_Skill
(SkillName, SkillTypeId, Url, Comment, Sequence)
select 
'Go', SkillTypeId, 'https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/serverapp', 'This site is being served by Go!', 1
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'C#', SkillTypeId, '', '.Net Legacy and Core', 2
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Node', SkillTypeId, '', '', 3
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Python', SkillTypeId, '', '', 4
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Others', SkillTypeId, 'https://github.com/JimiofEden?tab=repositories', 'Ruby, PHP, MATLAB', 5
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
;

WITH c (SkillTypeName) as (
   values ('Frontend')
)
INSERT INTO temp_Skill
(SkillName, SkillTypeId, Url, Comment, Sequence)
select 
'React', SkillTypeId, 'https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/clientapp', 'This site is being rendered with React!', 1
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Sass/SCSS, Less, CSS', SkillTypeId, '', '', 2
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Angular', SkillTypeId, '', '', 3
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Knockout', SkillTypeId, '', '', 4
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
;

WITH c (SkillTypeName) as (
   values ('Database')
)
INSERT INTO temp_Skill
(SkillName, SkillTypeId, Url, Comment, Sequence)
select 
'Postgres', SkillTypeId, 'https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/database', 'This data is being stored in Postgres!', 1
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Graphql', SkillTypeId, 'https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/clientapp/src/util/queries', 'This data is being queried by Graphql', 2
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'SQL Server', SkillTypeId, '', '', 3
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'MySQL', SkillTypeId, '', '', 4
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
;

WITH c (SkillTypeName) as (
   values ('Server')
)
INSERT INTO temp_Skill
(SkillName, SkillTypeId, Url, Comment, Sequence)
select 
'Docker', SkillTypeId, 'https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/server', 'This site''s host was built with a docker image!', 1
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'AWS', SkillTypeId, '', 'This site is being hosted by an AWS instance!', 2
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Azure', SkillTypeId, '', '', 3
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'IIS', SkillTypeId, '', '', 4
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Apache', SkillTypeId, '', '', 5
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
;

WITH c (SkillTypeName) as (
   values ('Misc.')
)
INSERT INTO temp_Skill
(SkillName, SkillTypeId, Url, Comment, Sequence)
select 
'Continuous Integration', SkillTypeId, '', 'Teamcity, Octopus Deploy', 1
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Version Control', SkillTypeId, 'https://github.com/JimiofEden/PersonalSite-GoReact', 'Git (also SVN)', 2
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Unit Testing', SkillTypeId, '', 'Jest, NUnit', 3
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
union
select 
'Project Management', SkillTypeId, '', '', 4
from dbo.SkillType st, c
where st.SkillTypeName = c.SkillTypeName
;


INSERT INTO dbo.Skill 
(SkillName, SkillTypeId, Url, Comment, Sequence

, Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime)
select 
SkillName, SkillTypeId, Url, Comment, Sequence

, Deleted, CreatedBy, CreatedDateTime, LastModifiedBy, LastModifiedDateTime
from temp_Skill
ON CONFLICT (SkillName, SkillTypeId)
DO
    UPDATE SET
    Url = EXCLUDED.Url,
    Comment = EXCLUDED.Comment,
    Sequence = EXCLUDED.Sequence,

    Deleted = EXCLUDED.Deleted,
    LastModifiedBy = EXCLUDED.LastModifiedBy,
    LastModifiedDateTime = EXCLUDED.LastModifiedDateTime
;

COMMIT
--ROLLBACK