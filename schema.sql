
-- workspaces
CREATE TABLE "workspaces" (
  "id"          bigserial PRIMARY KEY,
  "name"        text NOT NULL,
  "created_at"  timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "workspaces" IS 'organizes the data with workspaces';
COMMENT ON COLUMN "workspaces"."id" IS 'workspace id';
COMMENT ON COLUMN "workspaces"."name" IS 'workspace name';

-- sources
CREATE TABLE "sources" (
  "id"            bigserial PRIMARY KEY,
  "workspace_id"  bigint NOT NULL REFERENCES "workspaces" ("id") ON DELETE CASCADE,
  "generator"     text NOT NULL,
  "source_info"   jsonb,
  "generated_at"  timestamp with time zone,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "sources" IS 'source info of imported data';
COMMENT ON COLUMN "sources"."id" IS 'source id';
COMMENT ON COLUMN "sources"."generator" IS 'data generator name';
COMMENT ON COLUMN "sources"."source_info" IS 'extra information about the source';
COMMENT ON COLUMN "sources"."generated_at" IS 'source original timestamp';

-- hosts
CREATE TABLE "hosts" (
  "id"            bigserial PRIMARY KEY,
  "source_id"     bigint NOT NULL REFERENCES "sources" ("id") ON DELETE CASCADE,
  "address"       inet NOT NULL,
  "state"         text NOT NULL,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON COLUMN "hosts"."id" IS 'host id';
COMMENT ON COLUMN "hosts"."address" IS 'host address';

-- services
CREATE TABLE "services" (
  "id"            bigserial PRIMARY KEY,
  "host_id"       bigint NOT NULL REFERENCES "hosts" ("id") ON DELETE CASCADE,
  "protocol"      text NOT NULL,
  "port"          integer NOT NULL,
  "state"         text NOT NULL,
  "service"       text,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "services" IS 'running service of a host';
COMMENT ON COLUMN "services"."id" IS 'service id';
COMMENT ON COLUMN "services"."protocol" IS 'service protocol (tcp, etc.)';
COMMENT ON COLUMN "services"."port" IS 'service port';
COMMENT ON COLUMN "services"."state" IS 'service status';

-- notes
CREATE TABLE "notes" (
  "id"            bigserial PRIMARY KEY,
  "host_id"       bigint NOT NULL REFERENCES "hosts" ("id") ON DELETE CASCADE,
  "title"         text NOT NULL,
  "content"       text NOT NULL,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "notes" IS 'extra information of a host';
COMMENT ON COLUMN "notes"."title" IS 'information name';
COMMENT ON COLUMN "notes"."content" IS 'note content';

-- issues
CREATE TYPE issue_level AS ENUM (
  'critical',
  'high',
  'medium',
  'low',
  'info'
);
CREATE TABLE "issues" (
  "id"            bigserial PRIMARY KEY,
  "host_id"       bigint NOT NULL REFERENCES "hosts" ("id") ON DELETE CASCADE,
  "level"         issue_level NOT NULL,
  "title"         text NOT NULL,
  "content"       text NOT NULL,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "issues" IS 'issues of a host';
COMMENT ON COLUMN "issues"."level" IS 'issue level';
COMMENT ON COLUMN "issues"."title" IS 'issue name';
COMMENT ON COLUMN "issues"."content" IS 'issue content';

-- stats
CREATE VIEW "issues_summaries" AS
  WITH
    "counts" AS (
      SELECT
        -- count of issues by level
        COUNT(*) FILTER (WHERE "i"."level" = 'info') AS "info",
        COUNT(*) FILTER (WHERE "i"."level" = 'low') AS "low",
        COUNT(*) FILTER (WHERE "i"."level" = 'medium') AS "medium",
        COUNT(*) FILTER (WHERE "i"."level" = 'high') AS "high",
        COUNT(*) FILTER (WHERE "i"."level" = 'critical') AS "critical"
      FROM "issues" AS "i"
    ),
    "titles" AS (
      -- top 10 issues titles
      SELECT
        "title", "level",
        COUNT("id") AS "number"
      FROM "issues"
      GROUP BY "title", "level"
      ORDER BY "level", "number" DESC
      LIMIT 10
    )
  SELECT
    -- an id to make go-pg happy
    1 as "id",

    -- see "counts"
    ROW_TO_JSON("counts") AS "stats",

    -- see "titles"
    JSON_AGG(ROW_TO_JSON("titles")) AS "titles"
  FROM "counts", "titles"
  GROUP BY "counts".*;
COMMENT ON VIEW "issues_summaries" IS 'summary of issues with stats';
