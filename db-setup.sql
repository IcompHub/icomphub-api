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
    profile_picture_id TEXT,
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
    image_id TEXT,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_technologies
BEFORE UPDATE ON technologies
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
    thumbnail_id TEXT,
    data JSONB NOT NULL,
    class_group_id INTEGER NOT NULL REFERENCES class_groups(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_projects
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE project_images (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    image_id TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (image_id)
);

CREATE TRIGGER set_timestamp_project_images
BEFORE UPDATE ON project_images
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

CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    nickname TEXT,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,

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

CREATE TABLE member_roles (
    member_id INTEGER NOT NULL REFERENCES members(id) ON DELETE CASCADE,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,

    assigned_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (member_id, role_id)
);

CREATE TABLE project_technologies (
    id SERIAL PRIMARY KEY,
    status status_enum NOT NULL DEFAULT 'waiting_approval',
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    technology_id INTEGER NOT NULL REFERENCES technologies(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    UNIQUE (project_id, technology_id)
);

CREATE TRIGGER set_timestamp_project_technologies
BEFORE UPDATE ON project_technologies
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

-- Admin User
INSERT INTO public.users (slug, nickname, full_name, personal_email, "password", status, system_role)
VALUES
('nelson', 'Nelson', 'Nelson Pereira de Carvalho Neto', 'nelson.dev@test.com', '$2a$10$mJnlRbmI3Z2t2EcBSJP4AuCw3N9ZdduIpK8FRvWyP3RMntrzqsOXC', 'active', 'admin')

-- Additional Admin User
INSERT INTO public.users (slug, nickname, full_name, personal_email, "password", status, system_role)
VALUES
('admin2', 'AdminTwo', 'Admin Two', 'admin2@test.com', '$2a$10$mJnlRbmI3Z2t2EcBSJP4AuCw3N9ZdduIpK8FRvWyP3RMntrzqsOXC', 'active', 'admin');

-- Normal Users
INSERT INTO public.users (slug, nickname, full_name, personal_email, "password", status, system_role)
VALUES
('jdoe', 'JohnD', 'John Doe', 'john.doe@test.com', '$2a$10$mJnlRbmI3Z2t2EcBSJP4AuCw3N9ZdduIpK8FRvWyP3RMntrzqsOXC', 'active', 'user'),
('asmith', 'AliceS', 'Alice Smith', 'alice.smith@test.com', '$2a$10$mJnlRbmI3Z2t2EcBSJP4AuCw3N9ZdduIpK8FRvWyP3RMntrzqsOXC', 'active', 'user');

-- Technologies
INSERT INTO technologies (slug, name, status)
VALUES
('react', 'React', 'active'),
('nestjs', 'NestJS', 'active'),
('postgresql', 'PostgreSQL', 'active'),
('docker', 'Docker', 'active'),
('tailwind', 'Tailwind CSS', 'active');

-- Class Groups
INSERT INTO class_groups (slug, name, status)
VALUES
('group-a', 'Class Group A', 'active'),
('group-b', 'Class Group B', 'active');

-- Projects
INSERT INTO projects (slug, name, status, data, class_group_id)
VALUES
('proj-analytics', 'Analytics Platform', 'active', '{}', 1),
('proj-dashboard', 'Dashboard Builder', 'active', '{}', 2);

-- Get project IDs
-- Assuming default SERIALs, project IDs should be 1 and 2
-- Get user IDs to assign members (assuming users inserted in order above):
-- nelson (1), admin2 (2), jdoe (3), asmith (4)

-- Members for Project 1
INSERT INTO members (nickname, status, project_id, user_id)
VALUES
('nelson_proj1', 'active', 1, 1),
('jdoe_proj1', 'active', 1, 3);

-- Members for Project 2
INSERT INTO members (nickname, status, project_id, user_id)
VALUES
('admin2_proj2', 'active', 2, 2),
('asmith_proj2', 'active', 2, 4);

-- Link technologies to projects (2 per project)
-- Technologies assumed to be IDs 1 to 5
INSERT INTO project_technologies (project_id, technology_id, status)
VALUES
(1, 1, 'active'), -- React
(1, 2, 'active'), -- NestJS
(2, 3, 'active'), -- PostgreSQL
(2, 4, 'active'); -- Docker

INSERT INTO roles (slug, name, status)
VALUES
('project-manager', 'Project Manager', 'active'),
('developer', 'Developer', 'active'),
('designer', 'Designer', 'active');

-- Member 1 (nelson_proj1): Project Manager
INSERT INTO member_roles (member_id, role_id) VALUES (1, 1);

-- Member 2 (jdoe_proj1): Developer
INSERT INTO member_roles (member_id, role_id) VALUES (2, 2);

-- Member 3 (admin2_proj2): Project Manager + Developer
INSERT INTO member_roles (member_id, role_id) VALUES 
(3, 1),
(3, 2);

-- Member 4 (asmith_proj2): Designer
INSERT INTO member_roles (member_id, role_id) VALUES (4, 3);