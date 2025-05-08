CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TYPE status_enum AS ENUM ('active', 'inactive', 'waiting_approval');

CREATE TYPE system_role_enum AS ENUM ('admin', 'user', 'professor');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    nickname TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    personal_email TEXT NOT NULL UNIQUE CHECK (position('@' in personal_email) > 1),
    institutional_email TEXT UNIQUE CHECK (
        institutional_email IS NULL OR position('@' in institutional_email) > 1
    ),
    registration TEXT UNIQUE,
    password TEXT NOT NULL,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    system_role system_role_enum NOT NULL DEFAULT 'user',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TYPE entity_target_enum AS ENUM (
    'users',
    'class_groups',
    'projects',
    'roles',
    'technologies'
);

CREATE TYPE release_action_enum AS ENUM (
    'create',
    'update',
    'delete'
);

CREATE TYPE release_status_enum AS ENUM (
    'pending',
    'approved',
    'rejected',
    'cancelled'
);

CREATE TABLE releases (
    id SERIAL PRIMARY KEY,

    target_table entity_target_enum NOT NULL,
    target_id INTEGER NOT NULL,
    
    version INTEGER NOT NULL DEFAULT 1,

    action release_action_enum NOT NULL,
    data JSONB NOT NULL,

    status release_status_enum NOT NULL DEFAULT 'pending',

    requested_by INTEGER REFERENCES users(id),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (target_table, target_id, version)
);

CREATE TRIGGER set_timestamp_releases
BEFORE UPDATE ON releases
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TYPE release_event_type_enum AS ENUM (
    'submitted',
    'pending',
    'commented',
    'rejected',
    'approved',
    'cancelled',
    'reopened'
);

CREATE TABLE release_events (
    id SERIAL PRIMARY KEY,
    release_id INTEGER NOT NULL REFERENCES releases(id) ON DELETE CASCADE,

    "type" release_event_type_enum NOT NULL,
    actor_id INTEGER REFERENCES users(id),
    comment TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE technologies (
    id SERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_technologies
BEFORE UPDATE ON technologies
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_roles
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TYPE class_group_role_enum AS ENUM ('instructor', 'co_instructor', 'assistant', 'student');

CREATE TABLE class_groups (
    id SERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_class_groups
BEFORE UPDATE ON class_groups
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE class_group_users (
    id SERIAL PRIMARY KEY,
    class_group_id INTEGER NOT NULL REFERENCES class_groups(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    class_group_role class_group_role_enum NOT NULL DEFAULT 'student',
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_class_group_users
BEFORE UPDATE ON class_group_users
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    data JSONB NOT NULL,
    class_group_id INTEGER NOT NULL REFERENCES class_groups(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_projects
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    nickname TEXT,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    UNIQUE (project_id, nickname),
    UNIQUE (project_id, user_id),
    UNIQUE (project_id, user_id, nickname)
);

CREATE TRIGGER set_timestamp_members
BEFORE UPDATE ON members
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE projects_technologies (
    id SERIAL PRIMARY KEY,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    technology_id INTEGER NOT NULL REFERENCES technologies(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    UNIQUE (project_id, technology_id)
);

CREATE TRIGGER set_timestamp_projects_technologies
BEFORE UPDATE ON projects_technologies
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

INSERT INTO public.users (slug, nickname, full_name, personal_email, "password", status, system_role)
VALUES
('nelson', 'Nelson', 'Nelson Pereira de Carvalho Neto', 'nelson.dev@test.com', 'teste123', 'active', 'admin')