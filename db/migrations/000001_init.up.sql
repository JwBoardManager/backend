-- ENUMs para otimizar espaço e melhorar performance
CREATE TYPE meeting_type_enum AS ENUM ('Midweek Meeting', 'Weekend Meeting');
CREATE TYPE assignment_type_enum AS ENUM (
    'Reading',
    'Demonstration Student',
    'Demonstration Assistant',
    'Discourse',
    'Prayer',
    'Chairman',
    'Watchtower Reader',
    'Watchtower Conductor',
    'Congregation Bible Study Conductor',
    'Congregation Bible Study Reader',
    'Attendants',
    'Sound Operator',
    'Microphone Operator',
    'Video Operator',
    'Platform Assistant',
    'Field Service Conductor',
    'Public Talk Speaker'
);

CREATE TYPE weekday_enum AS ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday');
CREATE TYPE cleaning_type_enum AS ENUM ('Before Meeting', 'After Meeting');
-- ENUM para tipos de visita
CREATE TYPE visit_type_enum AS ENUM ('Nobody Home', 'Busy', 'Visited', 'Letter');

CREATE TYPE user_role_enum AS ENUM ('Elder', 'Ministerial_Servant', 'auxiliar_pioneer', 'regular_pioneer', 'student', 'publisher');

-- ENUM para categoria da visita
CREATE TYPE visit_category_enum AS ENUM ('Campaign', 'Normal');

-- Tabela de Usuários
CREATE TABLE users (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    name TEXT NOT NULL,
    email TEXT UNIQUE,
    password TEXT NOT NULL,
    role user_role_enum NOT NULL DEFAULT 'student'
);

-- Tabela de Grupos
CREATE TABLE groups (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    name TEXT NOT NULL UNIQUE,
    leader_id BIGINT NOT NULL REFERENCES users(id) ON DELETE SET NULL, -- Responsável pelo grupo
    assistant_id BIGINT REFERENCES users(id) ON DELETE SET NULL -- Ajudante do grupo
);

-- Relacionamento: Usuários pertencem a um Grupo
CREATE TABLE user_groups (
    user_id BIGINT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE
);

-- Tabela de Reuniões
CREATE TABLE meetings (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    meeting_type meeting_type_enum NOT NULL,
    meeting_date DATE NOT NULL
);

-- Tabela de Salas para reuniões
CREATE TABLE rooms (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    meeting_id BIGINT NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    room_name TEXT NOT NULL
);

-- Tabela de Tipos de Sessões (evita repetição de texto)
CREATE TABLE session_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- Tabela de Sessões
CREATE TABLE sessions (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    room_id BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    session_type_id INT NOT NULL REFERENCES session_types(id),
    UNIQUE (room_id, session_type_id)
);

-- Tabela de Tipos de Subsessões (evita repetição de texto)
CREATE TABLE subsession_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- Tabela de Subsessões
CREATE TABLE subsessions (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    session_id BIGINT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    subsession_type_id INT NOT NULL REFERENCES subsession_types(id),
    duration_minutes SMALLINT NOT NULL CHECK (duration_minutes > 0)
);

CREATE TABLE assignments (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    meeting_id BIGINT REFERENCES meetings(id) ON DELETE CASCADE, -- Para designações da reunião inteira
    subsession_id BIGINT REFERENCES subsessions(id) ON DELETE CASCADE, -- Para designações de uma parte da reunião
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    assignment_type assignment_type_enum NOT NULL,
    CHECK (
        (meeting_id IS NOT NULL AND subsession_id IS NULL) OR
        (meeting_id IS NULL AND subsession_id IS NOT NULL)
    ), -- Garante que a designação está em UMA das categorias
    CONSTRAINT unique_meeting_assignment UNIQUE (meeting_id, user_id, assignment_type),
    CONSTRAINT unique_subsession_assignment UNIQUE (subsession_id, user_id, assignment_type)
) WITH (fillfactor = 70);



-- Tabela de Designações de Limpeza para os Grupos
CREATE TABLE cleaning_assignments (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    meeting_id BIGINT NOT NULL REFERENCES meetings(id) ON DELETE CASCADE,
    cleaning_type cleaning_type_enum NOT NULL, -- Antes ou Depois da reunião
    UNIQUE (meeting_id, cleaning_type) -- Garante que não haja duas equipes na mesma limpeza
);

-- Tabela de Carrinhos
CREATE TABLE carts (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    location TEXT NOT NULL, -- Local onde o carrinho está alocado
    description TEXT -- Informação adicional sobre o carrinho
);

-- Tabela de Turnos para o Serviço em Carrinhos
CREATE TABLE cart_shifts (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    cart_id BIGINT NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    shift_day INT NOT NULL, -- Representa os dias da semana
    start_time TIME NOT NULL CHECK (start_time >= '07:00:00' AND start_time <= '21:00:00'), 
    end_time TIME NOT NULL CHECK (end_time > start_time AND end_time <= '21:00:00'),
    UNIQUE (cart_id, shift_day, start_time) -- Evita turnos duplicados no mesmo carrinho
);

-- Tabela de Designação de Publicadores ao Serviço em Carrinhos
CREATE TABLE cart_assignments (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    shift_id BIGINT NOT NULL REFERENCES cart_shifts(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    UNIQUE (shift_id, user_id) -- Evita duplicação de publicadores no mesmo turno
);

CREATE TABLE territories (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    territory_number INT NOT NULL UNIQUE, -- Número identificador do território
    location TEXT NOT NULL, -- Descrição da localização
    shapefile BYTEA NOT NULL, -- Arquivo de forma do território
    completed_at DATE DEFAULT NULL -- Data de conclusão do território (NULL se ainda não foi concluído)
);

-- Tabela de Registros de Casas Visitadas
CREATE TABLE house_visits (
    id BIGINT PRIMARY KEY, -- Snowflake ID
    territory_id BIGINT NOT NULL REFERENCES territories(id) ON DELETE CASCADE,
    house_number TEXT NOT NULL, -- Número da casa (pode incluir letras, ex: "123A")
    visit_date DATE NOT NULL, -- Data da visita
    visit_type visit_type_enum NOT NULL, -- Tipo da visita
    visit_category visit_category_enum NOT NULL, -- Se foi campanha ou normal
    UNIQUE (territory_id, house_number, visit_date) -- Garante que não haja visitas duplicadas no mesmo dia para a mesma casa
);


-- Índices para otimizar buscas
CREATE INDEX idx_meetings_date ON meetings (meeting_date);
CREATE INDEX idx_sessions_room ON sessions (room_id);
CREATE INDEX idx_assignments_user ON assignments (user_id);
CREATE INDEX idx_cart_shifts_day ON cart_shifts (shift_day);
CREATE INDEX idx_cart_assignments_user ON cart_assignments (user_id);
CREATE INDEX idx_cleaning_meeting ON cleaning_assignments (meeting_id);
