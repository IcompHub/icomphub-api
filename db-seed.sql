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