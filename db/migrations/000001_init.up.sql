-- ENUMs para otimizar espaço e melhorar performance
CREATE TYPE meeting_type_enum AS ENUM ('Midweek Meeting', 'Weekend Meeting');

CREATE TYPE assignment_type_enum AS ENUM (
    'Reading', 'Demonstration Student', 'Demonstration Assistant', 'Discourse', 'Prayer',
    'Chairman', 'Watchtower Reader', 'Watchtower Conductor', 'Congregation Bible Study Conductor',
    'Congregation Bible Study Reader', 'Attendants', 'Sound Operator', 'Microphone Operator',
    'Video Operator', 'Platform Assistant', 'Field Service Conductor', 'Public Talk Speaker'
);

CREATE TYPE weekday_enum AS ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday');

CREATE TYPE cleaning_type_enum AS ENUM ('Before Meeting', 'After Meeting');

CREATE TYPE visit_type_enum AS ENUM ('Nobody Home', 'Busy', 'Visited', 'Letter');

CREATE TYPE visit_category_enum AS ENUM ('Campaign', 'Normal');

-- Tabela de Usuários
CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE,
    role TEXT NOT NULL CHECK (
        role IN ('ELDER', 'MINISTERIAL_SERVANT', 'AUXILIAR_PIONEER', 'REGULAR_PIONEER', 'STUDENT', 'PUBLISHER')
    )
);

-- Tabela de Grupos
CREATE TABLE groups (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    leader_id BIGINT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    assistant_id BIGINT REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE user_groups (
    user_id BIGINT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE
);

-- Tabela de Reuniões
CREATE TABLE meetings (
    id BIGINT PRIMARY KEY,  -- Agora `id` é a única chave primária
    meeting_type meeting_type_enum NOT NULL,
    meeting_date DATE NOT NULL  -- Continua existindo, mas não faz parte da PK
);

-- Tabela de Salas para reuniões
CREATE TABLE rooms (
    id BIGINT PRIMARY KEY,
    meeting_id BIGINT NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    room_name TEXT NOT NULL
);

-- Tabela de Sessões
CREATE TABLE session_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE sessions (
    id BIGINT PRIMARY KEY,
    room_id BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    session_type_id INT NOT NULL REFERENCES session_types(id),
    UNIQUE (room_id, session_type_id)
);

-- Tabela de Subsessões
CREATE TABLE subsession_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE subsessions (
    id BIGINT PRIMARY KEY,
    session_id BIGINT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    subsession_type_id INT NOT NULL REFERENCES subsession_types(id),
    duration_minutes SMALLINT NOT NULL CHECK (duration_minutes > 0)
);

-- Tabela de Designações
CREATE TABLE assignments (
    id BIGINT PRIMARY KEY,
    meeting_id BIGINT,
    meeting_date DATE,
    subsession_id BIGINT REFERENCES subsessions(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    assignment_type assignment_type_enum NOT NULL,    
    CHECK (
        (meeting_id IS NOT NULL AND meeting_date IS NOT NULL AND subsession_id IS NULL) OR
        (meeting_id IS NULL AND meeting_date IS NULL AND subsession_id IS NOT NULL)
    ),
    CONSTRAINT unique_meeting_assignment UNIQUE (meeting_id, meeting_date, user_id, assignment_type),
    CONSTRAINT unique_subsession_assignment UNIQUE (subsession_id, user_id, assignment_type)
) WITH (fillfactor = 70);


-- Designações de Limpeza
CREATE TABLE cleaning_assignments (
    id BIGINT PRIMARY KEY,
    group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    meeting_id BIGINT NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    cleaning_type cleaning_type_enum NOT NULL,
    UNIQUE (meeting_id, cleaning_type)
);

-- Tabela de Carrinhos e Turnos
CREATE TABLE carts (
    id BIGINT PRIMARY KEY,
    location TEXT NOT NULL,
    description TEXT
);

CREATE TABLE cart_shifts (
    id BIGINT PRIMARY KEY,
    cart_id BIGINT NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    shift_day weekday_enum NOT NULL,
    start_time TIME NOT NULL CHECK (start_time >= '07:00:00' AND start_time <= '21:00:00'),
    end_time TIME NOT NULL CHECK (end_time > start_time AND end_time <= '21:00:00'),
    UNIQUE (cart_id, shift_day, start_time)
);

CREATE TABLE cart_assignments (
    id BIGINT PRIMARY KEY,
    shift_id BIGINT NOT NULL REFERENCES cart_shifts(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    UNIQUE (shift_id, user_id)
);

-- Tabela de Territórios e Visitas
CREATE TABLE territories (
    id BIGINT PRIMARY KEY,
    territory_number INT NOT NULL UNIQUE,
    location TEXT NOT NULL,
    shapefile BYTEA NOT NULL,
    completed_at DATE DEFAULT NULL
);

CREATE TABLE house_visits (
    id BIGINT PRIMARY KEY,
    territory_id BIGINT NOT NULL REFERENCES territories(id) ON DELETE CASCADE,
    house_number TEXT NOT NULL,
    visit_date DATE NOT NULL,
    visit_type visit_type_enum NOT NULL,
    visit_category visit_category_enum NOT NULL,
    UNIQUE (territory_id, house_number, visit_date)
);

-- Índices para otimizar buscas
CREATE INDEX idx_meetings_date ON meetings (meeting_date);
CREATE INDEX idx_sessions_room ON sessions (room_id);
CREATE INDEX idx_assignments_user ON assignments (user_id);
CREATE INDEX idx_cart_shifts_day ON cart_shifts (shift_day);
CREATE INDEX idx_cart_assignments_user ON cart_assignments (user_id);
CREATE INDEX idx_cleaning_meeting ON cleaning_assignments (meeting_id);
CREATE INDEX idx_house_visits_territory ON house_visits (territory_id);
