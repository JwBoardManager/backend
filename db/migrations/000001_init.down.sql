-- Removendo índices
DROP INDEX IF EXISTS idx_meetings_date;
DROP INDEX IF EXISTS idx_sessions_room;
DROP INDEX IF EXISTS idx_assignments_user;
DROP INDEX IF EXISTS idx_cart_shifts_day;
DROP INDEX IF EXISTS idx_cart_assignments_user;
DROP INDEX IF EXISTS idx_cleaning_meeting;

-- Removendo tabelas
DROP TABLE IF EXISTS house_visits;
DROP TABLE IF EXISTS territories;
DROP TABLE IF EXISTS cart_assignments;
DROP TABLE IF EXISTS cart_shifts;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS cleaning_assignments;
DROP TABLE IF EXISTS assignments;
DROP TABLE IF EXISTS subsessions;
DROP TABLE IF EXISTS subsession_types;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS session_types;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS meetings;
DROP TABLE IF EXISTS user_groups;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS users;

-- Removendo ENUMs
DROP TYPE IF EXISTS visit_category_enum;
DROP TYPE IF EXISTS visit_type_enum;
DROP TYPE IF EXISTS cleaning_type_enum;
DROP TYPE IF EXISTS weekday_enum;
DROP TYPE IF EXISTS assignment_type_enum;
DROP TYPE IF EXISTS meeting_type_enum;
DROP TYPE IF EXISTS user_role_enum;
