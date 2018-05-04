
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
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "sources" IS 'source info of imported data';
COMMENT ON COLUMN "sources"."id" IS 'source id';
COMMENT ON COLUMN "sources"."generator" IS 'data generator name';

-- hosts
CREATE TABLE "hosts" (
  "id"            bigserial PRIMARY KEY,
  "source_id"     bigint NOT NULL REFERENCES "sources" ("id") ON DELETE CASCADE,
  "address"       inet NOT NULL,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON COLUMN "hosts"."id" IS 'host id';
COMMENT ON COLUMN "hosts"."address" IS 'host address';

-- services
CREATE TABLE "services" (
  "id"            bigserial PRIMARY KEY,
  "host_id"       bigint NOT NULL REFERENCES "hosts" ("id") ON DELETE CASCADE,
  "port"          integer NOT NULL,
  "service"       text,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "services" IS 'running service of a host';
COMMENT ON COLUMN "services"."id" IS 'service id';
COMMENT ON COLUMN "services"."port" IS 'service port';

-- infos
CREATE TABLE "infos" (
  "id"            bigserial PRIMARY KEY,
  "host_id"       bigint NOT NULL REFERENCES "hosts" ("id") ON DELETE CASCADE,
  "name"          text NOT NULL,
  "data"          jsonb NOT NULL,
  "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);
COMMENT ON TABLE "infos" IS 'extra information of a host';
COMMENT ON COLUMN "infos"."name" IS 'information name';
COMMENT ON COLUMN "infos"."data" IS 'schema-less data';
