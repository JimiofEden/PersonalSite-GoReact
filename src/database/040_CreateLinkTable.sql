-- Table: dbo.Link

-- DROP TABLE dbo.Link;

CREATE TABLE IF NOT EXISTS dbo.Link
(
    Id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    LinkName character(50) COLLATE pg_catalog."default" NOT NULL,
    Url character(255) COLLATE pg_catalog."default" NOT NULL,
    Deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedDateTime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CreatedBy CHARACTER(127) NOT NULL DEFAULT CURRENT_USER,
    LastModifiedDateTime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    LastModifiedBy CHARACTER(127) NOT NULL DEFAULT CURRENT_USER,
    CONSTRAINT "PK_Link" PRIMARY KEY (Id),
    UNIQUE (LinkName)
)

TABLESPACE pg_default;

ALTER TABLE dbo.Link
    OWNER to jimi;

COMMENT ON TABLE dbo.Link
    IS 'Creates a table to contain Link';